package variables_test

import (
	"testing"

	// Note that the namespace is the package name, no matter the file name
	"github.com/spinozanose/go-hello-world/variables"
)

func TestAssignments(t *testing.T) {
	variables.Assignments()
	variables.Composition()
	variables.Tags()
}