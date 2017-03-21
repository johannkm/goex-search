package main

import(
  "fmt"
)

func main() {

    token, err := GetApiToken()
    if err != nil {
      panic(err)
    }

    res, err := YelpSearch("location=Frederick, MD", token)
    fmt.Println(res,err)

}
