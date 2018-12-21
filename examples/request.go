package main

import (
	"fmt"
	"log"
)

func request() {
	log.Println("================= REQUEST =================")
	if err := checkEnv(); err != nil {
		fmt.Println(err)
		return
	}

	// req, err := c.ListRequests("")
	req, err := c.ViewRequest("3108958")

	if err != nil {
		log.Println(err)
		return
	}
	log.Println(req)
}
