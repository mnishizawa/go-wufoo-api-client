package main

import (
	"github.com/itembase/go-wufoo-api-client/wufoo"
	"github.com/itembase/go-wufoo-api-client/wufoo/api"
	"log"
)

func main() {
	config := wufoo.Config{"<API_KEY>", "<subdomain>"}
	client := wufoo.Client{config}

//	GetOneForm(&client)
//	GetAllForms(&client)
}

func GetAllForms(client *wufoo.Client) {
	formApi := api.FormsApi{client}

	collection, err := formApi.Forms(false)
	if err != nil {
		log.Fatal("Unable to get forms from API")
	}

	log.Printf("Amount of forms fetched: %d", len(collection.Forms))
}

func GetOneForm(client *wufoo.Client) {
	formApi := api.FormsApi{client}

	form, err := formApi.FormsDetails("rwcd0mx003au1i", false)
	if err != nil {
		log.Fatal("Unable to get form from API")
	}

	log.Printf("Fetched form: %+v", form)
}
