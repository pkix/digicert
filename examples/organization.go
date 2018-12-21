package main

import (
	"fmt"
	"log"
)

func organization() {
	log.Println("================= ORGANIZATION =================")
	if err := checkEnv(); err != nil {
		fmt.Println(err)
		return
	}

	orgs, err := c.ListAllOrganizations()
	if err != nil {
		log.Println(err)
	}
	log.Println(orgs)

}
