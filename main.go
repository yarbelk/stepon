package main

import (
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/yarbelk/stepon/libstepon"
)

var opts = struct {
	Positional struct {
		Units int     `positional-arg-name:"units"`
		Price float64 `positional-arg-name:"price"`
		State string  `positional-arg-name:"state"`
	} `positional-args:"yes" required:"yes"`
}{}

var parser *flags.Parser

func init() {
	parser = flags.NewParser(&opts, flags.None)
}

func main() {
	if _, err := parser.Parse(); err != nil {
		parser.WriteHelp(os.Stderr)
		fmt.Println(err.Error())
		os.Exit(1)
	}
	p := opts.Positional
	fmt.Println(stepon.GetTotal(p.Units, p.Price, p.State))
}
