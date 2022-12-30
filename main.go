package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/araddon/dateparse"
)

var (
	version = "unset"
	commit  = "unset"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "-v" {
		fmt.Printf("%s-%s", version, commit)
		os.Exit(0)
	}

	input := strings.Join(os.Args[1:], " ")
	parsed, err := dateparse.ParseAny(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(parsed.Format(time.RFC1123Z))
}
