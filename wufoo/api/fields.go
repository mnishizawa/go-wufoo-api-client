package api

// Field structure represent wufoo field details payload
// For more details please check: http://help.wufoo.com/articles/en_US/SurveyMonkeyArticleType/The-Fields-API
type Field struct {
	Title         string
	Type          string
	ID            string
	IsRequired    string `json:",omitempty"`
	HasOtherField bool   `json:",omitempty"`

	SubFields []struct {
		ID          string
		Label       string
		DefaultVal  string
		ChoicesText string
		ColumnId    string
	} `json:",omitempty"`

	Choices []struct {
		Label string
		Score string
	} `json:",omitempty"`

	Page          string `json:",omitempty"`
	IsSystem      bool   `json:",omitempty"`
	ClassNames    string `json:",omitempty"`
	Instructions  string `json:",omitempty"`
	DefaultVal    string `json:",omitempty"`
	PurchaseTotal string `json:",omitempty"`
	Currency      string `json:",omitempty"`
	Status        string `json:",omitempty"`
}

// FieldCollection structure represents Wufoo Field API response
type FieldCollection struct {
	Fields []Field
}

// FieldsApi struct represents access to Wufoo Field API
type FieldsApi struct {
	Client *Client
}

// Fields method returns form fields details.
//
// Method receive the following parameters:
//     formIdentifier - form hash
//
//     system - includes system fields in response payload
//
// For more details please visit: http://help.wufoo.com/articles/en_US/SurveyMonkeyArticleType/The-Fields-API
//
func (api FieldsApi) Fields(formIdentifier string, system bool) (collection *FieldCollection, err error) {
	collection, err = api.request("forms/"+formIdentifier+"/fields", system)
	return
}

// FieldsReport method returns form fields details.
//
// Method receive the following parameters:
//     formIdentifier - form hash
//
//     system - includes system fields in response payload
//
// For more details please visit: http://help.wufoo.com/articles/en_US/SurveyMonkeyArticleType/The-Forms-API
//
func (api FieldsApi) FieldsReport(formIdentifier string, system bool) (collection *FieldCollection, err error) {
	collection, err = api.request("reports/"+formIdentifier+"/fields", system)
	return
}

// request is internal method to make a request to get fields list
func (api FieldsApi) request(apiUrl string, system bool) (*FieldCollection, error) {
	params := make(map[string]string)

	if system {
		params["system"] = "true"
	}

	collection := FieldCollection{make([]Field, 0)}

	err := api.Client.Get(apiUrl, params, &collection)
	if err != nil {
		return nil, err
	}

	return &collection, nil
}
