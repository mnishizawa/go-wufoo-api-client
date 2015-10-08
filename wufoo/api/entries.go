package api

import "strconv"

// FieldCollection structure represents Wufoo Field API response
type EntriesCollection struct {
	Entries []map[string]interface{}
}

// FieldsApi struct represents access to Wufoo Field API
type EntriesApi struct {
	Client *Client
}

// PostEntrieResponse represents response after creating a new entry
type PostEntrieResponse struct {
	Success     int
	EntryId     int
	EntryLink   string
	ErrorText   string
	FieldErrors []struct {
		ID        int
		ErrorText string
	}
}

// Fields method returns form fields details.
//
// For more details please visit: http://help.wufoo.com/articles/en_US/SurveyMonkeyArticleType/The-Entries-POST-API
//
func (api EntriesApi) Entries(formIdentifier string, page int, perPage int) (collection *EntriesCollection, err error) {
	collection, err = api.request("forms/"+formIdentifier+"/entries", page, perPage)
	return
}

// FieldsReport method returns form fields details.
//
// For more details please visit: http://help.wufoo.com/articles/en_US/SurveyMonkeyArticleType/The-Entries-POST-API
//
func (api EntriesApi) EntriesReport(formIdentifier string, page int, perPage int) (collection *EntriesCollection, err error) {
	collection, err = api.request("reports/"+formIdentifier+"/entries", page, perPage)
	return
}

// PostEntries method returns form fields details.
//
// For more details please visit: http://help.wufoo.com/articles/en_US/SurveyMonkeyArticleType/The-Entries-POST-API
//
func (api EntriesApi) PostEntries(formIdentifier string, postData map[string]string) (*PostEntrieResponse, error) {
	response := new(PostEntrieResponse)
	err := api.Client.Post("forms/"+formIdentifier+"/entries", postData, response)
	return response, err
}

// request is internal method to make a request to get fields list
func (api EntriesApi) request(apiUrl string, page int, perPage int) (*EntriesCollection, error) {
	params := make(map[string]string)

	if page > 0 && perPage > 0 {
		params["pageStart"] = strconv.Itoa(page - 1)
		params["pageSize"] = strconv.Itoa(perPage)
	}

	collection := EntriesCollection{make([]map[string]interface{}, 0)}

	err := api.Client.Get(apiUrl, params, &collection)
	if err != nil {
		return nil, err
	}

	return &collection, nil
}
