package main_test

import (
	"testing"
)

func TestPieceAttr(t *testing.T) {
	testCases(t, "Piece Attributes", []testCase{
		{"id",
			[]string{"-i", "opt_piece_attr.html", "-o", "-c", "li", "-p", "url=attr[id]:a"},
		},
		{"href",
			[]string{"-i", "opt_piece_attr.html", "-o", "-c", "li", "-p", "url=attr[href]:a"},
		},
		//{"both_fields",},
		//{"noexist"},
		//{},
	})
}
