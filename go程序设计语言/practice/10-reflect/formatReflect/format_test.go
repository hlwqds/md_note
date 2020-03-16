package format

import (
	"testing"
)

func TestAny(t *testing.T) {
	var testInt int = 1
	var wantInt string = "1"
	if got := Any(testInt); got != wantInt {
		t.Errorf("Any(%v) = %s, want %s", testInt, got, wantInt)
	}

	var testBool bool = true
	var wantBool string = "true"
	if got := Any(testBool); got != wantBool {
		t.Errorf("Any(%v) = %s, want %s", testBool, got, wantBool)
	}
}
