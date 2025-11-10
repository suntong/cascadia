////////////////////////////////////////////////////////////////////////////
// Program: cascadia
// Purpose: go cascadia CSS selection from command line
// Authors: Tong Sun (c) 2016-2025, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"errors"
	"regexp"
	"strings"
)

type PieceStyle int

const (
	PieceStyleTEXT PieceStyle = iota
	PieceStyleRAW
	PieceStyleATTR
	PieceStyleGOQR
)

type PieceStyleMap struct {
	Keys          []string
	Values        map[string]string
	PieceStyles   map[string]PieceStyle
	GoqrSelectors map[string]string
	GoqrAttrs     map[string]string
}

var pieceStyles = map[string]PieceStyle{
	"RAW":  PieceStyleRAW,
	"ATTR": PieceStyleATTR,
	"GOQR": PieceStyleGOQR,
}

//==========================================================================
// cli parameter handling

// DecodeSlice implements cli.SliceDecoder
// NOTE: if SliceDecoder not implemented, the Decode method would be only invoked once
func (PieceStyleMap) DecodeSlice() {}

// Decode implements cli.Decoder interface
func (m *PieceStyleMap) Decode(s string) error {
	if m.Values == nil {
		m.Values = make(map[string]string)
		m.PieceStyles = make(map[string]PieceStyle)
		m.GoqrSelectors = make(map[string]string)
		m.GoqrAttrs = make(map[string]string)
	}
	matches := regexp.MustCompile("(.*?)=((.*?):)?(.*)").FindStringSubmatch(s)
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

	if style == PieceStyleGOQR {
		// selector.attr(attributeName)
		parts := strings.Split(val, ".attr(")
		if len(parts) != 2 || !strings.HasSuffix(parts[1], ")") {
			return errors.New("GOQR format error. To get help, run: " + progname)
		}
		selector := parts[0]
		attr := strings.TrimSuffix(parts[1], ")")
		m.GoqrSelectors[key] = selector
		m.GoqrAttrs[key] = attr
	}

	m.Keys = append(m.Keys, key)
	m.PieceStyles[key] = style
	m.Values[key] = val
	return nil
}
