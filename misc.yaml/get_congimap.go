package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// It's a short code for getting config map from /var/lib/myconfig
func main() {
	fd, err := os.Open("/var/lib/myconfig")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	c, err := ioutil.ReadAll(fd)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(c))
	log.Fatal(http.ListenAndServe(":9999", nil))
}
