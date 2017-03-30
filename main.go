package main

import (
	"encoding/json"
	"io/ioutil"
)

var yelpToken *ApiToken

const (
	Conf_Path = "api-config.json"
)

type ApiKeys struct {
	Yelp struct {
		Id     string `json:"client_id"`
		Secret string `json:"client_secret"`
	} `json:"yelp"`
	GoogleMaps struct {
		Key string `json:"key"`
	} `json:"google-maps"`
}

func main() {

	conf, err := ReadConfigFile(Conf_Path)
	if err != nil {
		panic(err)
	}

	yelpToken, err = GetApiToken(conf)
	if err != nil {
		panic(err)
	}

	// Serve()
	// RunTraining( conf )
	ProcessText()

}

func ReadConfigFile(path string) (*ApiKeys, error) {

	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var a = new(ApiKeys)
	err_ := json.Unmarshal([]byte(configFile), &a)
	return a, err_
}
