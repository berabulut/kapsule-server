# webhooks-server

Listens [kapsule](https://github.com/berabulut/kapsule), [kapsule-ui](https://github.com/berabulut/kapsule-ui) and [kapsule-server](https://github.com/berabulut/kapsule-server). 

Deployment conditions:

- kapsule:

	- When a GitHub action succeeds (Currently there is only one GH Action)

- kapsule-ui:

	- When a GitHub action succeeds (Currently there is only one GH Action)

- kapsule-server:

	- Merged Pull Request



## How To Run

- Development

	`go run main.go`

- Deployment

	`sh run.sh`
