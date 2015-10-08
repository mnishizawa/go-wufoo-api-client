package api

// Form structure represent wufoo form details payload
// For more details please check: http://help.wufoo.com/articles/en_US/SurveyMonkeyArticleType/The-Forms-API
type Form struct {
	Name             string
	Description      string
	RedirectMessage  string
	Url              string
	Email            string
	IsPublic         string
	Language         string
	StartDate        string
	EndDate          string
	EntryLimit       string
	DateCreated      string
	DateUpdated      string
	Hash             string
	LinkFields       string
	LinkEntries      string
	LinkEntriesCount string
}

// FormCollection structure represents Wufoo Form API response
type FormCollection struct {
	Forms []Form
}

// FormsApi struct represents access to Wufoo Form API
type FormsApi struct {
	Client *Client
}

// Forms method returns details about all available forms for the API key which was set up in config. It can be used
// to create a list of all forms and dynamically generate a form embed snippet to use in your application.
//
// Method receive the following parameters:
//
//     includeTodayCount - will give you today’s entry count for the form.
//
// For more details please visit: http://help.wufoo.com/articles/en_US/SurveyMonkeyArticleType/The-Forms-API
//
func (api FormsApi) Forms(includeTodayCount bool) (*FormCollection, error) {
	params := make(map[string]string)

	if includeTodayCount {
		params["IncludeTodayCount"] = "true"
	} else {
		params["IncludeTodayCount"] = "false"
	}

	collection := FormCollection{make([]Form, 0)}

	err := api.Client.Get("forms", params, &collection)
	if err != nil {
		return nil, err
	}

	return &collection, nil
}

// FormsDetails method returns details about one form.
//
// Method receive the following parameters:
//     formIdentifier - if not nil Forms will give you information about just one form.
//                      The call with formIdentifier = nil will return all forms
//
//     includeTodayCount - will give you today’s entry count for the form.
//
// For more details please visit: http://help.wufoo.com/articles/en_US/SurveyMonkeyArticleType/The-Forms-API
//
func (api FormsApi) FormsDetails(formIdentifier string, includeTodayCount bool) (*Form, error) {
	params := make(map[string]string)

	if includeTodayCount {
		params["IncludeTodayCount"] = "true"
	} else {
		params["IncludeTodayCount"] = "false"
	}

	collection := FormCollection{make([]Form, 0)}

	err := api.Client.Get("forms/"+formIdentifier, params, &collection)
	if err != nil {
		return nil, err
	}

	var form Form

	if len(collection.Forms) > 0 {
		form = collection.Forms[0]
	}

	return &form, nil
}
