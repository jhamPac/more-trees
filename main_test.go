package main

import (
	"fmt"
	"testing"
)

type testcase struct {
	className      string
	node           Node
	expectedResult bool
}

func TestHasClass(t *testing.T) {
	cases := []testcase{
		testcase{
			className:      "foo",
			node:           Node{classes: "foo bar"},
			expectedResult: true,
		},
		testcase{
			className:      "foo",
			node:           Node{classes: "bar que ok"},
			expectedResult: false,
		},
		testcase{
			className:      "bar",
			node:           Node{classes: ""},
			expectedResult: false,
		},
	}

	for i, test := range cases {
		result := test.node.hasClass(test.className)
		t.Run(fmt.Sprintf("testing for case number %d", i+1), func(t *testing.T) {
			if result != test.expectedResult {
				t.Error("wrong")
			}
		})
	}
}
