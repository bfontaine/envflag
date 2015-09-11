// Package envflag provides a way to override environment variables with flags
package envflag

import (
	"errors"
	"flag"
	"os"
	"strings"
)

// MalformedValue is returned by Value.Set if the given key/value pair is not
// valid.
var MalformedValue = errors.New("Malformed environment key/value pair, expected <key>=<value>")

// Value implements the flag.Value interface and can be used as a target for
// flag.Var to override environment variables. It doesn’t save any value.
type Value struct{}

var _ flag.Value = &Value{}

// NewValue returns a pointer on a Value
func NewValue() *Value { return &Value{} }

// Set implements the flag.Value interface
func (v *Value) Set(s string) error {
	parts := strings.SplitN(s, "=", 2)
	if len(parts) != 2 {
		return MalformedValue
	}

	return os.Setenv(parts[0], parts[1])
}

func (v *Value) String() string {
	// we want to stay light and thus don't keep track of the flags.
	return "{env}"
}

// Setup creates a flag that will be used to override environment variables. It
// must be called before flag.Parse(). The flag can occur multiple times in the
// command-line.
func Setup(flagName, usage string) {
	flag.Var(&Value{}, flagName, usage)
}

// AutoSetup is a shortcut for Setup("e", "Override environment variables"). It
// must be called before flag.Parse().
func AutoSetup() {
	flag.Var(&Value{}, "e", "Override environment variables")
}
