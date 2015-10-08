package wufoo

const WUFOO_API_URL_PATTERN = "https://%s.wufoo.com/api/v3/%s.json"

// Config provides structure to configure wufoo client
type Config struct {
	ApiKey    string
	Subdomain string
}
