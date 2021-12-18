package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	First string
}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := person{First: "krzys"}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println("Encoded bad data", err)
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	var p1 person

	err := json.NewDecoder(r.Body).Decode(&p1)
	if err != nil {
		log.Println("Decoded bad data", err)
	}
	log.Println("Decoded data person: ", p1)
}

func main() {

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)

	http.ListenAndServe(":8082", nil)

	/* 	p1 := person{First: "krzys"}
	   	p2 := person{First: "janusz"}

	   	xp := []person{p1, p2}

	   	bs, err := json.Marshal(xp)
	   	if err != nil {
	   		log.Panic(err)
	   	}
	   	fmt.Println("marshal: ", string(bs))

	   	xp2 := []person{}

	   	err = json.Unmarshal(bs, &xp2)
	   	if err != nil {
	   		log.Panic(err)
	   	}

	   	fmt.Println("basic data: ", xp2) */

}
