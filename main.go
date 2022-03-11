package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

func PortScan(host string) string {
	wg := sync.WaitGroup{}
	listTCPAddr := []string{}
	re := regexp.MustCompile(`(?:[0-9]+)$`)

	// TCP Ports
	for i := 1; i <= 65535; i++ {
		address := fmt.Sprintf("%s:%d", host, i)

		wg.Add(1)
		go func() {
			defer wg.Done()
			if CheckTCPConnection(address, 5) {
				listTCPAddr = append(listTCPAddr, address)
			}

		}()
	}
	wg.Wait()

	openTCPPorts := []string{}
	for _, j := range listTCPAddr {
		portNumbers := re.FindAllString(j, -1)
		openTCPPorts = append(openTCPPorts, portNumbers...)
	}

	csvTCPPorts := fmt.Sprint(strings.Join(openTCPPorts, ","))
	fmt.Printf("[+] Open TCP Ports are: %s\n", csvTCPPorts)

	return csvTCPPorts
}

func CheckTCPConnection(address string, timeout int) bool {
	_, err := net.DialTimeout("tcp", address, time.Second*time.Duration(timeout))
	return err == nil
}

func CheckUDPConnection(address string, timeout int) bool {
	_, err := net.DialTimeout("udp", address, time.Second*time.Duration(timeout))
	return err == nil
}

func main() {
	ip := flag.String("i", "", "IP Address")
	flag.Parse()

	// Print Usage
	if os.Args[1] == "-h" || os.Args[1] == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if len(os.Args) < 1 {
		flag.PrintDefaults()
		fmt.Println("[-] Check number of arguments!")
		os.Exit(1)
	}

	fmt.Println("[i] Scanning in progress...")
	PortScan(*ip)

	fmt.Println("[i] Done!")

}
