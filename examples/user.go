package main

import (
	"fmt"
	"log"
)

func user() {
	log.Println("================= USER =================")
	if err := checkEnv(); err != nil {
		fmt.Println(err)
		return
	}
	username := c.CheckUserName("jack.z@ssl.report")
	// user, err := c.ViewUser("1546247")
	// user, err := c.ListUsers("88217")
	user, err := c.ListRoles("92661")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(user)
	log.Println("does that available - ", username)
	log.Println("================= End =================")
}
