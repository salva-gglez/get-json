package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type astronaut struct {
	Craft string `json:"craft"`
	Name  string `json:"name"`
}

type spacePeople struct {
	People  []astronaut `json:"people"`
	Message string      `json:"message"`
	Number  int         `json:"number"`
}

func main() {
	/*
		text := `{"people": [{"craft": "ISS", "name": "Sergey Rizhikov"}, {"craft": "ISS",
		 "name": "Andrey Borisenko"}, {"craft": "ISS", "name": "Shane Kimbrough"}, {"craft":
		 "ISS", "name": "Oleg Novitskiy"}, {"craft": "ISS", "name": "Thomas Pesquet"}, {"craft":
		 "ISS", "name": "Peggy Whitson"}], "message": "success", "number": 6}`
		textBytes := []byte(text)
	*/

	url := "http://api.open-notify.org/astros.json"
	spaceClient := http.Client{
		Timeout: time.Second * 2, // 2 seconds
	}

	// Call Online API
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("User-Agent", "my-golang-agent")

	res, err := spaceClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	if res.StatusCode != 200 {
		fmt.Printf("Response error from server (%d)\n", res.StatusCode)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	p := spacePeople{}
	err = json.Unmarshal(body, &p)
	if err != nil {
		log.Fatalf("unable to parse value: %q, error: %s",
			string(body), err.Error())
		return
	}
	fmt.Printf("%+v\n", p)
}
