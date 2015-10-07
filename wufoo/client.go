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
	Config Config
}

// get function makes a GET request to wufoo API and returns
func (c Client) Get(api string, params map[string]string, response interface{}) (err error) {
	req, err := http.NewRequest("GET", c.PrepareUrl(api, params), nil)
	if err != nil {
		return
	}

	req.SetBasicAuth(c.Config.ApiKey, "footastic")

	resp, err := http.DefaultClient.Do(req)
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
func (c Client) PrepareUrl(api string, params map[string]string) (urlValue string) {
	urlValue = fmt.Sprintf(WUFOO_API_URL_PATTERN, c.Config.Subdomain, api)

	if params != nil && len(params) > 0 {
		getParams := url.Values{}
		for key, value := range params {
			getParams.Add(key, value)
		}

		urlValue += "?" + getParams.Encode()
	}

	return
}
