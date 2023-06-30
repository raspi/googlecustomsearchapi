package googlecustomsearchapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const API_URL = "https://customsearch.googleapis.com/customsearch/v1"

// Client uses Google's custom search REST API
// See https://developers.google.com/custom-search/v1/reference/rest/v1/cse/list
type Client struct {
	apiKey   string
	engineId string
	cl       *http.Client
}

func New(httpClient *http.Client, apikey string, engineid string) *Client {
	return &Client{
		apiKey:   apikey,
		engineId: engineid,
		cl:       httpClient,
	}
}

// Search searches Google's custom search REST API with given query
func (s Client) Search(query string, start uint, custom map[string]string) (Result, error) {
	if query == `` {
		return Result{}, fmt.Errorf(`empty query`)
	}

	q := url.Values{}
	q.Set(`c2coff`, `1`)
	q.Set(`hl`, `en`)
	q.Set(`lr`, `lang_en`)
	q.Set(`num`, `10`)
	q.Set(`start`, fmt.Sprintf(`%d`, start))
	q.Set(`q`, query)
	q.Set(`cx`, s.engineId)
	q.Set(`key`, s.apiKey)

	if custom != nil {
		for k, v := range custom {
			switch k {
			case `q`, `key`, `cx`:
				// Do not allow to change query or API keys
				continue
			}

			q.Set(k, v)
		}
	}

	resp, err := s.cl.Get(API_URL + `?` + q.Encode())
	if err != nil {
		return Result{}, err
	}

	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{}, err
	}

	if resp.StatusCode == http.StatusOK {
		var r Result
		err = json.Unmarshal(content, &r)
		if err != nil {
			return Result{}, err
		}

		return r, nil
	}

	var tmperr Error
	err = json.Unmarshal(content, &tmperr)
	if err != nil {
		return Result{}, err
	}

	return Result{}, tmperr.Error
}
