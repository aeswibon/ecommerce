package cmd

import (
	"log"
	"net"
	"time"
)

// WaitForService waits for a service to start
func WaitForService(host string) {
	log.Printf("Waiting for service %s to start", host)
	for {
		log.Println("Trying to connect to service...")
		if conn, err := net.Dial("tcp", host); err == nil {
			conn.Close()
			log.Printf("Service %s started", host)
			return
		}
		time.Sleep(1 * time.Second)
	}
}