package utils

import (
	"github.com/playwright-community/playwright-go"
)

type BrowserManager struct {
	pw      *playwright.Playwright
	browser playwright.Browser
}

func NewBrowserManager() (*BrowserManager, error) {
	pw, err := playwright.Run()
	if err != nil {
		return nil, err
	}

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	if err != nil {
		return nil, err
	}

	return &BrowserManager{
		pw:      pw,
		browser: browser,
	}, nil
}

func (bm *BrowserManager) NewTestContext() (playwright.BrowserContext, playwright.Page, error) {
	context, err := bm.browser.NewContext()
	if err != nil {
		return nil, nil, err
	}

	page, err := context.NewPage()
	return context, page, err
}

func (bm *BrowserManager) Close() error {
	if err := bm.browser.Close(); err != nil {
		return err
	}
	return bm.pw.Stop()
}
