package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/araddon/dateparse"
)

var (
	version = "unset"
	commit  = "unset"

	defaultFormat = time.RFC1123Z
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "-v" {
		fmt.Printf("%s-%s", version, commit)
		os.Exit(0)
	}

	input := strings.Join(os.Args[1:], " ")

	// assume no date provided, if there is one, parse it and use that.
	parsed := time.Now()

	if input != "" {
		var err error
		parsed, err = dateparse.ParseAny(input)
		if err != nil {
			log.Fatal(err)
		}
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', 0)

	fmt.Fprintln(w, "Human:\t", parsed.Format(defaultFormat))
	fmt.Fprintln(w, "UTC:\t", parsed.UTC().Format(defaultFormat))
	fmt.Fprintln(w, "Unix:\t", parsed.Unix())
	fmt.Fprintln(w, "Unix(ms):\t", parsed.UnixMilli())
	w.Flush()
}
