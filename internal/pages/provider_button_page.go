package pages

import (
	. "github.com/onsi/gomega"
	pw "github.com/playwright-community/playwright-go"
)

type ProviderButtonPage struct {
	page         pw.Page
	baseUrl      string
	providerName string
}

func NewProviderButtonPage(page pw.Page, baseUrl string, providerName string) *ProviderButtonPage {
	return &ProviderButtonPage{page: page, baseUrl: baseUrl, providerName: providerName}
}

func (p *ProviderButtonPage) SignIn() {
	_, err := p.page.Goto(p.baseUrl)
	Expect(err).ToNot(HaveOccurred(), "Provider button page not loading")

	btn := p.page.Locator("button", pw.PageLocatorOptions{HasText: "Sign in with Dex"})

	err = btn.Click()
	Expect(err).NotTo(HaveOccurred(), "Provider button wasn't found")
}
