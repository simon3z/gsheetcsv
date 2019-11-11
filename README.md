# Google Sheets CSV Tool

## Introduction

The Google Sheets CSV tool is an utility to import and export data to and from Google Sheets using the CSV format.

Examples:

    $ gsheetcsv dump --range A1:L100 --sheet [sheet-id]
    ...

    $ gsheetcsv update -range A1 --sheet [sheet-id] < values.csv


## Building

    go build -ldflags "-X main.oAuth2ClientID=[client-id] -X main.oAuth2ClientSecret=[client-secret]"

Substitute [client-id] and [client-secret] with the credentials prepared on the Google Project you created for this application.
