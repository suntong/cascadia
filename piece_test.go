package main_test

import (
	"testing"
)

func TestPieceAttr(t *testing.T) {
	testCases(t, "Piece Attributes", []testCase{
		{"id",
			[]string{"-i", "opt_piece_attr.html", "-o", "-c", "li > a", "-p", "url=ATTR:id"},
		},
		{"href",
			[]string{"-i", "opt_piece_attr.html", "-o", "-c", "li > a", "-p", "url=ATTR:href"},
		},
		//{"both_fields",},
		//{"noexist"},
		//{},
	})
}
