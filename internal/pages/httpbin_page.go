package pages

import (
	"encoding/json"
	"net/url"

	. "github.com/onsi/gomega"
	pw "github.com/playwright-community/playwright-go"
)

type HttpbinPage struct {
	page    pw.Page
	baseUrl string
}

func NewHttpbinPage(page pw.Page, baseUrl string) *HttpbinPage {
	return &HttpbinPage{page: page, baseUrl: baseUrl}
}

func (p *HttpbinPage) GetHeaders() map[string]string {
	uri, err := url.JoinPath(p.baseUrl, "/headers")
	Expect(err).NotTo(HaveOccurred(), "URL isn't valid")

	resp, err := p.page.Goto(uri)
	Expect(err).NotTo(HaveOccurred(), "Couldn't open Httpbin headers endpoint")

	body, err := resp.Body()
	Expect(err).NotTo(HaveOccurred(), "Couldn't load Httpbin headers response")

	type HeadersResp struct {
		Headers map[string]string `json:"headers"`
	}

	var result HeadersResp
	err = json.Unmarshal(body, &result)
	Expect(err).NotTo(HaveOccurred(), "Couldn't parse Httpbin headers response")

	return result.Headers
}
