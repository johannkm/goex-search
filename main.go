package main

import (
	"encoding/json"
	"io/ioutil"
)

var yelpToken *ApiToken // yelp authenticator
var apiKeys *ApiKeys

const (
	Conf_Path = "api-config.json" // path to api key file
)

type ApiKeys struct { // hold api credentials
	Yelp struct {
		Id     string `json:"client_id"`
		Secret string `json:"client_secret"`
	} `json:"yelp"`
	GoogleMaps struct {
		Key string `json:"key"`
	} `json:"google-maps"`
	Watson struct {
		ToneAnalysis struct {
			Username string `json: password`
			Password string `json:"password"`
		} `json:"tone_analysis"`
		LanguageUnderstanding struct {
			Username string `json: password`
			Password string `json:"password"`
		} `json:"language_understanding"`
	} `json:"watson"`
	TextApi struct {
		AppId string `json:"app_id"`
		Key   string `json:"key"`
	} `json:"text-analysis"`
}

func main() {

	conf, err := ReadConfigFile(Conf_Path) // parse api key file
	if err != nil {
		panic(err)
	}
	apiKeys = conf

	yelpToken, err = GetApiToken(conf) // get yelp authenticator
	if err != nil {
		panic(err)
	}

	Serve() // start server

}

func ReadConfigFile(path string) (*ApiKeys, error) { // parse api keys from json

	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var a = new(ApiKeys)
	err_ := json.Unmarshal([]byte(configFile), &a)
	return a, err_
}
