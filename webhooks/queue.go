package main

// Resp is deployment_status.
type readOp struct {
	resp chan bool
}

// Deployment_status is true when there is an ongoing deployment.
// Resp is whether writing is successfull or not.
type writeOp struct {
	deployment_status bool
	resp              chan bool
}

var (
	reads  chan readOp
	writes chan writeOp
	queue  []string
)

func init() {

	reads = make(chan readOp)
	writes = make(chan writeOp)

	// Read and write deployment status.
	go func() {
		var state = false
		for {
			select {
			case read := <-reads:
				read.resp <- state
			case write := <-writes:
				state = write.deployment_status
				write.resp <- true
			}
		}
	}()
}

func orderQueue() {
	if len(queue) > 0 {
		queue = queue[1:]
	}
}

func ongoingDeployment() bool {
	read := readOp{
		resp: make(chan bool)}
	reads <- read

	resp := <-read.resp

	return resp
}

func setDeploymentStatus(status bool) {
	write := writeOp{
		deployment_status: status,
		resp:              make(chan bool)}
	writes <- write
	<-write.resp
}
