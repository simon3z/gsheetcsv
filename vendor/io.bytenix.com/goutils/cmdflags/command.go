package cmdflags

import (
	"flag"
	"fmt"
)

type (
	// Command is the interface of a command
	Command interface {
		Init(*flag.FlagSet) error
		Execute() error
	}

	CommandFlags struct {
	}
)

// RunCommands parse flags and runs the relevant command
func RunCommand(flags map[string]Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("command not provided")
	}

	cmd := args[0]

	if c, ok := flags[cmd]; ok {
		f := flag.NewFlagSet(cmd, flag.ExitOnError)

		if err := c.Init(f); err != nil {
			return err
		}

		f.Parse(args[1:])

		return c.Execute()
	}

	return fmt.Errorf("command not found: %s", cmd)
}
