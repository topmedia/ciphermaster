package main

import (
	"crypto/tls"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	host        = flag.String("host", "192.168.110.123", "Load Balancer Hostname/IP")
	user        = flag.String("user", "bal", "Administrator username")
	pass        = flag.String("pass", "2fourall", "Administrator password")
	ciphers     = flag.String("cipers", "ECDHE-RSA-AES256-GCM-SHA384", "Ciphers to assign to services, colon-separated")
	skip        = flag.String("skip", "", "If the service matches this string, no ciphers will be changed.")
	listciphers = flag.Bool("listciphers", false, "List available ciphers and exit")
)

func availableCiphers() []string {
	return []string{
		"ECDHE-RSA-AES256-GCM-SHA384",
		"ECDHE-ECDSA-AES256-GCM-SHA384",
		"ECDHE-RSA-AES256-SHA384",
		"ECDHE-ECDSA-AES256-SHA384",
		"ECDHE-RSA-AES256-SHA",
		"ECDHE-ECDSA-AES256-SHA",
		"DHE-DSS-AES256-GCM-SHA384",
		"DHE-RSA-AES256-GCM-SHA384",
		"DHE-RSA-AES256-SHA256",
		"DHE-DSS-AES256-SHA256",
		"DHE-RSA-AES256-SHA",
		"DHE-DSS-AES256-SHA",
		"DHE-RSA-CAMELLIA256-SHA",
		"DHE-DSS-CAMELLIA256-SHA",
		"ECDH-RSA-AES256-GCM-SHA384",
		"ECDH-ECDSA-AES256-GCM-SHA384",
		"ECDH-RSA-AES256-SHA384",
		"ECDH-ECDSA-AES256-SHA384",
		"ECDH-RSA-AES256-SHA",
		"ECDH-ECDSA-AES256-SHA",
		"AES256-GCM-SHA384",
		"AES256-SHA256",
		"AES256-SHA",
		"CAMELLIA256-SHA",
		"ECDHE-RSA-DES-CBC3-SHA",
		"ECDHE-ECDSA-DES-CBC3-SHA",
		"EDH-RSA-DES-CBC3-SHA",
		"EDH-DSS-DES-CBC3-SHA",
		"ECDH-RSA-DES-CBC3-SHA",
		"ECDH-ECDSA-DES-CBC3-SHA",
		"DES-CBC3-SHA",
		"DHE-RSA-SEED-SHA",
		"DHE-DSS-SEED-SHA",
		"SEED-SHA",
		"ECDHE-RSA-RC4-SHA",
		"ECDHE-ECDSA-RC4-SHA",
		"ECDH-ECDSA-RC4-SHA",
		"RC4-SHA",
		"RC4-MD5ECDHE-RSA-AES128-GCM-SHA256",
		"ECDHE-ECDSA-AES128-GCM-SHA256",
		"ECDHE-RSA-AES128-SHA256",
		"ECDHE-ECDSA-AES128-SHA256",
		"ECDHE-RSA-AES128-SHA",
		"ECDHE-ECDSA-AES128-SHA",
		"DHE-DSS-AES128-GCM-SHA256",
		"DHE-RSA-AES128-GCM-SHA256",
		"DHE-RSA-AES128-SHA256",
		"DHE-DSS-AES128-SHA256",
		"DHE-RSA-AES128-SHA",
		"DHE-DSS-AES128-SHA",
		"DHE-RSA-CAMELLIA128-SHA",
		"DHE-DSS-CAMELLIA128-SHA",
		"ECDH-RSA-AES128-GCM-SHA256",
		"ECDH-ECDSA-AES128-GCM-SHA256",
		"ECDH-RSA-AES128-SHA256",
		"ECDH-ECDSA-AES128-SHA256",
		"ECDH-RSA-AES128-SHA",
		"ECDH-ECDSA-AES128-SHA",
		"AES128-GCM-SHA256",
		"AES128-SHA256",
		"AES128-SHA",
		"CAMELLIA128-SHA",
	}
}

type Data struct {
	VSList []VS `xml:"Success>Data>VS"`
}

type VS struct {
	Index           int    `xml:"Index"`
	Nickname        string `xml:"NickName"`
	SSLAcceleration string `xml:"SSLAcceleration"`
}

func RESTURL(path string) string {
	return fmt.Sprintf("https://%s:%s@%s/access/%s", *user, *pass, *host, path)
}

func main() {
	flag.Parse()

	if *listciphers {
		fmt.Println("Available ciphers for use with -ciphers:")
		for _, cipher := range availableCiphers() {
			fmt.Println(cipher)
		}
		os.Exit(0)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	res, err := client.Get(RESTURL("listvs"))

	if err != nil {
		log.Fatalf("Error fetching available virtual services: %v", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	bodyFixed := strings.Replace(string(body), "ISO-8859-1", "UTF-8", 1)

	if err != nil {
		log.Fatalf("Error fetching available virtual services: %v", err)
	}

	var data Data
	err = xml.Unmarshal([]byte(bodyFixed), &data)

	for _, vs := range data.VSList {
		if vs.SSLAcceleration == "N" {
			log.Printf("Skipping non-SSL Service %s", vs.Nickname)
			continue
		}

		if *skip != "" && strings.Contains(vs.Nickname, *skip) {
			log.Printf("Skipping Service %s because it matches -skip", vs.Nickname)
			continue
		}

		log.Printf("Setting ciphers of %s to %s", vs.Nickname, *ciphers)

		res, err = client.Get(RESTURL(fmt.Sprintf("modvs?vs=%d&Ciphers=%s", vs.Index, *ciphers)))

		if err != nil {
			log.Fatalf("Error modifying VS %s: %v", vs.Nickname, err)
		}
	}
}
