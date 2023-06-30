////////////////////////////////////////////////////////////////////////////
// Program: cascadia
// Purpose: go cascadia CSS selection from command line
// Authors: Tong Sun (c) 2016-2023, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"errors"
	"regexp"
)

type PieceStyle int

const (
	PieceStyleTEXT PieceStyle = iota
	PieceStyleRAW
	PieceStyleATTR
)

type PieceStyleMap struct {
	Keys        []string
	Values      map[string]string
	PieceStyles map[string]PieceStyle
}

var pieceStyles = map[string]PieceStyle{
	"RAW":  PieceStyleRAW,
	"ATTR": PieceStyleATTR,
}

//==========================================================================
// cli parameter handling

// DecodeSlice implements cli.SliceDecoder
// NOTE: if SliceDecoder not implemented, the Decode method would be only invoked once
func (PieceStyleMap) DecodeSlice() {}

// Decode implements cli.Decoder interface
func (m *PieceStyleMap) Decode(s string) error {
	if (m.Values) == nil {
		m.Values = make(map[string]string)
		m.PieceStyles = make(map[string]PieceStyle)
	}
	matches := regexp.MustCompile("(.*)=((.*?):)?(.*)").FindStringSubmatch(s)
	if len(matches) < 4 {
		return errors.New("format error. To get help, run: " + progname)
	}
	key := matches[1]
	ptp := matches[3] // piece type
	val := matches[4]
	style := PieceStyleTEXT
	style, ok := pieceStyles[ptp]
	//fmt.Println("]", key, ptp, style, ok, val)
	if len(ptp) != 0 && !ok {
		return errors.New("Piece style specification error. To get help, run: " + progname)
	}
	m.Keys = append(m.Keys, key)
	m.PieceStyles[key] = style
	m.Values[key] = val
	return nil
}
