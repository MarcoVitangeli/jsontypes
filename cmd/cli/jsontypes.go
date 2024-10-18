package main

import (
	"os"
	"fmt"
	"io"
	"github.com/MarcoVitangeli/jsontypes/gen"
	"strconv"
)

const (
	defaultMaxDepth = 10
)

func main() {
	args := os.Args[1:]
	
	if len(args) >= 1 {
		runWithArgs(args)	
	} else {
		fmt.Println("Invalid program usage, see help below:\n")
		usage()
	}
}

func runWithArgs(args []string) {
	p := args[0]
	f, err := os.Open(p)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: error opening file: %s", err)
		os.Exit(-1)
	}
	
	bs, err := io.ReadAll(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: error reading file: %s", err)
		os.Exit(-2)
	}

	cliArgs, err := parseArgs(args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: error parsing arguments: %s", err)
		os.Exit(-3)
	}


	if err := gen.Gen(bs, cliArgs...); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: error generating go types: %s", err)
		os.Exit(-4)
	}
}

func parseArgs(args []string) ([]gen.GenOption, error) {
	var g []gen.GenOption
	i := 0
	for i < len(args) {
		v := args[i]
		switch v {
		case "-d", "--depth":
			if len(args) <= i+1 {
				return nil, fmt.Errorf("ERROR: missing depth argument")
			}
			if v, err := strconv.ParseUint(args[i+1], 10, 64); err != nil {
				return nil, fmt.Errorf("invalid depth value provided, must be a number")
			} else {
				g = append(g, gen.WithDepth(uint(v)))
				i += 2
			}
		default:
			return nil, fmt.Errorf("invalid argument provided: %s", v)
		}
	}
	return g, nil
}

func usage() {
	fmt.Printf(`Program usage: jsontypes [PATH] [OPTIONS]
	PATH:	path to json file for Go type generation
	OPTIONS:
		-d,--depth <DEPTH>	number. Max depth for JSON object traversal. Defaults to 10.`)
}
