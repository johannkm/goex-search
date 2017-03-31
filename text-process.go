package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
)

const (
  ParallelDotsUrl = "https://apis.paralleldots.com/emotion"
)

func ProcessText(creds *ApiKeys) {

  client := &http.Client{}
  req, err := http.NewRequest("GET", "https://gateway.watsonplatform.net/tone-analyzer/api/v3/tone", nil)
  req.SetBasicAuth(creds.Watson.Username, creds.Watson.Password)
  q := req.URL.Query()
  q.Add("version", "2016-05-19")
  q.Add("text", "I love this place! Great for grabbing a bite for lunch.")
  q.Add("sentences", "false")
  req.URL.RawQuery = q.Encode()

  resp, err := client.Do(req)
  if err != nil{
    panic(err)
  }
  bodyText, err := ioutil.ReadAll(resp.Body)
  s := string(bodyText)
  fmt.Println(s)

}

// func ParallelDots(key string){
//   form := url.Values{}
//   form.Add("sentence1", "what is going on here")
//   form.Add("apikey", key)
//   fmt.Println(form.Encode())
//   // resp,err := http.PostForm(ParallelDotsUrl, Post_Format, bytes.NewBuffer([]byte(form.Encode())))
//   resp,err := http.PostForm(ParallelDotsUrl, form)
//   if err != nil {
//     panic(err)
//   }
//   defer resp.Body.Close()
//   body, err := ioutil.ReadAll(resp.Body)
//   // if err != nil {
//   //   panic(err)
//   // }
//   // _, err = io.Copy(os.Stdout, resp.Body)
//   fmt.Println(string(body))
//   if err != nil {
//     log.Fatal(err)
//   }
// }
