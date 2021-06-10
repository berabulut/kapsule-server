package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

type WebHook struct {
	Ref string      `json:"ref"`
	PR  PullRequest `json:"pull_request"`
}

type PullRequest struct {
	Merged bool `json:"merged"`
}

func HookHandler(folder string) (http.HandlerFunc) { 
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

		if body.PR.Merged {
			command := fmt.Sprintf("./pull.sh %s", folder)
			cmd, err := exec.Command(command).Output()
		
			if err != nil {
				log.Fatal(err)
			}			
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

func executeCommand ()

func main() {
	address := ":4043"

	http.HandleFunc("/kapsule", HookHandler("kapsule"))
	http.HandleFunc("/kapsule-ui", HookHandler("kapsule-ui"))
	http.HandleFunc("/kapsule-server", HookHandler("kapsule-server"))
	
	log.Println("Starting server on address", address)
	
	err := http.ListenAndServe(address, nil)
	
	if err != nil {
		panic(err)
	}
}
