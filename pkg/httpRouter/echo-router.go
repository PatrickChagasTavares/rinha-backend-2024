package httpRouter

import (
	"context"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/patrickchagastavares/rinha-backend-2024/pkg/validator"
)

type echoRouter struct {
	router *echo.Echo
}

func NewEchoRouter() Router {

	router := echo.New()
	if os.Getenv("env") != "local" {
		router.HideBanner = true
		router.HidePort = true
	}
	router.Use(
		middleware.Recover(),
	)

	return &echoRouter{
		router: router,
	}
}

func (r *echoRouter) Get(path string, f HandlerFunc) {
	r.router.GET(path, func(ctx echo.Context) error {
		f(newEchoContext(ctx))
		return nil
	})
}

func (r *echoRouter) Post(path string, f HandlerFunc) {
	r.router.POST(path, func(ctx echo.Context) error {
		f(newEchoContext(ctx))
		return nil
	})
}

func (r *echoRouter) Server(port string) error {
	return http.ListenAndServe(port, r.router)
}

func (m *echoRouter) ParseHandler(h http.HandlerFunc) HandlerFunc {
	return func(c Context) {
		h(c.GetResponseWriter(), c.GetRequestReader())
	}
}

type echoContext struct {
	r echo.Context
	v validator.Validator
}

func newEchoContext(ctx echo.Context) Context {
	return &echoContext{
		r: ctx,
		v: validator.New(),
	}
}

func (c *echoContext) Context() context.Context {
	return c.r.Request().Context()
}

func (c *echoContext) String(statusCode int, value string) {
	c.r.String(statusCode, value)
}

func (c *echoContext) JSON(statusCode int, data any) {
	c.r.JSON(statusCode, data)
}

func (c *echoContext) DecodeJSON(data any) error {
	return c.r.Bind(data)
}

func (c *echoContext) GetQuery(query string) string {
	return c.r.QueryParam(query)
}

func (c *echoContext) GetParam(param string) string {
	return c.r.Param(param)
}

func (c *echoContext) GetResponseWriter() http.ResponseWriter {
	return c.r.Response().Writer
}

func (c *echoContext) GetRequestReader() *http.Request {
	return c.r.Request().Response.Request
}

func (c *echoContext) Validate(input any) error {
	return c.v.Validate(input)
}
