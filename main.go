package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/eitarox/eitaroxgit/object"
	"github.com/eitarox/eitaroxgit/store"
)

func main() {
	hashString := os.Args[1]
	hash, err := hex.DecodeString(hashString)
	if err != nil {
		log.Fatal(err)
	}

	client, err := store.NewClient(".")
	if err != nil {
		log.Fatal(err)
	}
	if err := client.WalkHistory(hash, func(commit *object.Commit) error {
		fmt.Println(commit)
		fmt.Println("")
		return nil
	}); err != nil {
		log.Fatal(err)
	}

}
