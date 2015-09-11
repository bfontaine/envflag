# envflag

**enflag** is a lightweight Go library to allow environment variables
overriding with command-line flags.

## Install

    go get github.com/bfontaine/envflag

## Example

This example is available in `example/main.go`:

```go
package main

import (
    "flag"
    "fmt"
    "os"

    "github.com/bfontaine/envflag"
)

func main() {
    envflag.AutoSetup()
    flag.Parse()

    fmt.Printf("The key is '%s'\n", os.Getenv("KEY"))
}
```

Once compiled, here is how to use it:

```
$ ./example -e KEY=42
The key is '42'

$ export KEY=43
$ ./example
The key is '43'

$ export KEY=43
$ ./example -e KEY=41
The key is '41'

$ ./example -e KEY=41 -e KEY=abcd
The key is 'abcd'
```
