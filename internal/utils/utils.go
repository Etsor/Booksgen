package utils

import (
    "os"
    "slices"
)

// HasArg checks if the specified argument is passed in command-line.
func HasArg(arg string) bool {
    return slices.Contains(os.Args, arg)
}
