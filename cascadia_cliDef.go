////////////////////////////////////////////////////////////////////////////
// Program: cascadiaC
// Purpose: cascadia wrapper
// Authors: Tong Sun (c) 2018, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"github.com/go-easygen/cli"
	clix "github.com/go-easygen/cli/ext"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

//==========================================================================
// cascadiaC

type rootT struct {
	cli.Helper
	Filei    *clix.Reader    `cli:"*i,in" usage:"The html/xml file to read from (or stdin)"`
	Fileo    *clix.Writer    `cli:"*o,out" usage:"The output file (or stdout)"`
	CSS      []string        `cli:"*c,css" usage:"CSS selectors (can provide more if not using --piece)"`
	Piece    MapStringString `cli:"p,piece" usage:"sub CSS selectors within -css to split that block up into pieces\n\t\t\tformat: PieceName=[RAW:]selector_string\n\t\t\tRAW: will return the selected as-is; else the text will be returned"`
	Deli     string          `cli:"d,delimiter" usage:"delimiter for pieces csv output" dft:"\t"`
	WrapHTML bool            `cli:"w,wrap-html" usage:"wrap up the output with html tags"`
	Base     string          `cli:"b,base" usage:"base href tag used in the wrapped up html"`
	Quiet    bool            `cli:"q,quiet" usage:"be quiet"`
}

var root = &cli.Command{
	Name: "cascadiaC",
	Desc: "cascadia wrapper\nVersion " + version + " built on " + date,
	Text: "Command line interface to go cascadia CSS selectors package" +
		"\n\nUsage:\n  cascadia -i in -c css -o [Options...]",
	Argv: func() interface{} { return new(rootT) },
	Fn:   cascadiaC,

	NumOption: cli.AtLeast(3),
}

// Template for main starts here
////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The OptsT type defines all the configurable options from cli.
//  type OptsT struct {
//  	Filei	*clix.Reader
//  	Fileo	*clix.Writer
//  	CSS	[]string
//  	Piece	MapStringString
//  	Deli	string
//  	WrapHTML	bool
//  	Base	string
//  	Quiet	bool
//  	Verbose int
//  }

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

//  var (
//          progname  = "cascadiaC"
//          version   = "0.1.0"
//          date = "2018-05-22"

//  	rootArgv *rootT
//  	// Opts store all the configurable options
//  	Opts OptsT
//  )

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
//  func main() {
//  	cli.SetUsageStyle(cli.ManualStyle) // up-down, for left-right, use NormalStyle
//  	//NOTE: You can set any writer implements io.Writer
//  	// default writer is os.Stdout
//  	if err := cli.Root(root,).Run(os.Args[1:]); err != nil {
//  		fmt.Fprintln(os.Stderr, err)
//  	}
//  	fmt.Println("")
//  }

// Template for main dispatcher starts here
//==========================================================================
// Main dispatcher

//  func cascadiaC(ctx *cli.Context) error {
//  	ctx.JSON(ctx.RootArgv())
//  	ctx.JSON(ctx.Argv())
//  	fmt.Println()

//  	return nil
//  }

// Template for CLI handling starts here
