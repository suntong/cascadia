////////////////////////////////////////////////////////////////////////////
// Program: cascadia
// Purpose: go cascadia CSS selection from command line
// Authors: Tong Sun (c) 2016-2021, All rights reserved
////////////////////////////////////////////////////////////////////////////

//go:generate sh -v cascadia_cliGen.sh

package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/andybalholm/cascadia"
	"github.com/mkideal/cli"
	"golang.org/x/net/html"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

const (
	IsRaw       = "RAW"
	WrapHTMLEnd = `</body>`
)

type OutputStyle int

const (
	OutputStyleRAW OutputStyle = iota
	OutputStyleATTR
	OutputStyleTEXT
)

type MapStringString struct {
	Keys         []string
	Values       map[string]string
	OutputStyles map[string]OutputStyle
	AttrName     map[string]string
}

// The OptsT type defines all the configurable options from cli.
type OptsT struct {
	CSS      []string
	TextOut  bool
	TextRaw  bool
	Piece    MapStringString
	Deli     string
	WrapHTML bool
	Style    string
	Base     string
	Quiet    bool
	Verbose  int
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	progname = "cascadia"
	version  = "1.2.6"
	date     = "2021-11-28"

	rootArgv *rootT
	// Opts store all the configurable options
	Opts OptsT
)

var WrapHTMLBeg string

////////////////////////////////////////////////////////////////////////////
// Function definitions

func main() {
	cli.SetUsageStyle(cli.DenseNormalStyle)
	if err := cli.Root(root).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println("")
}

//==========================================================================
// css selection

func CascadiaC(ctx *cli.Context) error {
	// ctx.JSON(ctx.RootArgv())
	// fmt.Println()
	// ctx.JSON(ctx.Argv())
	// fmt.Println()

	argv := ctx.Argv().(*rootT)
	WrapHTMLBeg = fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<base href="%s">
%s
</head>
<body>`, argv.Base, argv.Style)

	Opts.CSS, Opts.Piece, Opts.Deli,
		Opts.WrapHTML, Opts.TextOut, Opts.TextRaw, Opts.Quiet =
		argv.CSS, argv.Piece, argv.Deli,
		argv.WrapHTML, argv.TextOut, argv.TextRaw, argv.Quiet
	Cascadia(argv.Filei, argv.Fileo, Opts)
	argv.Filei.Close()
	argv.Fileo.Close()
	return nil
}

//--------------------------------------------------------------------------

// Cascadia filters the input buffer/stream `bi` with CSS selectors array `Opts.CSS` and write to the output buffer/stream `bw`.
func Cascadia(bi io.Reader, bw io.Writer, Opts OptsT) error {
	cssa, piece, deli, wrapHTML, textOut, textRaw, beQuiet :=
		Opts.CSS, Opts.Piece, Opts.Deli,
		Opts.WrapHTML, Opts.TextOut, Opts.TextRaw, Opts.Quiet
	if wrapHTML {
		fmt.Fprintln(bw, WrapHTMLBeg)
	}
	if len(piece.Values) == 0 {
		// no sub CSS selectors -- none-block selection mode
		if textOut {
			doc, err := goquery.NewDocumentFromReader(bi)
			abortOn("Input", err)

			for _, css := range cssa {
				// Process each item block
				doc.Find(css).Each(func(index int, item *goquery.Selection) {
					//fmt.Printf("] #%d: %s\n", index, item.Text())
					if textRaw {
						fmt.Fprintf(bw, "%s%s",
							item.Text(), deli)
					} else {
						fmt.Fprintf(bw, "%s%s",
							strings.TrimSpace(item.Text()), deli)
					}
					fmt.Fprintf(bw, "\n")
				})
			}
		} else {
			doc, err := html.Parse(bi)
			abortOn("Input", err)
			for _, css := range cssa {
				c, err := cascadia.Compile(css)
				abortOn("CSS Selector string "+css, err)

				// https://godoc.org/github.com/andybalholm/cascadia
				ns := c.MatchAll(doc)
				if !beQuiet {
					fmt.Fprintf(os.Stderr, "%d elements for '%s':\n", len(ns), css)
				}
				for _, n := range ns {
					html.Render(bw, n)
					fmt.Fprintf(bw, "\n")
				}
			}
		}
	} else {
		// have sub CSS selectors within -css -- block selection mode
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
		doc.Find(cssa[0]).Each(func(index int, item *goquery.Selection) {
			//fmt.Printf("] #%d: %s\n", index, item.Text())
			for _, key := range piece.Keys {
				//fmt.Printf("] %s: %s\n", key, piece.Values[key])
				switch piece.OutputStyles[key] {
				case OutputStyleRAW:
					html.Render(bw, item.Find(piece.Values[key]).Get(0))
					fmt.Fprintf(bw, deli)
				case OutputStyleATTR:
					fmt.Fprintf(bw, "%s%s",
						item.Find(piece.Values[key]).AttrOr(piece.AttrName[key], ""), deli)
				case OutputStyleTEXT:
					fmt.Fprintf(bw, "%s%s",
						item.Find(piece.Values[key]).Contents().Text(), deli)
				}
			}
			fmt.Fprintf(bw, "\n")
		})
	}
	if wrapHTML {
		fmt.Fprintln(bw, WrapHTMLEnd)
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
		m.OutputStyles = make(map[string]OutputStyle)
		m.AttrName = make(map[string]string)
	}
	matches := regexp.MustCompile("(.*)=(.*)").FindStringSubmatch(s)
	if len(matches) < 2 {
		return errors.New("format error. To get help, run: " + progname)
	}
	key := matches[1]
	val := matches[2]
	index := strings.Index(val, ":")
	if index > 0 {
		style := val[:index]
		val = val[index+1:]
		if style == IsRaw {
			m.OutputStyles[key] = OutputStyleRAW
		} else if strings.HasPrefix(style, "attr[") && strings.HasSuffix(style, "]") {
			m.OutputStyles[key] = OutputStyleATTR
			m.AttrName[key] = style[5 : len(style)-1]
		} else {
			m.OutputStyles[key] = OutputStyleTEXT
		}
	} else {
		m.OutputStyles[key] = OutputStyleTEXT
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
