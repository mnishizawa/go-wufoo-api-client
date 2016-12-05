package api

import (
	"net/url"
	"strconv"
)

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
		ID        string
		ErrorText string
	}
}

type FilterGroup struct {
	Filters  []Filter
	Grouping string
}

type Filter struct {
	FieldId    string
	Operator   string
	MatchValue string
}

type Sort struct {
	FieldId   string
	Direction string
}

type Page struct {
	Offset int
	Size   int
}

func (grp FilterGroup) Size() int {
	return len(grp.Filters)
}

func (grp FilterGroup) QueryString() string {
	//append filters unencoded to the url
	var url string
	numFilters := len(grp.Filters)

	if grp.Filters != nil && grp.Size() > 0 {
		for idx, filter := range grp.Filters {
			filterSpec := "Filter" + strconv.Itoa(idx+1) + "=" + filter.FieldId + "+" + filter.Operator + "+" + filter.MatchValue
			url += filterSpec
			if idx < numFilters {
				url += "&"
			}
		}
		url += "match=" + grp.Grouping
	}

	return url
}

// Fields method returns form fields details.
//
// For more details please visit: http://help.wufoo.com/articles/en_US/SurveyMonkeyArticleType/The-Entries-POST-API
//
func (api EntriesApi) Entries(formIdentifier string, filters *FilterGroup, sort *Sort, page *Page) (collection *EntriesCollection, err error) {
	collection, err = api.request("forms/"+formIdentifier+"/entries", filters, sort, page)
	return
}

// FieldsReport method returns form fields details.
//
// For more details please visit: http://help.wufoo.com/articles/en_US/SurveyMonkeyArticleType/The-Entries-POST-API
//
func (api EntriesApi) EntriesReport(formIdentifier string, filters *FilterGroup, sort *Sort, page *Page) (collection *EntriesCollection, err error) {
	collection, err = api.request("reports/"+formIdentifier+"/entries", filters, sort, page)
	return
}

// PostEntries method returns form fields details.
//
// For more details please visit: http://help.wufoo.com/articles/en_US/SurveyMonkeyArticleType/The-Entries-POST-API
//
func (api EntriesApi) PostEntries(formIdentifier string, postData url.Values) (*PostEntrieResponse, error) {
	response := new(PostEntrieResponse)
	err := api.Client.Post("forms/"+formIdentifier+"/entries", postData, response)
	return response, err
}

// request is internal method to make a request to get fields list
func (api EntriesApi) request(apiUrl string, filters *FilterGroup, sort *Sort, page *Page) (*EntriesCollection, error) {
	params := make(map[string]string)

	if page != nil {
		params["pageStart"] = strconv.Itoa(page.Offset - 1)
		params["pageSize"] = strconv.Itoa(page.Size)
	}

	if sort != nil {
		params["sort"] = sort.FieldId
		params["sortDirection"] = sort.Direction
	}

	//if filters != nil {
	//	for idx, filter := range filters.Filters {
	//		params["Filter"+strconv.Itoa(idx)] = filter.FieldId+"+"+filter.Operator+"+"+filter.MatchValue
	//	}
	//
	//	params["match"] = filters.Grouping
	//}

	collection := EntriesCollection{make([]map[string]interface{}, 0)}

	err := api.Client.Get(apiUrl, params, filters, &collection)
	if err != nil {
		return nil, err
	}

	return &collection, nil
}
