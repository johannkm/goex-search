package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	// "github.com/labstack/echo/middleware"
	"net/http"
	"net/url"
	"os"
)

type SearchForm struct { // general search request from spa
	Term     string `json:"term"`
	Location string `json:"location"`
	Limit string `json:"limit"`
}

type SummaryForm struct { // place summary request from spa
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func Serve() {

	e := echo.New()
	e.File("/", "index.html")
	e.Static("/dist", "dist")
	e.Static("/static", "static")
	e.POST("/places", postPlaces)
	e.POST("/summary", postSummary)

	// e.Use(middleware.CORS()) // TODO: remove for production

	e.Logger.Fatal(e.Start(":"+os.Getenv("PORT")))
}

func postPlaces(c echo.Context) (err error) { // respond to place search post

	var u = new(SearchForm)

	if err = c.Bind(u); err != nil {
		return err
	}

	form := url.Values{}
	form.Add("location", u.Location)
	form.Add("term", u.Term)
	form.Add("limit", u.Limit)

	res, err := YelpSearch(form.Encode(), yelpToken)
	if err != nil {
		fmt.Println(err)
	}

	out, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}

	return c.String(http.StatusOK, string(out))
}

func postSummary(c echo.Context) (err error) { // respond to place summary post
	var u = new(SummaryForm)
	if err = c.Bind(u); err != nil {
		return err
	}

	text, gReviews, err := SearchGoogleReviews(u, apiKeys)
	if err != nil {
		fmt.Println(err)
	}

	res, err := Summarize(text, apiKeys)
	if err != nil {
		fmt.Println(err)
	}

	res.Reviews = gReviews

	out, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}

	return c.String(http.StatusOK, string(out))
}
