package main

import (
	"flag"

	"github.com/ivan-kostko/goq/cmd"
)

var (
	serviceRpcPort  = flag.String("rpc", "", "The service port to listen for RPC calls")
	serviceRestPort = flag.String("rest", "", "The service port to listen for REST calls")
	webAppPort      = flag.String("web", "", "The web application UI port")
)

func main() {
	flag.Parse()
	runner := cmd.Compose()
	runner()
}
