[![Travis CI](https://img.shields.io/travis/AntoineAugusti/updown/master.svg?style=flat-square)](https://travis-ci.org/AntoineAugusti/updown)
[![Software License](https://img.shields.io/badge/License-MIT-orange.svg?style=flat-square)](https://github.com/AntoineAugusti/updown/blob/master/LICENSE.md)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/AntoineAugusti/updown)
[![Coverage Status](https://img.shields.io/codecov/c/github/antoineaugusti/updown.svg?style=flat-square)](http://codecov.io/github/AntoineAugusti/updown?branch=master)

# Updown Go client
This is a Go client for [updown.io](https://updown.io). Updown lets you monitor websites and online services for an affordable price.

## Installation
Once you have a working Go installation locally, you can grab this package with the following command:
```
go get github.com/antoineaugusti/updown
```

## Documentation
Head over to the [Go documentation](https://godoc.org/github.com/AntoineAugusti/updown) to see available methods and models and to https://updown.io/api for the Updown API.

### Creating a client
The client will be required to perform all actions against the API.
```go
package main

import (
    "github.com/antoineaugusti/updown"
)

func main() {
    // Your API key can be retrieved at https://updown.io/settings/edit
    // You can give a custom HTTP client
    client := updown.NewClient("your-api-key", nil)
}
```

### Listing all checks
```go
result, HTTPResponse, err := client.Check.List()
```

### Getting an Updown token for a check's alias
```go
name := "Google"
token, err := client.Check.TokenForAlias(name)
```
This method returns results from a memory cache by default if it's available. The first time, a request against the API will be performed.

### Getting a check by its token
```go
token := "foo"
result, HTTPResponse, err := client.Check.Get(token)
```

### Getting downtimes for a check
```go
token, page := "foo", 1 // 100 results per page
result, HTTPResponse, err := client.Downtime.List(token, page)
```

### Adding a new check
```go
// See the struct for additional parameters
item := updown.CheckItem{URL: "https://google.fr"}
result, HTTPResponse, err := client.Check.Add(item)
```

### Updating a check
```go
token := "foo"
// See the struct for additional parameters
updated := updown.CheckItem{URL: "https://google.com"}
result, HTTPResponse, err := client.Check.Update(token, updated)
```

### Removing a check
```go
token := "foo"
result, HTTPResponse, err := client.Check.Remove(token)
```

### Getting metrics for a check
```go
token, group := "foo", "host"
from, to := "2016-04-01 00:00:00 +0200", "2016-04-15 00:00:00 +0200"
result, HTTPResponse, err := client.Metric.List(token, group, from, to)
```
