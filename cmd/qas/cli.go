package main

import (
	"flag"
	"fmt"

	"os"
)

const (
	description = "Easily manage multiple FLOSS repositories."
	name        = "qas"
)

func cliParse() (*string, *bool, *bool) {
	backup := flag.String("backup", "", "archive floss projects listed on NAMES")
	grab := flag.Bool("grab", false, "grab floss projects")
	verbose := flag.Bool("verbose", false, "display more information")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "\nqas - Easily manage multiple FLOSS repositories.\n")
		fmt.Fprintln(flag.CommandLine.Output(), "\nUsage information:")
		flag.PrintDefaults()
	}
	flag.Parse()

	if !*grab && *backup == "" {
		flag.Usage()
		os.Exit(0)
	}

	return backup, grab, verbose
}
