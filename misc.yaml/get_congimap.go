package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type iface interface {
}

// It's a short code for getting config map from /var/lib/myconfig
func main() {
	fd, err := os.Open("/var/lib/myconfig/mongoHosts")
	defer func() {
		if r := recover(); r != nil {
			log.Println("error: ", r)
		}
		log.Println("close file")
		fd.Close()
	}()

	if err != nil {
		log.Println(err)
	} else {
		c, err := ioutil.ReadAll(fd)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(c))
	}

	http.HandleFunc("/health", func(res http.ResponseWriter, req *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			log.Println(err)
		}
		log.Println("get request, ", hostname)
		res.WriteHeader(200)
	})

	log.Fatal(http.ListenAndServe(":9999", nil))
}
