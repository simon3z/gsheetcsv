package oauth2utils

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"golang.org/x/oauth2"
)

func NewOAuth2Client(config *oauth2.Config, tokenFile string) (*http.Client, error) {
	tok, err := oauth2TokenFromFile(tokenFile)

	if os.IsNotExist(err) {
		tok, err = oauth2TokenFromWeb(config)

		if err != nil {
			return nil, err
		}

		err = oauth2TokenSave(tokenFile, tok)

		if err != nil {
			return nil, err
		}
	}

	return config.Client(context.Background(), tok), nil
}

func oauth2TokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	cmd := exec.Command("xdg-open", authURL)
	err := cmd.Run()

	if err == nil {
		fmt.Printf("Follow the authorization procedure on your browser and type the authorization code: ")
	} else {
		fmt.Printf("Go to the following link in your browser then type the authorization code: \n%v\n", authURL)
	}

	var authCode string

	if _, err := fmt.Scan(&authCode); err != nil {
		return nil, err
	}

	token, err := config.Exchange(context.TODO(), authCode)

	if err != nil {
		return nil, err
	}

	return token, err
}

func oauth2TokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	token := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(token)

	return token, err
}

func oauth2TokenSave(path string, token *oauth2.Token) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)

	if err != nil {
		return err
	}

	defer f.Close()

	json.NewEncoder(f).Encode(token)

	return nil
}
