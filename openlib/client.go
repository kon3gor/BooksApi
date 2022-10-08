package openlib

import (
	"fmt"
	"net/http"
	"log"
	"strings"
	"net/url"
)

const OpenLibUrl = "openlibrary.org"

func FetchSmthFromOpenLib(
	method string, 
	path string, 
	params map[string]string,
) *http.Response {
	client := &http.Client{}
	realPath := fmt.Sprintf("%s.json", path)
	url := &url.URL{
		Scheme: "https",
		Host: OpenLibUrl,
		Path: realPath,
		RawQuery: buildQueryString(params),
	}

	req, err  := http.NewRequest(method, url.String(), nil)
	if (err != nil) {
		log.Fatalln(err)
		return nil
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return res
}

func buildQueryString(params map[string]string) string {
	var sb strings.Builder
	for k, v := range params {
		fmt.Fprintf(&sb, "%s=%s&", k, v)
	}
	return strings.TrimSuffix(sb.String(), "&")
}
