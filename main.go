////////////////////////////////////////////////////////////////////////////
// Program: cascadia
// Purpose: go cascadia CSS selection from command line
// Authors: Tong Sun (c) 2016, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"

	"github.com/andybalholm/cascadia"
	"github.com/mkideal/cli"
	"golang.org/x/net/html"
)

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var progname = "cascadia"
var buildTime = "2016-12-10"

var rootArgv *rootT

////////////////////////////////////////////////////////////////////////////
// Function definitions

func main() {
	if err := cli.Root(root).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println("")
}

func cascadiaC(ctx *cli.Context) error {
	// ctx.JSON(ctx.RootArgv())
	// ctx.JSON(ctx.Argv())
	// fmt.Println()

	argv := ctx.Argv().(*rootT)
	doc, err := html.Parse(argv.Filei)
	argv.Filei.Close()
	abortOn("Input", err)

	ss := argv.CSS
	c := cascadia.MustCompile(ss)
	ns := c.MatchAll(doc)
	fmt.Fprintf(os.Stderr, "%d elements for '%s':\n", len(ns), ss)
	for _, n := range ns {
		html.Render(argv.Fileo, n)
		fmt.Println()
	}
	argv.Fileo.Close()
	return nil
}

// abortOn will quit on anticipated errors gracefully without stack trace
func abortOn(errCase string, e error) {
	if e != nil {
		fmt.Printf("[%s] %s error: %v\n", progname, errCase, e)
		os.Exit(1)
	}
}
