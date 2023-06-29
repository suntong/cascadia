package main_test

import (
	"testing"
)

func TestPiece(t *testing.T) {
	testCases(t, "Piece raw html", []testCase{
		{"piece_raw_a", []string{
			"-i", "opt_piece_attr.html", "-o", "-c", "ul > li", "-p", "Source=RAW:a",
		}},
	})
	testCases(t, "Piece Attributes", []testCase{
		{"id", []string{
			"-i", "opt_piece_attr.html", "-o", "-c", "li > a", "-p", "id=ATTR:id",
		}},
		{"href", []string{
			"-i", "opt_piece_attr.html", "-o", "-c", "li > a", "-p", "url=ATTR:href",
		}},
		{"both_fields", []string{"-i", "opt_piece_attr.html", "-o", "-c", "li > a",
			"-p", "id=ATTR:id", "-p", "url=ATTR:href",
		}},

		{"noexist", []string{
			"-i", "opt_piece_attr.html", "-o", "-c", "li > a", "-p", "NoExist=ATTR:noexist",
		}},

		{"script_src", []string{
			"-i", "opt_piece_script.html", "-o", "-c", "html > head > script", "-p", "SourceJS=ATTR:src",
		}},
		//{},
	})
}
