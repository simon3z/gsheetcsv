package main

import (
	"encoding/csv"
	"flag"
	"os"

	gsheets "google.golang.org/api/sheets/v4"
	"io.bytenix.com/gsheetcsv/sheets"
)

type (
	// UpdateCommand represents a spreadsheet update command
	UpdateCommand struct {
		GSheetCommand
		Flags struct {
			Sheet  string
			Range  string
			TabSep bool
		}
	}
)

// Init initializes the UpdateCommand structure and flags
func (c *UpdateCommand) Init(f *flag.FlagSet) error {
	if err := c.GSheetCommand.Init(); err != nil {
		return err
	}

	f.StringVar(&c.Flags.Sheet, "sheet", "", "Google Sheet ID")
	f.StringVar(&c.Flags.Range, "range", "", "Google Sheet Range")
	f.BoolVar(&c.Flags.TabSep, "tab", false, "Use tab as fields delimiter")

	return nil
}

// Execute executes the UpdateCommand business logic
func (c *UpdateCommand) Execute() error {
	rng, err := sheets.RangeFromString(c.Flags.Range)

	if err != nil {
		return err
	}

	r := csv.NewReader(os.Stdin)

	if c.Flags.TabSep {
		r.Comma = '\t'
	}

	values, err := r.ReadAll()

	if err != nil {
		return err
	}

	rng.End.Row = rng.Start.Row + sheets.Row(len(values))
	rng.End.Column = rng.Start.Column + sheets.Column(len(values[0])-1)

	resp, err := GetSpreadsheetValues(c.OAuth2Client, c.Flags.Sheet, rng, true)

	updates := []*gsheets.ValueRange{}

	for i := range values {
		for j := range values[i] {
			if values[i][j] != resp[i][j] {
				updates = append(updates, &gsheets.ValueRange{
					Range: sheets.Range{
						Sheet: rng.Sheet,
						Start: sheets.Cell{
							Column: rng.Start.Column + sheets.Column(j),
							Row:    rng.Start.Row + sheets.Row(i),
						},
					}.String(),
					Values: [][]interface{}{
						[]interface{}{values[i][j]},
					},
				})
			}
		}
	}

	s, err := gsheets.New(c.OAuth2Client)

	if err != nil {
		return err
	}

	_, err = s.Spreadsheets.Values.BatchUpdate(c.Flags.Sheet, &gsheets.BatchUpdateValuesRequest{
		ValueInputOption: "USER_ENTERED",
		Data:             updates,
	}).Do()

	if err != nil {
		return err
	}

	return nil
}
