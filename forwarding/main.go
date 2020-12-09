package main

import (
	"fmt"
	"io/ioutil"
	"log"

	oauth "example.com/forwarding/oauth"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

const credentialFileName string = "client_id.json"

func main() {
	b, err := ioutil.ReadFile(credentialFileName)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope, gmail.GmailSettingsSharingScope, gmail.GmailLabelsScope, gmail.GmailSettingsBasicScope, gmail.GmailSendScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := oauth.GetClient(config)

	gmailService, err := gmail.New(client)
	if err != nil {
		log.Fatalf("Unable to create gmail Client %v", err)
	}
	fmt.Printf("%v", gmailService)
}
