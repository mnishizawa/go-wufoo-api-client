package wufoo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"encoding/json"
)

// Client represents API client for Wufoo
type Client struct {
	config Config
}

// Creates a new Client object with the provided configuration
func New(config Config) Client {
	return &Client{config: config}
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
func (c Client) Forms(includeTodayCount bool) {
	params := make(map[string]string)
	params["IncludeTodayCount"] = includeTodayCount

	resp, err := http.Get(c.prepareUrl("forms", params))
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
func (c Client) FormsDetails(formIdentifier string, includeTodayCount bool) (err error) {
	params := make(map[string]string)
	params["IncludeTodayCount"] = includeTodayCount

	err = c.get("forms/"+formIdentifier, params, &response)
	if err != nil {
		return
	}
}

// get function makes a GET request to wufoo API and returns
func (c Client) get(api string, params map[string]string, response interface{}) (err error) {
	resp, err := http.Get(c.prepareUrl(api, params))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}

	return nil
}

// prepareUrl function returns URL for specific API with GET parameters if specified
func (c Client) prepareUrl(api string, params map[string]string) (url string) {
	url = fmt.Sprintf(WUFOO_API_URL_PATTERN, c.config.Subdomain, api)

	if params != nil && len(params) > 0 {
		getParams := url.Values{}
		for key, value := range params {
			getParams.Add(key, value)
		}

		url += "?" + getParams.Encode()
	}

	return
}
