package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Lol struct {
	XMLName xml.Name `xml:"lol"`
	Content string   `xml:",chardata"`
	OpenVPNData
}

type OpenVPNData struct {
	Key     string `xml:"key"`
	Cert    string `xml:"cert"`
	Ca      string `xml:"ca"`
	TlsAuth string `xml:"tls-auth"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Print(`Usage:
	ovpnsplit <file.ovpn>
`)
		os.Exit(1)
	}
	filepath := os.Args[1]

	filedata, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal("Couldn't open file", err)
		os.Exit(2)
	}

	data := Lol{}
	// Ugly hack. Otherwise, not considered as xml valid…
	filedata = append(append([]byte("<lol>\n"), filedata...), []byte("\n</lol>")...)
	xml.Unmarshal(filedata, &data)

	if data.Key != "" {
		ioutil.WriteFile(filepath+".key.pem", []byte(data.Key), 0600)
	}

	if data.Cert != "" {
		ioutil.WriteFile(filepath+".cert.pem", []byte(data.Cert), 0600)
	}

	if data.Ca != "" {
		ioutil.WriteFile(filepath+".ca.pem", []byte(data.Ca), 0600)
	}

	if data.TlsAuth != "" {
		ioutil.WriteFile(filepath+".tls-auth.pem", []byte(data.TlsAuth), 0600)
	}

	newnetmanfiledata := append(append(append(append([]byte(data.Content), []byte("\ncert "+filepath+".cert.pem\n")...), []byte("ca "+filepath+".ca.pem\n")...), []byte("key "+filepath+".key.pem\n")...), []byte("tls-auth "+filepath+".tls-auth.pem\n")...)
	ioutil.WriteFile(filepath+".new.ovpn", newnetmanfiledata, 0600)
}
