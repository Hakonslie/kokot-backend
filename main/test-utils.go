package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"restful/kokots"
)

func testHelperAddKokot(address, name string) (kokots.Kokot, error) {
	nkj, err := json.Marshal(name)
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
