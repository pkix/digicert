package main

import (
	"fmt"
	"log"
)

func container() {
	log.Println("================= CONTAINER =================")
	if err := checkEnv(); err != nil {
		fmt.Println(err)
		return
	}

	container, err := c.ViewAContainerOfParent("92661")
	// container, err := c.ListContainerTempaltes("88217")
	// pem, err := c.DownloadCertificate("3418073")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(container)

	// container, err := c.ViewAContainerTempl("88217", "6")
	// // pem, err := c.DownloadCertificate("3418073")
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// log.Println(container)
}
