package baseurl

import (
	"context"
	"fmt"
	"net/url"

	"google.golang.org/appengine"
)

// BaseURL type
type BaseURL struct {
	version string
	service string
	domain  string
}

// Option type
type Option func(*BaseURL) error

// New Create AppEngine URL
func New(c context.Context, options ...Option) (*url.URL, error) {
	b := &BaseURL{}
	for _, option := range options {
		err := option(b)
		if err != nil {
			return nil, err
		}
	}

	hostName := ""
	if b.domain == "" {
		hostName = appengine.DefaultVersionHostname(c)
	} else {
		hostName = b.domain
	}

	if b.service != "" {
		hostName = fmt.Sprintf("%s-dot-%s", b.service, hostName)
	}

	if b.version != "" {
		hostName = fmt.Sprintf("%s-dot-%s", b.version, hostName)
	}

	baseURL := fmt.Sprintf("https://%s", hostName)

	return url.Parse(baseURL)
}

// WithService function
// See https://cloud.google.com/appengine/docs/standard/go/microservices-on-app-engine#service_isolation
func WithService(s string) Option {
	return func(b *BaseURL) error {
		b.service = s
		return nil
	}
}

// WithVersion function
// See https://cloud.google.com/appengine/docs/standard/go/microservices-on-app-engine#versions_within_services
func WithVersion(v string) Option {
	return func(b *BaseURL) error {
		b.version = v
		return nil
	}
}

// WithDomain function
func WithDomain(d string) Option {
	return func(b *BaseURL) error {
		b.domain = d
		return nil
	}
}
