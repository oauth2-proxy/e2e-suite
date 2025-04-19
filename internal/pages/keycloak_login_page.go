package pages

import (
	. "github.com/onsi/gomega"
	pw "github.com/playwright-community/playwright-go"
)

func KeycloakLogin(page pw.Page, username, password string) {
	usernameInput := page.Locator("input#username")
	passwordInput := page.Locator("input#password")

	err := usernameInput.Fill(username)
	Expect(err).NotTo(HaveOccurred(), "Couldn't enter username")

	err = passwordInput.Fill(password)
	Expect(err).NotTo(HaveOccurred(), "Couldn't enter password")

	btn := page.Locator("button#kc-login", pw.PageLocatorOptions{HasText: "Sign In"})
	err = btn.Click()
	Expect(err).NotTo(HaveOccurred(), "Couldn't login")
}
