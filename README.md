# go-appengine-baseurl

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/star-integrations/go-appengine-baseurl/baseurl)
[![Build Status](https://travis-ci.org/star-integrations/go-appengine-baseurl.svg?branch=master)](https://travis-ci.org/star-integrations/go-appengine-baseurl)
[![codecov](https://codecov.io/gh/star-integrations/go-appengine-baseurl/branch/master/graph/badge.svg)](https://codecov.io/gh/star-integrations/go-appengine-baseurl)
[![Go Report Card](https://goreportcard.com/badge/github.com/star-integrations/go-appengine-baseurl)](https://goreportcard.com/report/github.com/star-integrations/go-appengine-baseurl)

Google App Engine URL Generator


## Installation

```sh
$ go get github.com/star-integrations/go-appengine-baseurl/baseurl
```

## Configuration

```go
import (
	"github.com/star-integrations/go-appengine-baseurl/baseurl"
)

func main() {
	...
	u, err := baseurl.New(ctx)
	u.String() // https://YOUR_PROJECT_ID.appspot.com

	u, err = baseurl.New(ctx, WithService("api"))
	u.String() // https://api-dot-YOUR_PROJECT_ID.appspot.com

	u, err = baseurl.New(ctx, WithVersion("1"), WithService("api"))
	u.String() // https://1-dot-api-dot-YOUR_PROJECT_ID.appspot.com

	u, err = baseurl.New(ctx, WithProject("SOME_PROJECT_ID"), WithService("api"))
	u.String() // https://api-dot-SOME_PROJECT_ID.appspot.com

	u, err = baseurl.New(ctx, WithDomain("SOME_PROJECT_ID.appspot.com"), WithService("api"))
	u.String() // https://api-dot-SOME_PROJECT_ID.appspot.com
	...
}

```

## LICENSE

MIT
