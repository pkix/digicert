package main

import (
	"fmt"
	"log"
)

func orders() {
	log.Println("================= ORDER =================")
	if err := checkEnv(); err != nil {
		fmt.Println(err)
		return
	}

	orders, err := c.ListOrders(10, 0)
	// order, err := c.ViewOrder("3863157")
	if err != nil {
		log.Println(err)
		return
	}

	for _, order := range orders.Orders {
		// log.Println(order.ID)
		log.Println(order.Certificate.CommonName)
		log.Println(order.Organization)
	}
	// log.Println(order.Certificate.ID)
	// log.Println(order.Certificate.CommonName)
}
