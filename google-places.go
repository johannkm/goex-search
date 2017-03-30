package main

import (
	"log"

	"github.com/kr/pretty"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

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