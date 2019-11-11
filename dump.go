package main

import (
	"encoding/csv"
	"flag"
	"os"

	"io.bytenix.com/gsheetcsv/sheets"
)

type (
	// DumpCommand represents a spreadsheet dump command
	DumpCommand struct {
		GSheetCommand
		Flags struct {
			Sheet   string
			Range   string
			TabSep  bool
			Formula bool
		}
	}
)

// Init initializes the DumpCommand structure and flags
func (c *DumpCommand) Init(f *flag.FlagSet) error {
	if err := c.GSheetCommand.Init(); err != nil {
		return err
	}

	f.StringVar(&c.Flags.Sheet, "sheet", "", "Google Sheet ID")
	f.StringVar(&c.Flags.Range, "range", "", "Google Sheet Range")
	f.BoolVar(&c.Flags.TabSep, "tab", false, "Use tab as fields delimiter")
	f.BoolVar(&c.Flags.Formula, "formula", false, "Print formulas")

	return nil
}

// Execute executes the DumpCommand business logic
func (c *DumpCommand) Execute() error {
	rng, err := sheets.RangeFromString(c.Flags.Range)

	if err != nil {
		return err
	}

	values, err := GetSpreadsheetValues(c.OAuth2Client, c.Flags.Sheet, rng, c.Flags.Formula)

	if err != nil {
		return err
	}

	w := csv.NewWriter(os.Stdout)

	if c.Flags.TabSep {
		w.Comma = '\t'
	}

	for _, v := range values {
		w.Write(v)
	}

	w.Flush()

	return nil
}
