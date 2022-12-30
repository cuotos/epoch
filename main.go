package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/araddon/dateparse"
)

func main() {
	input := strings.Join(os.Args[1:], " ")
	parsed, err := dateparse.ParseAny(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(parsed.Format(time.RFC1123Z))
}
