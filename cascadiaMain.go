////////////////////////////////////////////////////////////////////////////
// Program: cascadia
// Purpose: go cascadia CSS selection from command line
// Authors: Tong Sun (c) 2016-2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

//go:generate sh -v cascadiaCLIGen.sh

package main

import (
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/andybalholm/cascadia"
	"github.com/mkideal/cli"
	"golang.org/x/net/html"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

type MapStringString struct {
	Keys   []string
	Values map[string]string
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var progname = "cascadia"
var buildTime = "2017-03-22"

var rootArgv *rootT

////////////////////////////////////////////////////////////////////////////
// Function definitions

func main() {
	if err := cli.Root(root).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println("")
}

//==========================================================================
// css selection

func cascadiaC(ctx *cli.Context) error {
	// ctx.JSON(ctx.RootArgv())
	// ctx.JSON(ctx.Argv())
	// fmt.Println()

	argv := ctx.Argv().(*rootT)
	Cascadia(argv.Filei, argv.Fileo, argv.CSS, argv.Piece, argv.Deli)
	argv.Filei.Close()
	argv.Fileo.Close()
	return nil
}

//--------------------------------------------------------------------------

// Cascadia filters the input buffer/stream `bi` with CSS selectors `css` and write to the output buffer/stream `bw`.
func Cascadia(bi io.Reader, bw io.Writer, css string, piece MapStringString, deli string) error {
	if len(piece.Values) == 0 {
		doc, err := html.Parse(bi)
		abortOn("Input", err)
		c := cascadia.MustCompile(css)

		// https://godoc.org/github.com/andybalholm/cascadia
		ns := c.MatchAll(doc)
		fmt.Fprintf(os.Stderr, "%d elements for '%s':\n", len(ns), css)
		for _, n := range ns {
			html.Render(bw, n)
			fmt.Fprintf(bw, "\n")
		}
	} else {
		// fmt.Printf("%v\n", piece)

		// https://godoc.org/github.com/PuerkitoBio/goquery
		// for debug
		//doc, err := goquery.NewDocumentFromReader(strings.NewReader(testhtml))
		doc, err := goquery.NewDocumentFromReader(bi)
		abortOn("Input", err)

		// Print csv headers
		for _, key := range piece.Keys {
			fmt.Fprintf(bw, "%s%s", key, deli)
		}
		fmt.Fprintf(bw, "\n")

		// Process each item block
		doc.Find(css).Each(func(index int, item *goquery.Selection) {
			//fmt.Printf("] #%d: %s\n", index, item.Text())
			for _, key := range piece.Keys {
				//fmt.Printf("] %s\n", piece.Values[key])
				fmt.Fprintf(bw, "%s%s",
					item.Find(piece.Values[key]).Contents().Text(), deli)
			}
			fmt.Fprintf(bw, "\n")
		})
	}
	return nil
}

//==========================================================================
// cli parameter handling

// DecodeSlice implements cli.SliceDecoder
// NOTE: if SliceDecoder not implemented, the Decode method would be only invoked once
func (MapStringString) DecodeSlice() {}

// Decode implements cli.Decoder interface
func (m *MapStringString) Decode(s string) error {
	if (m.Values) == nil {
		m.Values = make(map[string]string)
	}
	matches := regexp.MustCompile("(.*)=(.*)").FindStringSubmatch(s)
	key := matches[1]
	val := matches[2]
	m.Keys = append(m.Keys, key)
	m.Values[key] = val
	return nil
}

//==========================================================================
// support functions

// abortOn will quit on anticipated errors gracefully without stack trace
func abortOn(errCase string, e error) {
	if e != nil {
		fmt.Printf("[%s] %s error: %v\n", progname, errCase, e)
		os.Exit(1)
	}
}

//==========================================================================
// for debug
// echo a | cascadia -i -o -c "[align=\"justify\"]" -p Bold="b"

const testhtml = `
<div class="container">
    <div class="row">
      <div class="col-lg-8">
        <p align="justify"><b>Name</b>Priyaka</p>
        <p align="justify"><b>Surname</b>Patil</p>
        <p align="justify"><b>Adress</b><br>India,Kolhapur</p>
        <p align="justify"><b>Hobbies&nbsp;</b><br>Playing</p>
        <p align="justify"><b>Eduction</b><br>12th</p>
        <p align="justify"><b>School</b><br>New Highschool</p>
       </div>
    </div>
</div>
`
