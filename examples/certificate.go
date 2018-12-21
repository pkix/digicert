package main

import (
	"fmt"
	"log"
)

func certificate() {
	log.Println("================= CERTIFICATE =================")
	if err := checkEnv(); err != nil {
		fmt.Println(err)
		return
	}
	// pem, err := c.ListAPIKeys()
	// pem, err := c.ListEmailValidations("4119921")
	pem, err := c.OrderStatus(5000)
	// pem, err := c.ListOrganizations("88117")
	// pem, err := c.DownloadPKCS7Certificate("2261880")
	// pem, err := c.ListDuplicateCertificates("3924340")
	// pem, err := c.DownloadCertificate("3118073")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(pem)
}
