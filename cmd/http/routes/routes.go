package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/SSH-Management/utils/v2"

	"github.com/SSH-Management/server/cmd/http/middleware"
	"github.com/SSH-Management/server/pkg/container"
)

const (
	RequestIdKey    = "request_id"
	CsrfTokenCookie = "csrf_cookie"
	CsrfTokenKey    = "csrf_token"
)

func Register(c *container.Container, router fiber.Router) {
	router.Use(pprof.New())
	router.Use(recover.New())
	router.Use(middleware.Context)

	router.Use(csrf.New(csrf.Config{
		KeyLookup:  "cookie:csrf_cookie",
		ContextKey: CsrfTokenKey,
		CookieName: CsrfTokenCookie,
		//Storage:        c.GetStorage(0),
		//CookieDomain:   c.Config.Csrf.CookieDomain,
		//CookieSecure:   c.Config.Csrf.Secure,
		CookieHTTPOnly: true,
		KeyGenerator: func() string {
			return utils.RandomString(32)
		},
	}))

	router.Use(requestid.New(requestid.Config{
		Generator: func() string {
			return utils.RandomString(32)
		},
		ContextKey: RequestIdKey,
	}))

	router.Get("/monitor", monitor.New())

	registerClientRoutes(c, router)
}
