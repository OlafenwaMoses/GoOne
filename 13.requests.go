package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	resp, err1 := http.Get("http://worldtimeapi.org/api/timezone/America/Argentina/Salta")

	if err1 != nil {
		fmt.Println("An error occured")
	}
	fmt.Println(resp)
	fmt.Println("====================")

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(string(body))
	}

	resp.Body.Close()

}
