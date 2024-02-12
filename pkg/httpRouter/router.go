package httpRouter

import (
	"context"
	"net/http"
)

type (
	Router interface {
		Server(port string) error
		Get(path string, f HandlerFunc)
		Post(path string, f HandlerFunc)
		ParseHandler(h http.HandlerFunc) HandlerFunc
	}

	HandlerFunc func(ctx Context)

	Context interface {
		// Context returns the request's context. To change the context, use
		// Clone or WithContext.
		//
		// The returned context is always non-nil; it defaults to the
		// background context.
		//
		// For outgoing client requests, the context controls cancellation.
		//
		// For incoming server requests, the context is canceled when the
		// client's connection closes, the request is canceled (with HTTP/2),
		// or when the ServeHTTP method returns.
		Context() context.Context
		// JSON serializes the given struct as JSON into the response body.
		// It also sets the Content-Type as "application/json".
		JSON(statusCode int, data any)
		// String writes the given string into the response body.
		String(statusCode int, value string)
		DecodeJSON(data any) error
		GetResponseWriter() http.ResponseWriter
		GetRequestReader() *http.Request
		// Param returns the value of the URL param.
		// It is a shortcut for c.Params.ByName(key)
		//
		//	router.GET("/user/:id", func(c *gin.Context) {
		//	    // a GET request to /user/john
		//	    id := c.Param("id") // id == "/john"
		//	    // a GET request to /user/john/
		//	    id := c.Param("id") // id == "/john/"
		//	})
		GetParam(param string) string
		Validate(input any) error
	}
)
