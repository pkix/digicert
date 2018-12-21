package main

import (
	"digicert"
	"errors"
	"log"
	"os"
)

var c *digicert.Client

func main() {
	container()
	certificate()
	domain()
	orders()
	organization()
	request()
	user()
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

// export DC_KEY="apikeys"
