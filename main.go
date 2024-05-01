package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Domain, hasMX,hasSPF,spfrecords,hasDMACR,dmarcRecord \n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error : could not read from input: %v \n", err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMACR bool
	var spfrecords, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v \n", err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtrecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v \n", err)
	}

	for _, record := range txtrecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfrecords = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error : %v\n", err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMACR = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v", domain, hasMX, hasSPF, spfrecords, hasDMACR, dmarcRecord)

}
