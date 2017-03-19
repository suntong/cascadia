////////////////////////////////////////////////////////////////////////////
// Program: cascadia
// Purpose: go cascadia CSS selection from command line
// Authors: Tong Sun (c) 2016, All rights reserved
////////////////////////////////////////////////////////////////////////////

//go:generate sh -v cascadiaGenCLI.sh

package main

import (
	"fmt"
	"io"
	"os"

	"github.com/andybalholm/cascadia"
	"github.com/mkideal/cli"
	"golang.org/x/net/html"
)

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var progname = "cascadia"
var buildTime = "2016-12-12"

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
	Cascadia(argv.Filei, argv.Fileo, argv.CSS)
	argv.Filei.Close()
	argv.Fileo.Close()
	return nil
}

// Cascadia filters the input buffer/stream `bi` with CSS selectors `css` and write to the output buffer/stream `bw`.
func Cascadia(bi io.Reader, bw io.Writer, css string) error {
	ss := css
	c := cascadia.MustCompile(ss)

	doc, err := html.Parse(bi)
	abortOn("Input", err)

	ns := c.MatchAll(doc)
	fmt.Fprintf(os.Stderr, "%d elements for '%s':\n", len(ns), ss)
	for _, n := range ns {
		html.Render(bw, n)
		fmt.Fprintf(bw, "\n")
	}
	return nil
}

// abortOn will quit on anticipated errors gracefully without stack trace
func abortOn(errCase string, e error) {
	if e != nil {
		fmt.Printf("[%s] %s error: %v\n", progname, errCase, e)
		os.Exit(1)
	}
}
