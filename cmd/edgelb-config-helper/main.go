package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	apiclient "github.com/iss-lab/edgelb-config-helper/client"
	"github.com/iss-lab/edgelb-config-helper/client/operations"
	"github.com/iss-lab/edgelb-config-helper/models"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	dnsHost := os.Getenv("DNS_HOST")
	publicHost := os.Getenv("PUBLIC_HOST")

	// edgelbHost := os.Getenv("EDGELB_HOST")
	// config, err := fetchConfig(edgelbHost)

	poolFile := os.Getenv("POOL_FILE")
	config := readConfig(poolFile)

	poolToLinks(config, dnsHost, publicHost)
}

func readConfig(poolFile string) *models.V2Pool {
	dat, err := ioutil.ReadFile(poolFile)
	check(err)

	pool := &models.V2Pool{}
	err = pool.UnmarshalBinary(dat)
	check(err)

	return pool
}

func fetchConfig(edgelbHost, poolName string) *models.V2Pool {
	transport := httptransport.New(edgelbHost, "", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)

	pingResp, err := client.Operations.Ping(operations.NewPingParamsWithTimeout(60 * time.Second))
	check(err)
	fmt.Printf("%#v\n", pingResp.Payload)

	configResp, err := client.Operations.GetConfigContainer(operations.NewGetConfigContainerParams())
	check(err)

	config := configResp.Payload
	for _, pool := range config.Pools {
		if pool.Name == poolName {
			return pool.V2
		}
	}

	log.Fatal(fmt.Sprintf("Could not find pool named: '%s'", poolName))
	return nil
}

func poolToLinks(pool *models.V2Pool, dnsHost, publicHost string) {
	fmt.Printf("    links:\n")
	for _, frontend := range pool.Haproxy.Frontends {
		for _, link := range frontend.LinkBackend.Map {
			fmt.Printf("      - %s -> %s\n", displayLink(publicHost, dnsHost, frontend.BindPort, link), link.Backend)
		}
		fmt.Printf("      - %s -> %s\n", displayHost(publicHost, dnsHost, frontend.BindPort, nil), frontend.LinkBackend.DefaultBackend)
	}
}

func displayHost(publicHost, dnsHost string, port *int32, link *models.V2FrontendLinkBackendMapItems0) string {
	host := dnsHost
	if *port != 80 && *port != 443 {
		host = fmt.Sprintf("%s:%d", publicHost, *port)
	}

	if link != nil {
		if len(link.HostReg) > 0 {
			hostCopy := strings.ReplaceAll(link.HostReg, "\\.", ".")
			hostCopy = strings.Replace(hostCopy, "^", "", 1)
			if strings.Contains(hostCopy, ".*$") {
				hostCopy = strings.Replace(hostCopy, ".*$", host, 1)
			} else {
				hostCopy = strings.TrimSuffix(hostCopy, ".*")
				hostCopy = hostCopy + host
			}
			host = strings.ReplaceAll(hostCopy, ".*", "*")
		} else if len(link.HostEq) > 0 {
			host = link.HostEq
		}
	}

	return host
}

// Basically reverse regex the linkBackend map item
func displayLink(publicHost, dnsHost string, port *int32, link *models.V2FrontendLinkBackendMapItems0) string {
	host := displayHost(publicHost, dnsHost, port, link)
	path := ""
	if link != nil {
		if len(link.PathReg) > 0 {
			path = link.PathReg
		} else if len(link.PathBeg) > 0 || len(link.PathEnd) > 0 {
			path = link.PathBeg + "/*/" + link.PathEnd
		}
	}

	return fmt.Sprintf("%s%s", host, path)
}

// ---
// vhosts:
// 	- postgresql-admin
// 		marathon: /data/postgresql-admin
// 		portName: postgresql-admin
// 	-

// TODO: Support "simple" vhost yaml syntax + mustache template for rest of pool json config
// TODO: Support pulling straight from API server as well as from static json file
// TODO: Support dcos cli based connection to api server via https and dcos authentication
// TODO: Improve display of backend information
// TODO: Add https and http scheme based on some simple yaml logic
// TODO: Correctly show dns vs public host based on protocol (http(s) vs tcp/tls)
