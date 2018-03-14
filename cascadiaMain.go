////////////////////////////////////////////////////////////////////////////
// Program: cascadia
// Purpose: go cascadia CSS selection from command line
// Authors: Tong Sun (c) 2016-2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

//go:generate sh -v cascadiaCLIGen.sh

package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/andybalholm/cascadia"
	"github.com/mkideal/cli"
	"golang.org/x/net/html"
	"strings"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

const (
	IsRaw       = "RAW:"
	WrapHTMLEnd = `</body>`
)

type MapStringString struct {
	Keys   []string
	Values map[string]string
	Raw    map[string]bool
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var progname = "cascadia"
var buildTime = "2017-04-17"

var rootArgv *rootT
var WrapHTMLBeg string

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
	WrapHTMLBeg = fmt.Sprintf(`<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <base href="%s">
  </head>
<body>`, argv.Base)

	Cascadia(
		argv.Filei,
		argv.Fileo,
		argv.CSS,
		argv.Piece,
		argv.NoHeader,
		argv.Deli,
		argv.WrapHTML,
		argv.RawAttr,
		argv.Text,
		argv.RawText,
		argv.Quiet,
	)
	argv.Filei.Close()
	argv.Fileo.Close()
	return nil
}

//--------------------------------------------------------------------------

// Cascadia filters the input buffer/stream `bi` with CSS selectors `css` and write to the output buffer/stream `bw`.
func Cascadia(bi io.Reader, bw io.Writer, css string, piece MapStringString, noHeader bool, deli string, wrapHTML bool, rawAttr string, text bool, rawText bool, beQuiet bool) error {
	if len(piece.Values) == 0 {
		doc, err := html.Parse(bi)
		abortOn("Input", err)
		c, err := cascadia.Compile(css)
		abortOn("CSS Selector string "+css, err)

		// https://godoc.org/github.com/andybalholm/cascadia
		ns := c.MatchAll(doc)
		if !beQuiet {
			fmt.Fprintf(os.Stderr, "%d elements for '%s':\n", len(ns), css)
		}

		if rawAttr != "" {
			if len(ns) != 1 {
				return fmt.Errorf("requested --raw-attr found more than 1 element")
			}

			for _, attr := range ns[0].Attr {
				if attr.Key == rawAttr {
					fmt.Fprintf(bw, attr.Val)
				}
			}
			return nil
		}

		if text == true {
			if len(ns) != 1 {
				return fmt.Errorf("requested --text but found more than 1 element")
			}

			fmt.Fprintf(bw, strings.Join(strings.Fields(getInnerText(ns[0])), " ")+"\n")
			return nil
		}

		if rawText == true {
			if len(ns) != 1 {
				return fmt.Errorf("requested --raw-text but found more than 1 element")
			}

			fmt.Fprintf(bw, getInnerText(ns[0])+"\n")
			return nil
		}

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

		if wrapHTML {
			fmt.Fprintln(bw, WrapHTMLBeg)
		}
		if noHeader == false {
			// Print csv headers
			for _, key := range piece.Keys {
				fmt.Fprintf(bw, "%s%s", key, deli)
			}
			fmt.Fprintf(bw, "\n")
		}

		// Process each item block
		doc.Find(css).Each(func(index int, item *goquery.Selection) {
			//fmt.Printf("] #%d: %s\n", index, item.Text())
			for _, key := range piece.Keys {
				//fmt.Printf("] %s: %s\n", key, piece.Values[key])
				if piece.Raw[key] {
					html.Render(bw, item.Find(piece.Values[key]).Get(0))
					fmt.Fprintf(bw, deli)
				} else {
					fmt.Fprintf(bw, "%s%s",
						item.Find(piece.Values[key]).Contents().Text(), deli)
				}
			}
			fmt.Fprintf(bw, "\n")
		})
		if wrapHTML {
			fmt.Fprintln(bw, WrapHTMLEnd)
		}
	}
	return nil
}

func getInnerText(node *html.Node) string {
	text := ""
	if node.Type == html.TextNode {
		text += node.Data
	} else if node.Type == html.ElementNode && node.FirstChild != nil {
		text += getInnerText(node.FirstChild)
	}
	if node.NextSibling != nil {
		text += getInnerText(node.NextSibling)
	}

	return text
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
		m.Raw = make(map[string]bool)
	}
	matches := regexp.MustCompile("(.*)=(.*)").FindStringSubmatch(s)
	if len(matches) < 2 {
		return errors.New("format error. To get help, run: " + progname)
	}
	key := matches[1]
	val := matches[2]
	if len(val) >= 4 && val[:4] == IsRaw {
		m.Raw[key] = true
		val = val[4:]
	}
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
