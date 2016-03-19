package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	flags "github.com/jessevdk/go-flags"
	"github.com/yarbelk/stepon/libstepon"
)

type options struct {
	Positional struct {
		Units int     `positional-arg-name:"units"`
		Price float64 `positional-arg-name:"price"`
		State string  `positional-arg-name:"state"`
	} `positional-args:"yes" required:"yes"`
}

var opts options

func parseAndRun(parser *flags.Parser, arglist []string) {
	if _, err := parser.ParseArgs(arglist); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	p := opts.Positional
	fmt.Println(stepon.GetTotal(p.Units, p.Price, p.State))
}

func runFromStdin() {
	var parser *flags.Parser
	var arglist []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		opts = options{}
		parser = flags.NewParser(&opts, flags.None)
		arglist = strings.Split(scanner.Text(), " ")

		parseAndRun(parser, arglist)
	}
	if err := scanner.Err(); err != nil {
		os.Exit(1)
	}
}

func main() {
	opts = options{}
	parser := flags.NewParser(&opts, flags.HelpFlag)
	switch len(os.Args) {
	case 1: // Read from stdin
		runFromStdin()
	case 4: // CMD units price state
		parseAndRun(parser, os.Args[1:])
	default:
		_, err := parser.Parse()
		fmt.Println(err.Error())
	}
}
