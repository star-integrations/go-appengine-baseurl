package baseurl

import (
	"context"
	"fmt"
	"net/url"

	"google.golang.org/appengine"
)

// NewBaseURL 接続先のAppEngineBaseURLを作成する
func NewBaseURL(c context.Context, version, service, domain string) (*url.URL, error) {
	hostName := ""
	if domain == "" {
		hostName = appengine.DefaultVersionHostname(c)
	} else {
		hostName = domain
	}
	if service != "" {
		hostName = fmt.Sprintf("%s-dot-%s", service, hostName)
	}
	if version != "" {
		hostName = fmt.Sprintf("%s-dot-%s", version, hostName)
	}
	hostName = fmt.Sprintf("https://%s", hostName)

	return url.Parse(hostName)
}
