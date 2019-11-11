package main

import (
	"net/http"
	"os"
	"path"

	"golang.org/x/oauth2"
	"io.bytenix.com/goutils/cmdflags"
	"io.bytenix.com/goutils/oauth2utils"
	"io.bytenix.com/goutils/sysutils"
)

var (
	oAuth2ClientID     string
	oAuth2ClientSecret string

	oAuth2Config = oauth2.Config{
		ClientID:     oAuth2ClientID,
		ClientSecret: oAuth2ClientSecret,
		RedirectURL:  "urn:ietf:wg:oauth:2.0:oob",
		Scopes: []string{
			"https://www.googleapis.com/auth/spreadsheets",
		},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://oauth2.googleapis.com/token",
		},
	}
)

type (
	// GSheetCommand represents the gsheetcsv base command
	GSheetCommand struct {
		AppDataPath  string
		OAuth2Client *http.Client
	}
)

// Init initializes the GSheetCommand structure
func (c *GSheetCommand) Init() error {
	appDataPath, err := sysutils.InitAppDataDirectory("gsheetcsv")

	if err != nil {
		return err
	}

	c.AppDataPath = appDataPath

	oAuth2Client, err := oauth2utils.NewOAuth2Client(&oAuth2Config, path.Join(c.AppDataPath, "token.config"))

	if err != nil {
		return err
	}

	c.OAuth2Client = oAuth2Client

	return nil
}

func main() {
	commands := map[string]cmdflags.Command{
		"dump":   &DumpCommand{},
		"update": &UpdateCommand{},
	}

	if err := cmdflags.RunCommand(commands, os.Args[1:]); err != nil {
		panic(err)
	}
}
