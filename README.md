# digicert

[![GoDoc](https://img.shields.io/badge/godoc-reference-5673AF.svg?style=flat-square)](https://godoc.org/github.com/pkix/digicert)
[![Go Report Card](https://goreportcard.com/badge/github.com/pkix/digicert)](https://goreportcard.com/report/github.com/pkix/digicert)

A Go library for interacting with
[Digicert CertCentral's API v2](https://www.digicert.com/services/v2/documentation). This library allows you to:

* Certificate Management.
* Container Management.
* Orders Management.
* Organization Management.
* Request Management.
* User Management.

## Getting Started

```go
package main

import (
	"github.com/pkix/digicert"
	"errors"
	"log"
	"os"
)

var c *digicert.Client

func main() {
	log.Println("================= CONTAINER =================")
	if err := checkEnv(); err != nil {
		fmt.Println(err)
		return
	}

	container, err := c.ViewAContainerOfParent("32000")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(container)
}

func checkEnv() error {
	if c == nil {
		var err error
		c, err = digicert.New(os.Getenv("DC_KEY"))
		if err != nil {
			log.Fatal(err)
		}
	}
	if c.AuthKey == "" {
		return errors.New("API key not defined")
	}
	return nil
}
```

# License

MIT License. See the [LICENSE](LICENSE) file for details.
