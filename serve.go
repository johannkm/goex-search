package main

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"net/url"
)

type SearchForm struct {
	Term     string `json:"term"`
	Location string `json:"location"`
}

func Serve() {

	e := echo.New()
	e.File("/", "index.html")
	e.Static("/dist", "dist")
	e.Static("/static", "static")
	e.POST("/places", post)

	e.Use(middleware.CORS()) // TODO: remove for production

	e.Logger.Fatal(e.Start(":8000"))
}

func post(c echo.Context) (err error) {

	var u = new(SearchForm)

	if err = c.Bind(u); err != nil {
		return
	}

	form := url.Values{}
	form.Add("location", u.Location)
	form.Add("term", u.Term)


	res, err := YelpSearch(form.Encode(), yelpToken)
	if err != nil {
		panic(err)
	}

	out, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	return c.String(http.StatusOK, string(out))
}
