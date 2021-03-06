package ast

import (
	"testing"
)

func TestUnusedAttr(t *testing.T) {
	nodes := map[string]Node{
		`0x7fe3e01416d0 <col:47> unused`: &UnusedAttr{
			Address:  "0x7fe3e01416d0",
			Position: "col:47",
			Children: []Node{},
		},
	}

	runNodeTests(t, nodes)
}
