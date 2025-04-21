package nginx

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

func TestNginxIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Nginx Integration Suite")
}

var _ = BeforeSuite(func() {
	var err error
	bm, err = utils.NewBrowserManager()
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	Expect(bm.Close()).To(Succeed())
})

var _ = Describe("Middleware Subrequest Flow", func() {
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

	It("should authenticate via Dex", func(ctx SpecContext) {
		container, err := utils.CreateContainer(ctx, "configs/basic.cfg", []string{"03_nginx_app"})
		Expect(err).NotTo(HaveOccurred(), "Couldn't create container")
		defer container.Terminate(ctx)

		baseUrl := "http://app.localtest.me:8080"

		pages.NewProviderButtonPage(page, baseUrl, "Dex").SignIn()
		pages.DexLogin(page, "admin@example.com", "password")

		httpbin := pages.NewHttpbinPage(page, baseUrl)
		headers := httpbin.GetHeaders()

		Expect(headers["Cookie"]).To(HavePrefix("_oauth2_proxy"))
		Expect(headers).To(Not(HaveKey("X-Access-Token")))
		Expect(headers).To(Not(HaveKey("X-Email")))
	})

	It("should forward X-* headers and tokens to upstream", func(ctx SpecContext) {
		container, err := utils.CreateContainer(ctx, "configs/pass-access-token-and-xauthrequest.cfg", []string{"03_nginx_app"})
		Expect(err).NotTo(HaveOccurred(), "Couldn't create container")
		defer container.Terminate(ctx)

		baseUrl := "http://app.localtest.me:8080"

		pages.NewProviderButtonPage(page, baseUrl, "Dex").SignIn()
		pages.DexLogin(page, "admin@example.com", "password")

		httpbin := pages.NewHttpbinPage(page, baseUrl)
		headers := httpbin.GetHeaders()

		Expect(headers["Cookie"]).To(HavePrefix("_oauth2_proxy"))
		Expect(headers).To(HaveKey("X-Access-Token"))
		Expect(headers).To(HaveKey("X-Email"))
		Expect(headers["X-Email"]).To(Equal("admin@example.com"))
	})
})
