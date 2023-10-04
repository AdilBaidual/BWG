package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func checkCurrencyCode(code string) (bool, error) {
	url := "https://openexchangerates.org/api/currencies.json"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := errors.New(fmt.Sprintf("request error to currency api %d", resp.Status))
		log.Fatalf(err.Error())
		return false, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
		return false, err
	}

	var currencyMap map[string]string
	err = json.Unmarshal(body, &currencyMap)
	if err != nil {
		return false, err
	}

	_, exists := currencyMap[code]
	return exists, nil
}
