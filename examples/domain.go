package main

import (
	"fmt"
	"log"
)

func domain() {
	log.Println("================= DOMAIN =================")
	if err := checkEnv(); err != nil {
		fmt.Println(err)
		return
	}
	domain, err := c.GetDomainControlEmails("432481")
	// domain, err := c.ListValidationTypes()
	// domain, err := c.ListDomains("88217")
	// domain, err := c.ViewADomain("432481")

	if err != nil {
		log.Println(err)
		return
	}
	log.Println(domain)
}
