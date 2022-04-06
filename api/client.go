package api

import (
	"context"
	"net/url"
)

type Client interface {
	Log(string, string, ...interface{})
	Do(context.Context, string, []string, url.Values, interface{}, interface{}, ...error) ([]byte, error)
}
