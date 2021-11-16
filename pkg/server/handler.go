package server

import (
	"io/ioutil"
	"log"
	"net/http"
)

type HttpHandler struct{}

func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Printf("Error when reading body %s\n", body)
	}

	log.Printf("Received message %s\n", body)

	if jsonDataOk(body) {
		insertDataInDb(body)
		res.Write([]byte("Data accepted."))
	} else {
		msg := "Data not recognized"
		res.Write([]byte(msg))
		log.Printf("%s\n", msg)
	}
}

func insertDataInDb(jsonData []byte) {
	data := getTemperatureData(jsonData)

	db := Connect()
	if db == nil {
		return
	}
	defer db.Close()
	insert(db, data)
}
