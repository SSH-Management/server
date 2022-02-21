package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/SSH-Management/server/pkg/constants"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/valyala/fasthttp"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/memory"

	"github.com/SSH-Management/utils/v2"

	"github.com/SSH-Management/server/pkg/config"
	sshmanagementapp "github.com/SSH-Management/server/pkg/http"
	"github.com/SSH-Management/server/pkg/http/middleware"
)

const (
	SessionCookieName  = "testing_ssh_management_cookie_name"
	SessionCookieValue = "test_ssh_management_value"
)

func GetValidator() (*validator.Validate, ut.Translator) {
	v := validator.New()
	english := en.New()
	uni := ut.New(english, english)
	englishTranslations, _ := uni.GetTranslator("en")

	return v, englishTranslations
}

func CreateSessionCookie() *http.Cookie {
	return &http.Cookie{
		Name:     SessionCookieName,
		Value:    SessionCookieValue,
		Path:     "/",
		Domain:   "localhost",
		Expires:  time.Now().Add(time.Hour),
		MaxAge:   0,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
}

func GetSessionWithNameAndValue(app *fiber.App, store *session.Store, name, value string) (*session.Session, func()) {
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	ctx.Request().Header.SetCookie(name, value)

	sess, err := store.Get(ctx)
	if err != nil {
		panic(err)
	}

	return sess, func() {
		app.ReleaseCtx(ctx)
	}
}

func GetSession(app *fiber.App, store *session.Store) (*session.Session, func()) {
	return GetSessionWithNameAndValue(app, store, SessionCookieName, SessionCookieValue)
}

func GetSessionWithValue(app *fiber.App, store *session.Store, value string) (*session.Session, func()) {
	return GetSessionWithNameAndValue(app, store, SessionCookieName, value)
}

func CreateApplication() (*fiber.App, *session.Store) {
	app := sshmanagementapp.CreateApplication("/static", "", false, config.Testing, nil)

	store := session.New(session.Config{
		Storage:        memory.New(),
		CookieHTTPOnly: true,
		Expiration:     1 * time.Hour,
		KeyLookup:      fmt.Sprintf("cookie:%s", SessionCookieName),
		CookieDomain:   "localhost",
		CookiePath:     "/",
		CookieSecure:   false,
		CookieSameSite: "strict",
		KeyGenerator: func() string {
			return utils.RandomString(32)
		},
	})

	return app, store
}

// TODO: CreateApplicationWithUser (*fiber.App, *session.Store, models.User)

func CreateApplicationWithSession() (*fiber.App, *session.Store) {
	app, store := CreateApplication()

	app.Use(middleware.Auth(store))

	return app, store
}

func getBody(headers http.Header, body interface{}) io.Reader {
	switch headers.Get("Content-Type") {
	case "application/json":
		jsonStr, err := json.Marshal(body)
		if err != nil {
			panic(err)
		}

		return bytes.NewReader(jsonStr)
	default:
		return nil
	}
}

type RequestModifier func(*http.Request) *http.Request

func WithHeaders(headers http.Header) RequestModifier {
	return func(req *http.Request) *http.Request {
		if headers.Get("Content-Type") == "" {
			headers.Set("Content-Type", "application/json")
		}

		if headers.Get("Accept") == "" {
			headers.Set("Accept", "application/json")
		}

		if headers.Get("User-Agent") == "" {
			headers.Set("User-Agent", constants.TestUserAgent)
		}

		req.Header = headers

		return req
	}
}

func WithBody(body interface{}) RequestModifier {
	return func(req *http.Request) *http.Request {
		newReq := httptest.NewRequest(req.Method, req.URL.Path, getBody(req.Header, body))
		newReq.Header = req.Header
		newReq.URL.RawQuery = req.URL.Query().Encode()

		for _, cookie := range req.Cookies() {
			newReq.AddCookie(cookie)
		}

		return newReq
	}
}

func WithCookies(cookies []*http.Cookie) RequestModifier {
	return func(req *http.Request) *http.Request {
		for _, cookie := range cookies {
			req.AddCookie(cookie)
		}

		return req
	}
}

func WithSessionCookie() RequestModifier {
	return WithCookies([]*http.Cookie{CreateSessionCookie()})
}

func MakeRequest(method, uri string, modifiers ...RequestModifier) *http.Request {
	var defaults []func(*http.Request) *http.Request

	switch method {
	case http.MethodPost, http.MethodPut, http.MethodPatch:
		defaults = []func(*http.Request) *http.Request{
			WithHeaders(http.Header{}),
			WithBody(nil),
		}
	default:
		defaults = []func(*http.Request) *http.Request{
			WithHeaders(http.Header{}),
		}
	}

	req := httptest.NewRequest(method, uri, nil)

	for _, modifier := range defaults {
		req = modifier(req)
	}

	for _, modifier := range modifiers {
		req = modifier(req)
	}

	return req
}

func Get(app *fiber.App, uri string, modifiers ...RequestModifier) *http.Response {
	req := MakeRequest(http.MethodGet, uri, modifiers...)

	res, err := app.Test(req, -1)
	if err != nil {
		panic("Cannot get response")
	}

	return res
}

func Post(app *fiber.App, uri string, modifiers ...RequestModifier) *http.Response {
	req := MakeRequest(http.MethodPost, uri, modifiers...)

	res, err := app.Test(req, -1)
	if err != nil {
		panic("Cannot get response")
	}

	return res
}

func Put(app *fiber.App, uri string, modifiers ...RequestModifier) *http.Response {
	req := MakeRequest(http.MethodPut, uri, modifiers...)

	res, err := app.Test(req, -1)
	if err != nil {
		panic("Cannot get response")
	}

	return res
}

func Patch(app *fiber.App, uri string, modifiers ...RequestModifier) *http.Response {
	req := MakeRequest(http.MethodPatch, uri, modifiers...)

	res, err := app.Test(req, -1)
	if err != nil {
		panic("Cannot get response")
	}

	return res
}

func Delete(app *fiber.App, uri string, modifiers ...RequestModifier) *http.Response {
	req := MakeRequest(http.MethodDelete, uri, modifiers...)

	res, err := app.Test(req, -1)
	if err != nil {
		panic("Cannot get response")
	}

	return res
}
