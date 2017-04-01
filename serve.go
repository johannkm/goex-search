package main

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"net/url"
	"fmt"
)

type SearchForm struct {
	Term     string `json:"term"`
	Location string `json:"location"`
}

type SummaryForm struct {
	Name string `json:"name"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func Serve() {

	e := echo.New()
	e.File("/", "index.html")
	e.Static("/dist", "dist")
	e.Static("/static", "static")
	e.POST("/places", postPlaces)
	e.POST("/summary", postSummary)

	e.Use(middleware.CORS()) // TODO: remove for production

	e.Logger.Fatal(e.Start(":8000"))
}

func postPlaces(c echo.Context) (err error) {

	var u = new(SearchForm)

	if err = c.Bind(u); err != nil {
		return err
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

func postSummary(c echo.Context) (err error) {
	var u = new(SummaryForm)
	if err = c.Bind(u); err != nil {
		return err
	}

	fmt.Println(u)

	text, err := SearchGoogleReviews(u, apiKeys)
	if err != nil {
		panic(err)
	}
	// reviews, err := YelpReviewSearch(u.BusinessId, yelpToken)
	// if err!=nil{
	// 	panic(err)
	// }


	//
	// fmt.Println(reviews)
	// text := ""
	// for x := range reviews.Reviews {
	// 	text = text + reviews.Reviews[x].Text
	// 	fmt.Println(x)
	// }

	res, err := Summarize(text, apiKeys)
	if err != nil {
		panic(err)
	}

	out, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	return c.String(http.StatusOK, string(out))
}
