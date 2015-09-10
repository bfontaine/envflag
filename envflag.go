package envflag

import (
	"errors"
	"flag"
	"os"
	"strings"
)

var MalformedValue = errors.New("Malformed environment key/value pair, expected <key>=<value>")

type Value []string

var _ flag.Value = &Value{}

// Set implements the flag.Value interface
func (v *Value) Set(s string) error {
	parts := strings.Split(s, "=")
	if len(parts) != 2 {
		return MalformedValue
	}

	*v = append(*v, s)

	return os.Setenv(parts[0], parts[1])
}

func (v *Value) String() string {
	return strings.Join(*v, ",")
}
