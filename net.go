package main

import (
	"net/http"
	"crypto/tls"
	"io/ioutil"
	"os"
	"fmt"
)

//NOTE, EVENTUALLY DO SSL
//Also, rename this from net >__>
func update(name string) {
	tr := &http.Transport {
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Get(fmt.Sprintf("http://padherder.com/api/%s/",name))

	if err != nil {	panic (err) }

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil { panic(err) }

	//println(string(body))
	//println(string(*resp))

	fo, err := os.Create(name)
	if err != nil { panic(err) }

	defer fo.Close()

	fo.Write(body)

	fo.Sync()
}


func read(name string) []byte {
	fo, err := ioutil.ReadFile(name)
	if err != nil { panic(err) }

	return fo
}

func getMon(ID int) []byte {
	tr := &http.Transport {
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Get(fmt.Sprintf("http://padherder.com/user-api/monster/%d/",ID))

	if err != nil {	panic (err) }

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil { panic(err) }

	return body
}
