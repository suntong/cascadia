////////////////////////////////////////////////////////////////////////////
// Program: cascadiaC
// Purpose: cascadia wrapper
// Authors: Tong Sun (c) 2016-2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"github.com/mkideal/cli"
	clix "github.com/mkideal/cli/ext"
)

////////////////////////////////////////////////////////////////////////////
// cascadiaC

type rootT struct {
	cli.Helper
	Filei    *clix.Reader    `cli:"*i,in" usage:"The html/xml file to read from (or stdin)"`
	Fileo    *clix.Writer    `cli:"*o,out" usage:"The output file (or stdout)"`
	CSS      string          `cli:"*c,css" usage:"CSS selectors"`
	Piece    MapStringString `cli:"p,piece" usage:"sub CSS selectors within -css to split that block up into pieces\n\t\t\tformat: PieceName=[RAW:]selector_string\n\t\t\tRAW: will return the selected as-is; else the text will be returned"`
	NoHeader bool            `cli:"no-header" usage:"ignore headers when using --piece" dft:"false"`
	Deli     string          `cli:"d,delimiter" usage:"delimiter for pieces csv output" dft:"\t"`
	WrapHTML bool            `cli:"w,wrap-html" usage:"wrap up the output with html tags"`
	Base     string          `cli:"b,base" usage:"base href tag used in the wrapped up html"`
	RawAttr  string          `cli:"a,attr,raw-attr" usage:"extract an attribute from selected element. Ignored when --piece is used"`
	Text     bool            `cli:"t,text" usage:"extract the human readable text from selected element. Ignored when --piece is used"`
	RawText  bool            `cli:"r,raw-text" usage:"extract the innerText from selected element. Ignored when --piece is used"`
	Quiet    bool            `cli:"q,quiet" usage:"be quiet"`
}

var root = &cli.Command{
	Name: "cascadiaC",
	Desc: "cascadia wrapper\nbuilt on " + buildTime,
	Text: "Command line interface to go cascadia CSS selectors package",
	Argv: func() interface{} { return new(rootT) },
	Fn:   cascadiaC,

	NumOption: cli.AtLeast(3),
}

// func main() {
// 	cli.SetUsageStyle(cli.ManualStyle) // up-down, for left-right, use NormalStyle
// 	//NOTE: You can set any writer implements io.Writer
// 	// default writer is os.Stdout
// 	if err := cli.Root(root,).Run(os.Args[1:]); err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 	}
// 	fmt.Println("")
// }

// func cascadiaC(ctx *cli.Context) error {
// 	ctx.JSON(ctx.RootArgv())
// 	ctx.JSON(ctx.Argv())
// 	fmt.Println()

// 	return nil
// }
