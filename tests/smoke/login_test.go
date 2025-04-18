package smoke

import (
	"testing"

	"github.com/oauth2-proxy/e2e/internal/pages"
	"github.com/oauth2-proxy/e2e/internal/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	pw "github.com/playwright-community/playwright-go"
)

var (
	bm *utils.BrowserManager
)

func TestOIDCLogin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OIDC Login Smoke Tests")
}

var _ = BeforeSuite(func() {
	var err error
	bm, err = utils.NewBrowserManager()
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	Expect(bm.Close()).To(Succeed())
})

var _ = Describe("OIDC Login Flow", func() {
	var (
		context pw.BrowserContext
		page    pw.Page
	)

	BeforeEach(func() {
		var err error
		context, page, err = bm.NewTestContext()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(context.Close()).To(Succeed())
	})

	It("should authenticate via Dex", func() {
		baseUrl := "http://oauth2-proxy.localtest.me:4180"

		pages.NewProviderButtonPage(page, baseUrl, "Dex").SignIn()
		pages.DexLogin(page, "admin@example.com", "password")

		httpbin := pages.NewHttpbinPage(page, baseUrl)
		headers := httpbin.GetHeaders()

		Expect(headers["X-Forwarded-Email"]).To(Equal("admin@example.com"))
	})
})
