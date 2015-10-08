package main

import (
	"github.com/itembase/go-wufoo-api-client/wufoo"
	"github.com/itembase/go-wufoo-api-client/wufoo/api"
	"log"
)

var formHash string = "hash"

func main() {
	config := wufoo.Config{"API key", "subdomain"}

	client := new(api.Client)
	client.Config = config

//	GetOneForm(client)
//	GetAllForms(client)
//	GetFormFields(client)
//	GetFormEntries(client)

//	postData := make(map[string]string)
//	postData = map[string]string{
//		"Field10":  "Name",
//		"Field11":  "Second",
//		"Field3":   "check@itembase.api",
//		"Field4":   "01512534976",
//		"Field118": "Test API wrapper",
//	}
//	PostFormEntry(client, postData)
}

func GetAllForms(client *api.Client) {
	collection, err := client.FormsApi().Forms(false)
	if err != nil {
		log.Fatalf("Unable to get form from API: %+v", err)
	}

	log.Printf("Amount of forms fetched: %d", len(collection.Forms))
}

func GetOneForm(client *api.Client) {
	form, err := client.FormsApi().FormsDetails(formHash, false)
	if err != nil {
		log.Fatalf("Unable to get form from API: %+v", err)
	}

	log.Printf("Fetched form: %+v", form)
}

func GetFormFields(client *api.Client) {
	fields, err := client.FieldsApi().Fields(formHash, false)
	if err != nil {
		log.Fatalf("Unable to get form from API: %+v", err)
	}

	log.Printf("Fetched form: %v", fields)
}

func GetFormEntries(client *api.Client) {
	fields, err := client.EntriesApi().Entries(formHash, 1, 10)
	if err != nil {
		log.Fatalf("Unable to get form from API: %+v", err)
	}

	log.Printf("Fetched form: %v", fields)
}

func PostFormEntry(client *api.Client, postData map[string]string) {
	response, err := client.EntriesApi().PostEntries(formHash, postData)
	if err != nil {
		log.Fatalf("Unable to get form from API: %+v", err)
	}

	log.Printf("Fetched form: %+v", response)
}
