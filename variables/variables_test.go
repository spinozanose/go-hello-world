package variables_test

import (
	"testing"

	"github.com/spinozanose/go-hello-world/variables"
)

func TestVariables(t *testing.T) {
	variables.DeclareVariables()
	variables.DeclareTypes()
	variables.Types()
	variables.Arrays()
	variables.Slices()
}

func TestMaps(t *testing.T) {
	variables.Maps()
}
