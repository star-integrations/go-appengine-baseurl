package baseurl

import (
	"fmt"
	"testing"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
)

func TestNew(t *testing.T) {
	inst, err := aetest.NewInstance(&aetest.Options{StartupTimeout: 120 * time.Second})
	if err != nil {
		t.Fatal(err)
	}
	defer inst.Close()
	c, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	// Default
	if url, err := New(c); err != nil || url.String() != "https://" {
		t.Errorf("unmatch DefaultVersionHostname")
	}

	// with Project
	if url, err := New(c, WithProject("project"), WithService("service")); err != nil || url.String() != "https://service-dot-project.appspot.com" {
		t.Errorf("unmatch service-project")
	}

	// with Service
	if url, err := New(c, WithService("service")); err != nil || url.String() != fmt.Sprintf("https://service-dot-%s", appengine.DefaultVersionHostname(c)) {
		t.Errorf("unmatch service-DefaultVersionHostname")
	}

	// with Version
	if url, err := New(c, WithService("service"), WithVersion("version")); err != nil || url.String() != fmt.Sprintf("https://version-dot-service-dot-%s", appengine.DefaultVersionHostname(c)) {
		t.Errorf("unmatch version-service-DefaultVersionHostname")
	}

	// with Domain
	if url, err := New(c, WithDomain("mydomain.com")); err != nil || url.String() != "https://mydomain.com" {
		t.Errorf("unmatch domain")
	}

	// with Domain and Service
	if url, err := New(c, WithService("service"), WithDomain("mydomain.com")); err != nil || url.String() != "https://service-dot-mydomain.com" {
		t.Errorf("unmatch service-domain")
	}

	// with Domain and Service and Version
	if url, err := New(c, WithVersion("version"), WithService("service"), WithDomain("mydomain.com")); err != nil || url.String() != "https://version-dot-service-dot-mydomain.com" {
		t.Errorf("unmatch version-service-domain")
	}
}
