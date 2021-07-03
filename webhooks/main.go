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
	"time"

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

	// Open log file, create if it doesn't exist
	f, err = os.OpenFile("/tmp/webhooks.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	// Set write logs to log file
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)

	// Load environment variables
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal(err)
	}
}

// Handle HTTP request
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

		// When kapsule-server repo has merged pull request deploy changes.
		if folder == "kapsule-server" && body.PR.Merged {

			log.Println("kapsule-server pull request merged")

			queue = append(queue, folder)
			go executeScript()
		}

		// When a GitHub action related to kapsule and kapsule-ui succeed deploy changes.
		if body.Action == "completed" && body.CheckRun.Conclusion == "success" && body.CheckRun.Name == ActionJobName {
			//command := fmt.Sprintf("./build.sh %s", folder)
			//cmd, err := exec.Command("./build.sh", folder).Output()

			log.Println("GH action succeeded!")

			queue = append(queue, folder)
			go executeScript()
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

// Run deployment script while respecting deployment queue.
func executeScript() {

	// We will wait until previous deployment finish
	for ongoingDeployment() {
		log.Println("Waiting end of current deployment!")
		time.Sleep(time.Second)
	}

	log.Printf("Executing script with parameter related **%s**", queue[0])
	setDeploymentStatus(true)

	cmd, err := exec.Command("./update.sh", queue[0]).CombinedOutput()
	output := string(cmd)
	log.Println(output)

	setDeploymentStatus(false)
	orderQueue()

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

	// Close log file
	defer f.Close()

	err := http.ListenAndServe(address, nil)

	if err != nil {
		log.Fatal(err)
	}
}
