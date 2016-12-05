package main

import (
	"github.com/mnishizawa/go-wufoo-api-client/wufoo"
	"github.com/mnishizawa/go-wufoo-api-client/wufoo/api"
	"log"
	"net/url"
)

var formHash string = "form_id" // set your form hash here

func main() {
	client := new(api.Client)
	client.Config = wufoo.Config{"API key", "subdomain"} // put your values

	//	GetOneForm(client)
	//	GetAllForms(client)
	//	GetFormFields(client)
	//	GetFormEntries(client)

//	postData := make(url.Values)
//	postData.Set("Field10", "Name")
//	postData.Set("Field11", "Second")
//	postData.Set("Field3", "check@itembase.api")
//	postData.Set("Field4", "01512534976")
//	postData.Set("Field118", "Test API wrapper")
//
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

func PostFormEntry(client *api.Client, postData url.Values) {
	response, err := client.EntriesApi().PostEntries(formHash, postData)
	if err != nil {
		log.Fatalf("Unable to get form from API: %+v", err)
	}

	log.Printf("Fetched form: %+v", response)
}
