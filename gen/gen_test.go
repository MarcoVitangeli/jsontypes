package gen

import (
	"testing"
)

var (
	jsonArrayRoot = `[
	{
		"foo": "bar"
	},
	{
		"foo2": "bar2"
	}
]`
)

func TestJsonValidationArrayAsRoot(t *testing.T) {
	err := Gen([]byte(jsonArrayRoot))
	if err == nil {
		t.Fatalf("no error when inserting invalid json array as root")
	}
}

func TestJsonValidation(t *testing.T) {
	var invalidJson = `asdasdasd
	asdasdasd`

	err := Gen([]byte(invalidJson))
	if err == nil {
		t.Fatalf("no error when inserting invalid json")
	}
}
