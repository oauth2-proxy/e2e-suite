package utils

import (
	"os"
	"strconv"

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
		Headless: playwright.Bool(getHeadless()),
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
	page.SetDefaultTimeout(10000)
	return context, page, err
}

func (bm *BrowserManager) Close() error {
	if err := bm.browser.Close(); err != nil {
		return err
	}
	return bm.pw.Stop()
}

func getHeadless() bool {
	val := os.Getenv("BROWSER_HEADLESS")
	if val == "" {
		return true
	}

	// Parse as boolean (accepts "false", "true", "0", "1", "t", "yes", "y", etc.)
	b, err := strconv.ParseBool(val)
	if err != nil {
		return true
	}
	return b
}
