package conditionals_test

import (
	"testing"

	"github.com/spinozanose/go-hello-world/conditionals"
)

func TestConditionals(t *testing.T) {
	conditionals.IfStatements()
	if conditionals.GuessTheNumber(10) == true {
		t.Errorf("Expected false, got true")
	}
	if conditionals.GuessTheNumber(60) == true {
		t.Errorf("Expected false, got true")
	}
	if conditionals.GuessTheNumber(50) == false {
		t.Errorf("Expected true, got false")
	}
	conditionals.CompareFloatingPointNumbers()
}