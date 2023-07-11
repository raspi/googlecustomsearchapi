# googlecustomsearchapi

![GitHub release (latest by date)](https://img.shields.io/github/v/release/raspi/googlecustomsearchapi?style=for-the-badge)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/raspi/googlecustomsearchapi?style=for-the-badge)
[![Go Report Card](https://goreportcard.com/badge/github.com/raspi/googlecustomsearchapi)](https://goreportcard.com/report/github.com/raspi/googlecustomsearchapi)

Google's custom search REST API v1 client for Go

## Example

```go
func main() {
	httpc := http.DefaultClient

	gs := googlecustomsearchapi.New(httpc, "api-key-here", "engine-id-here")

	custom := map[string]string{
		`siteSearchFilter`: `i`,
		`siteSearch`:       `www.imdb.com`,
	}

	res, err := gs.Search(`the room`, 0, custom)
	if err != nil {
		panic(err)
	}

	for _, r := range res.Items {
		fmt.Printf(`%v %v`+"\n", r.Link, r.Title)
	}

}
```

## References

* https://developers.google.com/custom-search/v1/using_rest
* https://developers.google.com/custom-search/v1/overview
* https://console.cloud.google.com/apis/api/customsearch.googleapis.com/metrics
 
