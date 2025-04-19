package keycloak

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

func TestKeycloakSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Keycloak Test Suite")
}

var _ = BeforeSuite(func() {
	var err error
	bm, err = utils.NewBrowserManager()
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	Expect(bm.Close()).To(Succeed())
})

var _ = Describe("Login Flow", func() {
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

	It("should authenticate", func(ctx SpecContext) {
		container, err := utils.CreateContainer(ctx, "configs/basic.cfg", []string{"02_keycloak_keycloak", "02_keycloak_upstream"})
		Expect(err).NotTo(HaveOccurred(), "Couldn't create container")
		defer container.Terminate(ctx)

		baseUrl := "http://oauth2-proxy.localtest.me:4180"

		pages.NewProviderButtonPage(page, baseUrl, "Keycloak").SignIn()
		pages.KeycloakLogin(page, "admin@example.com", "password")

		httpbin := pages.NewHttpbinPage(page, baseUrl)
		headers := httpbin.GetHeaders()

		Expect(headers["X-Forwarded-Email"]).To(Equal("admin@example.com"))
	})
})
