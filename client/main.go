///bin/golang-tip-to-run-as-a-shell-script "$0" ; exec /usr/bin/env go run "$0" "$@"
package main

import (
	"fmt"
	"log"
	"os"

	nano "../nano"
	// "github.com/joao-parana/msf/nano"
)

func main() {
	log.Println("Starting..")
	ip := os.Getenv("IP")
	if ip == "" {
		ip = "localhost"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Connecting to %s:%s\n", ip, port)
	repo := nano.NewRemoteRepository(fmt.Sprintf("%s:%s", ip, port))

	t1 := &nano.Thingey{
		ID:   "1",
		Data: "d1",
	}
	if err := repo.Create(t1); err != nil {
		log.Fatal(err)
	}
	log.Printf("Created %+v\n", t1.Data)

	if t2, err := repo.Get(t1.ID); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Got %+v\n", t2.Data)
	}

	if tList, err := repo.List(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Listed %d items\n", len(tList))
	}

	if err := repo.Delete(t1); err != nil {
		log.Fatal(err)
	}
	log.Printf("Deleted %+v\n", t1.Data)

	if tList, err := repo.List(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Listed %d items\n", len(tList))
	}
}
