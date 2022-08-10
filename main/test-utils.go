package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"restful/kokots"
	"time"
)

func runAppWithChannel() chan int {
	// Sleep to give time for the channel and server to close in between tests
	time.Sleep(time.Second / 4)
	c := make(chan int)
	go Run("5050", c)
	return c
}

func testHelperAddKokot(address, name string) (kokots.Kokot, error) {
	nkj, err := json.Marshal(name)
	fmt.Println(string(nkj))
	if err != nil {
		return kokots.Kokot{}, err
	}
	r, err := http.Post(address, "application/json", bytes.NewReader(nkj))
	if err != nil {
		return kokots.Kokot{}, err
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return kokots.Kokot{}, err
	}

	var addedResponse kokots.Kokot
	err = json.Unmarshal(b, &addedResponse)
	if err != nil {
		return kokots.Kokot{}, err
	}

	return addedResponse, nil
}
func testHelperGetKokot(address, id string) (kokots.Kokot, error) {
	r, err := http.Get(address + "/" + id)
	if err != nil {
		return kokots.Kokot{}, err
	} else if r.StatusCode != 200 {
		return kokots.Kokot{}, errors.New("Wrong status code")
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return kokots.Kokot{}, err
	}
	var k kokots.Kokot
	err = json.Unmarshal(b, &k)
	if err != nil {
		return kokots.Kokot{}, err
	}
	return k, nil
}
