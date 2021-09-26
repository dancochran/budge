package string

import (
	"testing"
)

// TestStringReverse calls string.StringReverse with a string, checking for a
// valid return value.
func TestStringReverse(t *testing.T) {
	inputStr := "Mountain"
	wantStr := "niatnuoM"
	reversedStr := StringReverse(inputStr)
	if reversedStr != wantStr {
		t.Fatalf(`StringReverse("Mountain") = %q, want match for %#q`, reversedStr, wantStr)
	}
}
