package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"io"
	"math/rand"
	"net/http"
	"restful/kokots"
	"strconv"
	"testing"
)

const address = "http://localhost:5050/kokots/v1"

func TestApp(t *testing.T) {
	go App("5050")
	r, err := http.Get(address)
	require.NoError(t, err)
	require.Equal(t, r.StatusCode, 200)

	b, _ := io.ReadAll(r.Body)
	var ks []kokots.Kokot
	err = json.Unmarshal(b, &ks)
	require.NoError(t, err)
	require.ElementsMatch(t, ks, make([]kokots.Kokot, 0))
}

func TestAddAndGet(t *testing.T) {
	go App("5050")

	res, err := testHelperAddKokot(address, "Louis Litt")
	require.NoError(t, err)
	require.Equal(t, "Louis Litt", res.Name)

	res, err = testHelperGetKokot(address, res.ID)
	require.NoError(t, err)
	require.Equal(t, "Louis Litt", res.Name)

}

func TestIsEmpty(t *testing.T) {
	go App("5050")
	r, err := http.Get(address)
	require.NoError(t, err)
	b, err := io.ReadAll(r.Body)
	require.NoError(t, err)

	var kokts []kokots.Kokot
	err = json.Unmarshal(b, &kokts)
	require.NoError(t, err)

	require.Equal(t, 0, len(kokts))
}

func TestAddMultiple(t *testing.T) {
	go App("5050")
	for i := 0; i < 10; i++ {
		newKokot := strconv.Itoa(rand.Int())
		_, err := testHelperAddKokot(address, newKokot)
		require.NoError(t, err)
	}

	r, err := http.Get(address)
	require.NoError(t, err)
	b, err := io.ReadAll(r.Body)

	var kokts []kokots.Kokot
	json.Unmarshal(b, &kokts)

	require.Equal(t, 10, len(kokts))
}

func TestUpdate(t *testing.T) {
	go App("5050")

	res, err := testHelperAddKokot(address, "Louis Litt")
	require.NoError(t, err)

	res.Name = "Louis Hitt"
	resM, err := json.Marshal(res)
	require.NoError(t, err)
	req, err := http.NewRequest(http.MethodPut, address+"/"+res.ID, bytes.NewReader(resM))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	_, err = http.DefaultClient.Do(req)
	require.NoError(t, err)

	res, err = testHelperGetKokot(address, res.ID)
	require.NoError(t, err)
	require.Equal(t, "Louis Hitt", res.Name)
}

func TestDelete(t *testing.T) {
	go App("5050")

	res, err := testHelperAddKokot(address, "Louis Litt")
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodDelete, address+"/"+res.ID, nil)
	require.NoError(t, err)
	_, err = http.DefaultClient.Do(req)
	require.NoError(t, err)

	res, err = testHelperGetKokot(address, res.ID)
	require.Equal(t, err.Error(), "Wrong status code")
	require.Error(t, err)
	require.Empty(t, res)

}
