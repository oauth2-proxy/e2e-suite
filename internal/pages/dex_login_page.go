package pages

import (
	. "github.com/onsi/gomega"
	pw "github.com/playwright-community/playwright-go"
)

func DexLogin(page pw.Page, username, password string) {
	usernameInput := page.Locator("input#login")
	passwordInput := page.Locator("input#password")

	err := usernameInput.Fill(username)
	Expect(err).NotTo(HaveOccurred(), "Couldn't enter username")

	err = passwordInput.Fill(password)
	Expect(err).NotTo(HaveOccurred(), "Couldn't enter password")

	btn := page.Locator("button", pw.PageLocatorOptions{HasText: "Login"})
	err = btn.Click()
	Expect(err).NotTo(HaveOccurred(), "Couldn't login")

	btn = page.Locator("button", pw.PageLocatorOptions{HasText: "Grant Access"})
	err = btn.Click()
	Expect(err).NotTo(HaveOccurred(), "Couldn't grant access for scopes")
}
