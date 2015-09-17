///bin/golang-tip-to-run-as-a-shell-script "$0" ; exec /usr/bin/env go run "$0" "$@"
package main

import (
	"log"

	"github.com/joao-parana/msf/nano"
	// nano "../nano"
)

func main() {
	log.Println("Starting cmd to test locat ..")
	repo := nano.NewLocalRepository()

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
