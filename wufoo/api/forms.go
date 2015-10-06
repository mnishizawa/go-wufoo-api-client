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
