package main

import (
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

func SearchGoogleReviews(args *SummaryForm, conf *ApiKeys) (string, *GooglePlaceReviews, error) { // get 5 reviews for a business
	c, err := maps.NewClient(maps.WithAPIKey(conf.GoogleMaps.Key))
	if err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
	}

	if len(resp.Results) > 0 {
		matchedPlace := resp.Results[0]
		fmt.Println("looking for: " + args.Name + " found: " + matchedPlace.Name)

		r2 := &maps.PlaceDetailsRequest{
			PlaceID: matchedPlace.PlaceID,
		}
		details, err := c.PlaceDetails(context.Background(), r2)
		if err != nil {
			fmt.Println(err)
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

func GetReviews(c *maps.Client, b *maps.PlaceDetailsRequest) ([]maps.PlaceReview, error) { // get reviews for a business
	resp, err := c.PlaceDetails(context.Background(), b)
	if err != nil {
		fmt.Println(err)
	}

	return resp.Reviews, err
}
