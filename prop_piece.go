////////////////////////////////////////////////////////////////////////////
// Program: cascadia
// Purpose: go cascadia CSS selection from command line
// Authors: Tong Sun (c) 2016-2023, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"errors"
	"regexp"
	"strings"
)

type PieceStyle int

const (
	PieceStyleRAW PieceStyle = iota
	PieceStyleATTR
	PieceStyleTEXT
)

type PieceStyleMap struct {
	Keys        []string
	Values      map[string]string
	PieceStyles map[string]PieceStyle
	AttrName    map[string]string
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
			m.PieceStyles[key] = PieceStyleRAW
		} else if strings.HasPrefix(style, "attr[") && strings.HasSuffix(style, "]") {
			m.PieceStyles[key] = PieceStyleATTR
			m.AttrName[key] = style[5 : len(style)-1]
		} else {
			m.PieceStyles[key] = PieceStyleTEXT
		}
	} else {
		m.PieceStyles[key] = PieceStyleTEXT
	}
	m.Keys = append(m.Keys, key)
	m.Values[key] = val
	return nil
}
