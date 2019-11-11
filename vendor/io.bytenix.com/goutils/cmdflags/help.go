package cmdflags

import (
	"flag"
)

type (
	// HelpCommand lore ipsum
	HelpCommand struct {
		Commands map[string]Command
	}
)

// Init lore ipsum
func (c *HelpCommand) Init(f *flag.FlagSet) error {
	return nil
}

// Execute lore ipsum
func (c *HelpCommand) Execute() error {
	/*k := make([]string, len(c.Commands))

	for i, v := range c.Commands {
		k[0] = i
	}

	sort.Strings(k)

	for _, v := range k {

	}*/

	return nil
}
