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

func getAstros() ([]byte, error) {
	url := "http://api.open-notify.org/astros.json"
	spaceClient := http.Client{
		Timeout: time.Second * 2, // 2 seconds
	}

	// Call Online API
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Set("User-Agent", "my-golang-agent")

	res, err := spaceClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
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
		return nil, err
	}

	return body, nil
}

func main() {
	/*
		text := `{"people": [{"craft": "ISS", "name": "Sergey Rizhikov"}, {"craft": "ISS",
		 "name": "Andrey Borisenko"}, {"craft": "ISS", "name": "Shane Kimbrough"}, {"craft":
		 "ISS", "name": "Oleg Novitskiy"}, {"craft": "ISS", "name": "Thomas Pesquet"}, {"craft":
		 "ISS", "name": "Peggy Whitson"}], "message": "success", "number": 6}`
		textBytes := []byte(text)
	*/

	p := spacePeople{}
	body, err := getAstros()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(body, &p)
	if err != nil {
		log.Fatalf("unable to parse value: %q, error: %s",
			string(body), err.Error())
		return
	}
	fmt.Printf("%+v\n", p)

	fmt.Printf("There are %d astronauts on space\n", p.Number)
	for _, a := range p.People {
		fmt.Printf("Astronaut - %s on %s \n", a.Name, a.Craft)
	}
}
