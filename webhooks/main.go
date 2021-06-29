package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

const ActionJobName = "build-and-push"

var (
	f   *os.File
	err error
)

type WebHook struct {
	Ref      string      `json:"ref"`
	Action   string      `json:"action"`
	CheckRun CheckRun    `json:"check_run"`
	PR       PullRequest `json:"pull_request"`
}

type CheckRun struct {
	Name       string `json:"name"`
	Conclusion string `json:"conclusion"`
}

type PullRequest struct {
	Merged bool `json:"merged"`
}

func init() {

	// open log file, create if it doesn't exist
	f, err = os.OpenFile("/tmp/webhooks.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	// set write logs to log file
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)

	// load environment variables
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal(err)
	}
}
func HookHandler(folder string) http.HandlerFunc {
	// Read body
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// Unmarshal
		var body WebHook
		err = json.Unmarshal(b, &body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		if folder == "kapsule-server" && body.PR.Merged {
			log.Println("kapsule-server pull request merged")
			go executeScript(folder)
		}

		if body.Action == "completed" && body.CheckRun.Conclusion == "success" && body.CheckRun.Name == ActionJobName {
			//command := fmt.Sprintf("./build.sh %s", folder)
			//cmd, err := exec.Command("./build.sh", folder).Output()
			log.Println("GH action succeeded!")
			go executeScript(folder)
		}

		output, err := json.Marshal(body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)
	}
}

func executeScript(folder string) {

	log.Printf("Executing script with parameter related **%s**", folder)

	cmd, err := exec.Command("./update.sh", folder).CombinedOutput()
	output := string(cmd)
	fmt.Println(output)

	if err != nil {
		log.Fatal(err)
	}

}

func main() {

	address := fmt.Sprintf(":%s", os.Getenv("WEBHOOKS_SERVER_PORT"))

	http.HandleFunc("/kapsule", HookHandler("kapsule"))
	http.HandleFunc("/kapsule-ui", HookHandler("kapsule-ui"))
	http.HandleFunc("/kapsule-server", HookHandler("kapsule-server"))

	log.Println("Starting server on address", address)

	// log file
	defer f.Close()

	err := http.ListenAndServe(address, nil)

	if err != nil {
		log.Fatal(err)
	}
}
