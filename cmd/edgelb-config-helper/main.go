package main

import (
	"fmt"
	"time"

	"log"
	"os"

	"github.com/go-openapi/strfmt"
	"github.com/iss-lab/edgelb-config-helper/client/operations"

	// "github.com/go-openapi/spec"

	httptransport "github.com/go-openapi/runtime/client"
	apiclient "github.com/iss-lab/edgelb-config-helper/client"
)

func main() {
	transport := httptransport.New(os.Getenv("EDGELB_HOST"), "", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	pingResp, err := client.Operations.Ping(operations.NewPingParamsWithTimeout(60 * time.Second))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", pingResp.Payload)

	configResp, err := client.Operations.GetConfigContainer(operations.NewGetConfigContainerParams())
	if err != nil {
		log.Fatal(err)
	}
	config := configResp.Payload
	configBytes, err := config.MarshalBinary()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", string(configBytes))

	fmt.Println("Pools:")

	for _, p := range config.Pools {
		// pBytes, err := p.MarshalBinary()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Printf("%s\n", string(pBytes))
		fmt.Printf("%s\n", p.Name)
	}
}
