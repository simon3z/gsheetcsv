package sheets

import (
	"fmt"
	"strconv"
	"strings"
)

// Row represents a spreadsheet row number
type Row uint

// Column represents a spreadsheet column number
type Column uint

// Sheet represents a spreadsheet name
type Sheet string

// Cell represents a spreadsheet cell location
type Cell struct {
	Column Column
	Row    Row
}

// Range represents a spreadsheet range of cells
type Range struct {
	Sheet Sheet
	Start Cell
	End   Cell
}

// RowFromString returns a Row from the relevant string
func RowFromString(s string) (Row, error) {
	v, err := strconv.ParseUint(s, 10, 64)
	return Row(v), err
}

func (r Row) String() string {
	return strconv.FormatInt(int64(r), 10)
}

// IsNil returns if the Row is nil
func (r Row) IsNil() bool {
	return r == 0
}

// ColumnFromString returns a Column from the relevant string
func ColumnFromString(s string) (Column, error) {
	c := Column(0)

	for i := 0; i < len(s); i++ {
		if s[i] < 'A' || s[i] > 'Z' {
			return Column(0), fmt.Errorf("parsing \"%s\": invalid syntax", s)
		}
		c = c*26 + Column(s[i]+1-'A')
	}

	return c, nil
}

func (c Column) String() string {
	s := ""

	for v := uint(c); v > 0; v = (v - 1) / 26 {
		s = string(uint('A')+(v-1)%26) + s
	}

	return s
}

// IsNil returns if the Column is nil
func (c Column) IsNil() bool {
	return c == 0
}

func (s Sheet) String() string {
	return string(s)
}

// IsNil returns if the Sheet is nil
func (s Sheet) IsNil() bool {
	return s == ""
}

// CellFromString returns a Cell from the relevant string
func CellFromString(s string) (Cell, error) {
	r := strings.IndexFunc(s, func(r rune) bool { return r >= '0' && r <= '9' })

	col, err := ColumnFromString(s[:r])

	if err != nil {
		return Cell{}, err
	}

	row, err := RowFromString(s[r:])

	if err != nil {
		return Cell{}, err
	}

	return Cell{col, row}, nil
}

func (c Cell) String() string {
	return c.Column.String() + c.Row.String()
}

// IsNil returns if the Cell is nil
func (c Cell) IsNil() bool {
	return c.Column.IsNil() && c.Row.IsNil()
}

// RangeFromString returns a Range from the relevant string
func RangeFromString(s string) (Range, error) {
	p := strings.SplitN(s, "!", 2)

	sheet := Sheet("")

	if len(p) > 1 {
		sheet = Sheet(p[0])
	}

	p = strings.SplitN(p[len(p)-1], ":", 2)

	start, err := CellFromString(p[0])

	if err != nil {
		return Range{}, err
	}

	end := Cell{start.Column, start.Row}

	if len(p) > 1 {
		end, err = CellFromString(p[1])

		if err != nil {
			return Range{}, err
		}
	}

	return Range{sheet, start, end}, nil
}

func (r Range) String() string {
	s := ""

	if !r.Sheet.IsNil() {
		s += r.Sheet.String() + "!"
	}

	s += r.Start.String()

	if !r.End.IsNil() && r.End != r.Start {
		s += ":" + r.End.String()
	}

	return s
}

// IsNil returns if the Cell is nil
func (r Range) IsNil() bool {
	return r.Sheet.IsNil() && r.Start.IsNil() && r.End.IsNil()
}
