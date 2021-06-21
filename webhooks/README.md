# webhooks-server

Listens **kapsule**, **kapsule-ui** and **kapsule-server** repos. When there is a merged pull request, server executes a shell script that pulls and builds related repo. 

## How To Run

- Development

	`go run main.go`

- Deployment

	`./deploy.sh`
