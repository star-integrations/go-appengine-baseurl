package baseurl

import (
	"fmt"
	"testing"

	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
)

func TestNewBaseURL(t *testing.T) {
	c, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	// ドメインが空でデフォルトを設定
	if url, err := NewBaseURL(c, "", "", ""); err != nil || url.String() != "https:" {
		t.Errorf("unmatch DefaultVersionHostname")
	}
	// デフォルトのサービスを設定
	if url, err := NewBaseURL(c, "", "service", ""); err != nil || url.String() != fmt.Sprintf("https://service-dot-%s", appengine.DefaultVersionHostname(c)) {
		t.Errorf("unmatch service-DefaultVersionHostname")
	}
	// デフォルトのバージョンを設定
	if url, err := NewBaseURL(c, "version", "service", ""); err != nil || url.String() != fmt.Sprintf("https://version-dot-service-dot-%s", appengine.DefaultVersionHostname(c)) {
		t.Errorf("unmatch service-DefaultVersionHostname")
	}
	// ドメインを設定
	if url, err := NewBaseURL(c, "", "", "mydomain.com"); err != nil || url.String() != "https://mydomain.com" {
		t.Errorf("unmatch domain")
	}
	// サービスを設定
	if url, err := NewBaseURL(c, "", "service", "mydomain.com"); err != nil || url.String() != "https://service-dot-mydomain.com" {
		t.Errorf("unmatch service-domain")
	}
	// バージョンを設定
	if url, err := NewBaseURL(c, "version", "service", "mydomain.com"); err != nil || url.String() != "https://version-dot-service-dot-mydomain.com" {
		t.Errorf("unmatch version-service-domain")
	}
}
