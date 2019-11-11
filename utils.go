package main

import (
	"fmt"
	"net/http"
	"strconv"

	gsheets "google.golang.org/api/sheets/v4"
	"io.bytenix.com/gsheetcsv/sheets"
)

func GetSpreadsheetValues(c *http.Client, id string, r sheets.Range, formula bool) ([][]string, error) {
	s, err := gsheets.New(c)

	if err != nil {
		return [][]string{}, err
	}

	var renderOption string

	if formula {
		renderOption = "FORMULA"
	} else {
		renderOption = "FORMATTED_VALUE"
	}

	cells, err := s.Spreadsheets.Values.Get(id, r.String()).ValueRenderOption(renderOption).DateTimeRenderOption("FORMATTED_STRING").Do()

	if err != nil {
		return [][]string{}, err
	}

	values := make([][]string, len(cells.Values))

	for i, v := range cells.Values {
		values[i], err = GetValueNStrings(v, int(r.End.Column-r.Start.Column+1))

		if err != nil {
			return nil, err
		}
	}

	return values, nil
}

func GetValueNStrings(v []interface{}, l int) ([]string, error) {
	s := make([]string, l)

	for i := range v {
		switch v[i].(type) {
		case string:
			s[i] = v[i].(string)
		case float64:
			s[i] = strconv.FormatFloat(v[i].(float64), 'f', -1, 64)
		case bool:
			if v[i].(bool) {
				s[i] = "TRUE"
			} else {
				s[i] = "FALSE"
			}
		default:
			return nil, fmt.Errorf("Cannot convert: %s", v[i])
		}
	}

	return s, nil
}
