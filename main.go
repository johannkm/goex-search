package main

import (
	"encoding/json"
	"io/ioutil"
)

var yelpToken *ApiToken
var apiKeys *ApiKeys

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
	ParallelDots struct {
		Key string `json:"key"`
	} `json:"paralleldots"`
	Watson struct {
		Username string `json: password`
		Password string `json:"password"`
	} `json:"watson"`
	TextApi struct {
		AppId string `json:"app_id"`
		Key string `json:"key"`
	} `json:"text-analysis"`
}

func main() {

	conf, err := ReadConfigFile(Conf_Path)
	if err != nil {
		panic(err)
	}
	apiKeys = conf

	yelpToken, err = GetApiToken(conf)
	if err != nil {
		panic(err)
	}

	Serve()
	// RunTraining( conf )
	// ProcessText(conf)

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
