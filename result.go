package googlecustomsearchapi

// JSON result from API
// Some fields were removed

type Result struct {
	//URL               URL               `json:"url"`
	Queries           Queries           `json:"queries"`
	SearchInformation SearchInformation `json:"searchInformation"`
	Items             []Item            `json:"items"`
}

type Item struct {
	//Kind             string `json:"kind"`
	Title       string `json:"title"`
	HTMLTitle   string `json:"htmlTitle"`
	Link        string `json:"link"`
	DisplayLink string `json:"displayLink"`
	Snippet     string `json:"snippet"`
	HTMLSnippet string `json:"htmlSnippet"`
	//CacheID          string `json:"cacheId"`
	FormattedURL     string `json:"formattedUrl"`
	HTMLFormattedURL string `json:"htmlFormattedUrl"`
}

type Queries struct {
	Request  []RequestInfo `json:"request"`
	NextPage []RequestInfo `json:"nextPage"`
}

type RequestInfo struct {
	//Title                  string `json:"title"`
	TotalResults string `json:"totalResults"`
	//SearchTerms  string `json:"searchTerms"`
	Count      int64 `json:"count"`
	StartIndex int64 `json:"startIndex"`
	//Language               string `json:"language"`
	Safe                   string `json:"safe"`
	DisableCNTwTranslation string `json:"disableCnTwTranslation"`
	Hl                     string `json:"hl"`
}

type SearchInformation struct {
	SearchTime float64 `json:"searchTime"`
	//FormattedSearchTime   string  `json:"formattedSearchTime"`
	TotalResults string `json:"totalResults"`
	//FormattedTotalResults string  `json:"formattedTotalResults"`
}

type URL struct {
	Type     string `json:"type"`
	Template string `json:"template"`
}
