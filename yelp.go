package main

import (
  "io/ioutil"
  "bytes"
  "encoding/json"
  "net/http"
  "net/url"
)

const (
  Conf_Path = "yelp-config.json"
  Token_Url = "https://api.yelp.com/oauth2/token"
  Post_Format = "application/x-www-form-urlencoded"
)

type ApiKeys struct {
  Id string `json:"client_id"`
  Secret string `json:"client_secret"`
}

type ApiToken struct {
  AccessToken string `json:"access_token"`
  TokenType string `json:"token_type"`
  ExpiresIn int64 `json:"expires_in"`
}

type Business struct {
  Rating float32 `json:"rating"`
  Price string `json:"price"`
  Id string `json:"id"`
  Name string `json:"name"`
  Url string `json:"url"`
  Phone string `json:"phone"`
  RatingCount int32 `json:"review_count"`
  ImgUrl string `json:"image_url"`
  Hours []struct {
    HoursType string `json:"hours_type"`
    Open []struct {
      Day int32   `json:"day"`
      Start string `json:"start"`
      End string `json:"end"`
      Overnight bool `json:"is_overnight"`
    } `json:"open"`
  } `json:"hours"`
  Categories []struct {
    Alias string `json:"alias"`
    Title string `json:"title"`
  } `json:"categories"`
  Coordinates struct {
    Latitude float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
  } `json:"coordinates"`
}

type SearchResp struct {
  Total int64 `json:"total"`
  Businesses []Business `json:"businesses"`
}


func GetApiToken() (*ApiToken, error){

  conf, err := ReadConfigFile(Conf_Path)
  if err != nil {
    return nil, err
  }

  form := url.Values{}
  form.Add("grant_type", "client_credentials")
  form.Add("client_id", conf.Id)
  form.Add("client_secret", conf.Secret)

  resp,err := PostId(Token_Url, Post_Format, form.Encode())
  if err != nil {
    panic(err)
  }
  cred,err := ParseApiToken(resp)
  return cred, err
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

func PostId(url string, format string, args string) ([]byte, error){

  argsBytes := bytes.NewBuffer( []byte(args) )
  resp,err := http.Post(url, format, argsBytes)

  if err != nil {
    return nil, err
  }

  defer resp.Body.Close()
  bodyBytes, err := ioutil.ReadAll(resp.Body)

  return []byte(bodyBytes), err
}


func ParseApiToken(body []byte) (*ApiToken, error){

  var c = new(ApiToken)
  err := json.Unmarshal(body, &c)

  return c, err
}


func MakeGet(url string, cred *ApiToken) (*http.Response, error) {

  auth := cred.TokenType+" "+cred.AccessToken

  client := &http.Client{}
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return nil, err
  }
  req.Header.Set("Authorization", auth)
  res, err := client.Do(req)

  return res, err

}


func YelpSearch(args string, cred *ApiToken)(*SearchResp, error){

  addr := "https://api.yelp.com/v3/businesses/search"+"?"+args
  resp, err := MakeGet(addr, cred)
  if err != nil {
    return nil, err
  }
  bodyBytes, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  var r = new(SearchResp)
  err_ := json.Unmarshal( []byte(bodyBytes) , &r)

  return r, err_

}
