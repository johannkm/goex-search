package main

import(
  // "strconv"
  "net/url"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "net/http"
  "encoding/json"
  // "fmt"
)

var token *ApiToken

type SearchForm struct {
  Term string `json:"term"`
  Location string `json:"location"`
}

func get(c echo.Context)(err error){

  var u = new(SearchForm)

  if err = c.Bind(u); err != nil {
    return
  }

  form := url.Values{}
  form.Add("location", u.Location)

  res, err := YelpSearch(form.Encode(), token)
  if err != nil {
    panic(err)
  }

  out, err := json.Marshal(res)
  if err != nil {
      panic (err)
  }

  return c.String(http.StatusOK, string(out))
}

func main() {

  var err error
  token, err = GetApiToken()
  if err != nil {
    panic(err)
  }

  e := echo.New()
  e.File("/", "index.html")
  e.Static("/dist", "dist")
  e.Static("/static", "static")
  e.POST("/places", get)
  // e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
  e.Use(middleware.CORS()) // TODO: remove for production

  e.Logger.Fatal( e.Start(":8000") )
}
