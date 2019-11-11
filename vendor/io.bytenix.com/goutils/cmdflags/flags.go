package cmdflags

import (
	"strings"
)

type (
	// ArrayFlag is used for command line flags with multiple values
	ArrayFlag []string
)

func (i *ArrayFlag) String() string {
	return strings.Join(*i, ", ")
}

// Set function adds a value to the array
func (i *ArrayFlag) Set(value string) error {
	*i = append(*i, value)
	return nil
}
