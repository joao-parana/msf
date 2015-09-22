///bin/golang-tip-to-run-as-a-shell-script "$0" ; exec /usr/bin/env go run "$0" "$@"
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"syscall"

	// "github.com/joao-parana/libchan/spdy"
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

	// tl, err := spdy.NewTransportListener(listener, spdy.NoAuthenticator)
	// if err != nil {
	//  log.Fatal(err)
	// }
	provider, err := spdy.NewSpdyStreamProvider(conn, true)
	// ...
	tl := provider.Listen()
	adapter := nano.NewThingeyAdapter()
	for {
		// t, err := tl.AcceptTransport()
		t, err := tl.Accept()
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

// RemoteCommand is the received command parameters to execute locally and return
type RemoteCommand struct {
	Cmd        string
	Args       []string
	Stdin      io.Reader
	Stdout     io.WriteCloser
	Stderr     io.WriteCloser
	StatusChan libchan.Sender
}

// CommandResponse is the response struct to return to the client
type CommandResponse struct {
	Status int
}

func xmain() {
	var listener net.Listener

	var err error
	listener, err = net.Listen("tcp", "localhost:9323")
	if err != nil {
		log.Fatal(err)
	}

	for {
		c, err := listener.Accept()
		if err != nil {
			log.Print(err)
			break
		}
		p, err := spdy.NewSpdyStreamProvider(c, true)
		if err != nil {
			log.Print(err)
			break
		}
		t := spdy.NewTransport(p)

		go func() {
			for {
				receiver, err := t.WaitReceiveChannel()
				if err != nil {
					log.Print(err)
					break
				}

				go func() {
					for {
						command := &RemoteCommand{}
						err := receiver.Receive(command)
						if err != nil {
							log.Print(err)
							break
						}

						cmd := exec.Command(command.Cmd, command.Args...)
						cmd.Stdout = command.Stdout
						cmd.Stderr = command.Stderr

						stdin, err := cmd.StdinPipe()
						if err != nil {
							log.Print(err)
							break
						}
						go func() {
							io.Copy(stdin, command.Stdin)
							stdin.Close()
						}()

						res := cmd.Run()
						command.Stdout.Close()
						command.Stderr.Close()
						returnResult := &CommandResponse{}
						if res != nil {
							if exiterr, ok := res.(*exec.ExitError); ok {
								returnResult.Status = exiterr.Sys().(syscall.WaitStatus).ExitStatus()
							} else {
								log.Print(res)
								returnResult.Status = 10
							}
						}

						err = command.StatusChan.Send(returnResult)
						if err != nil {
							log.Print(err)
						}
					}
				}()
			}
		}()
	}
}
