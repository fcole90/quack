package services

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func ComposeDuckDNSURI(domain string, token string) string {
	return fmt.Sprintf("https://www.duckdns.org/update?domains=%s&token=%s", domain, token)
}

func Update(domain string, token string) (string, error) {
	resp, err := http.Get(ComposeDuckDNSURI(domain, token))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	return string(body), err
}
