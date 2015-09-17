///bin/golang-tip-to-run-as-a-shell-script "$0" ; exec /usr/bin/env go run "$0" "$@"
package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/docker/libchan/spdy"
	// "github.com/joao-parana/msf/nano"
	nano "../nano"
)

func main() {
	log.Println("Starting Server..")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}

	tl, err := spdy.NewTransportListener(listener, spdy.NoAuthenticator)
	if err != nil {
		log.Fatal(err)
	}
	adapter := nano.NewThingeyAdapter()
	for {
		t, err := tl.AcceptTransport()
		if err != nil {
			log.Print(err)
			break
		}

		go func() {
			for {
				receiver, err := t.WaitReceiveChannel()
				if err != nil {
					log.Print(err)
					break
				}

				go func() {
					for {
						err := adapter.Listen(receiver)
						if err != nil {
							break
						}
					}
				}()
			}
		}()
	}

}
