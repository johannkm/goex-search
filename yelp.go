package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	Token_Url   = "https://api.yelp.com/oauth2/token"
	Post_Format = "application/x-www-form-urlencoded"
)

type ApiToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

type Business struct {
	Rating      float32 `json:"rating"`
	Price       string  `json:"price"`
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Url         string  `json:"url"`
	Phone       string  `json:"phone"`
	RatingCount int32   `json:"review_count"`
	ImgUrl      string  `json:"image_url"`
	Hours       []struct {
		HoursType string `json:"hours_type"`
		Open      []struct {
			Day       int32  `json:"day"`
			Start     string `json:"start"`
			End       string `json:"end"`
			Overnight bool   `json:"is_overnight"`
		} `json:"open"`
	} `json:"hours"`
	IsClosed   bool `json:"is_closed"`
	Categories []struct {
		Alias string `json:"alias"`
		Title string `json:"title"`
	} `json:"categories"`
	Coordinates struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"coordinates"`
	Analysis struct {
		Emotion string `json:"emotion"`
	} `json:"analysis"`
}

type YelpSearchResp struct {
	Total      int64      `json:"total"`
	Businesses []Business `json:"businesses"`
}

type YelpReviewSearchResp struct {
	Total   int32 `json:"total"`
	Reviews []struct {
		Text        string `json:"text"`
		Url         string `json:"url"`
		Rating      int32  `json:"rating"`
		TimeCreated string `json:"time_created"`
		User        struct {
			Name     string `json:"name"`
			ImageUrl string `json:"image_url"`
		} `json:"user"`
	} `json:"reviews"`
}

func GetApiToken(conf *ApiKeys) (*ApiToken, error) {

	form := url.Values{}
	form.Add("grant_type", "client_credentials")
	form.Add("client_id", conf.Yelp.Id)
	form.Add("client_secret", conf.Yelp.Secret)

	resp, err := PostId(Token_Url, Post_Format, form.Encode())
	if err != nil {
		panic(err)
	}
	cred, err := ParseApiToken(resp)
	return cred, err
}

func PostId(url string, format string, args string) ([]byte, error) {

	argsBytes := bytes.NewBuffer([]byte(args))
	resp, err := http.Post(url, format, argsBytes)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	return []byte(bodyBytes), err
}

func ParseApiToken(body []byte) (*ApiToken, error) {

	var c = new(ApiToken)
	err := json.Unmarshal(body, &c)

	return c, err
}

func MakeGet(url string, cred *ApiToken) (*http.Response, error) {

	auth := cred.TokenType + " " + cred.AccessToken

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", auth)
	res, err := client.Do(req)

	return res, err

}

func YelpSearch(args string, cred *ApiToken) (*YelpSearchResp, error) {

	addr := "https://api.yelp.com/v3/businesses/search" + "?" + args
	resp, err := MakeGet(addr, cred)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r = new(YelpSearchResp)
	err_ := json.Unmarshal([]byte(bodyBytes), &r)

	return r, err_

}

func YelpReviewSearch(id string, cred *ApiToken) (*YelpReviewSearchResp, error) {
	addr := "https://api.yelp.com/v3/businesses/" + id + "/reviews"

	resp, err := MakeGet(addr, cred)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r = new(YelpReviewSearchResp)
	err_ := json.Unmarshal([]byte(bodyBytes), &r)

	return r, err_
}
