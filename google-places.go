package main

import (
	"log"

	"github.com/kr/pretty"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"

	"fmt"
)

type GooglePlaceReview struct {
	AuthorName string `json:"author_name"`
	AuthorURL  string `json:"author_url"`
	Rating     int    `json:"rating"`
	Text       string `json:"text"`
	Time       int    `json:"time"`
}

type GooglePlaceReviews struct {
	Reviews []*GooglePlaceReview `json:"reviews"`
}

func SearchGoogleReviews(args *SummaryForm, conf *ApiKeys) (string, *GooglePlaceReviews, error) {
	c, err := maps.NewClient(maps.WithAPIKey(conf.GoogleMaps.Key))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	fmt.Println(args)

	r := &maps.NearbySearchRequest{
		Location: &maps.LatLng{
			Lat: args.Latitude,
			Lng: args.Longitude,
		},
		Keyword: args.Name,
		RankBy:  "distance",
	}

	resp, err := c.NearbySearch(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	if len(resp.Results) > 0 {
		matchedPlace := resp.Results[0]
		fmt.Println("looking for: " + args.Name + " found: " + matchedPlace.Name)

		r2 := &maps.PlaceDetailsRequest{
			PlaceID: matchedPlace.PlaceID,
		}
		details, err := c.PlaceDetails(context.Background(), r2)
		if err != nil {
			panic(err)
		}

		gReviews := new(GooglePlaceReviews)

		for x := range details.Reviews {
			gReview := new(GooglePlaceReview)
			gReview.AuthorName = details.Reviews[x].AuthorName
			gReview.AuthorURL = details.Reviews[x].AuthorURL
			gReview.Rating = details.Reviews[x].Rating
			gReview.Text = details.Reviews[x].Text
			gReview.Time = details.Reviews[x].Time
			gReviews.Reviews = append(gReviews.Reviews, gReview)
		}

		reviews := ""
		for x := range details.Reviews {
			reviews += details.Reviews[x].Text + " . "
		}
		return reviews, gReviews, nil

	}
	return "", nil, err
}

func RunTraining(conf *ApiKeys) {
	CollectPlaces(conf)
}

func CollectPlaces(conf *ApiKeys) {

	c, err := maps.NewClient(maps.WithAPIKey(conf.GoogleMaps.Key))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &maps.NearbySearchRequest{
		Location: &maps.LatLng{
			Lat: 39.4142688,
			Lng: -77.4105409,
		},
		Radius: 5000,
	}
	resp, err := c.NearbySearch(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	var reviewers = make(map[string]int32)

	for i := 0; i < len(resp.Results); i++ {
		pretty.Println(resp.Results[i].PlaceID)
		b := &maps.PlaceDetailsRequest{
			PlaceID: resp.Results[i].PlaceID,
		}
		revs, err := GetReviews(c, b)
		for j := 0; j < len(revs); j++ {
			pretty.Println(revs[j].AuthorURL)
			reviewers[revs[j].AuthorURL] = 1 + reviewers[revs[j].AuthorURL]
		}
		if err != nil {
			log.Fatalf("fatal error: %s", err)
		}
	}
	pretty.Println(reviewers)

}

func GetReviews(c *maps.Client, b *maps.PlaceDetailsRequest) ([]maps.PlaceReview, error) {
	resp, err := c.PlaceDetails(context.Background(), b)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	return resp.Reviews, err
}
