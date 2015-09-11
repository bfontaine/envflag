package envflag_test

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/bfontaine/envflag"
	"github.com/bfontaine/vanish/vanish"
	"github.com/stretchr/testify/assert"
)

func ExampleSetup() {
	envflag.Setup("e", "Override env variables")
	flag.Parse()

	fmt.Printf("The key is '%s'\n", os.Getenv("KEY"))

	// This example will print:
	// - "foobar" if it’s called as ./example -e KEY=foobar regardless of the
	//   environment variable "KEY"
	// - "abc123" if it’s called as ./example AND the environment variable
	//   "KEY" is set to "abc123"
	// - an empty key if no -e flag is passed and the environment variable
	//   "KEY" is empty or doesn’t exist
}

func TestMultipleValues(t *testing.T) {
	vanish.Env(func() {
		fs := flag.NewFlagSet("f", flag.ContinueOnError)
		fs.Var(envflag.NewValue(), "yo", "tutu")

		os.Unsetenv("A")
		os.Unsetenv("B")
		os.Unsetenv("C")

		err := fs.Parse([]string{
			"-yo", "A=42", "-yo", "B=yo", "-yo", "C=true", "-yo", "A=43",
		})

		assert.Nil(t, err)
		assert.Equal(t, "43", os.Getenv("A"))
		assert.Equal(t, "yo", os.Getenv("B"))
		assert.Equal(t, "true", os.Getenv("C"))
	})
}

func TestComplexValue(t *testing.T) {
	vanish.Env(func() {
		fs := flag.NewFlagSet("f", flag.ContinueOnError)
		fs.Var(envflag.NewValue(), "x", "tutu")
		os.Unsetenv("A")

		assert.Nil(t, fs.Parse([]string{"-x", "A=42 B=17"}))

		assert.Equal(t, "42 B=17", os.Getenv("A"))
	})
}
