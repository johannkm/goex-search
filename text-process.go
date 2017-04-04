package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	textapi "github.com/AYLIEN/aylien_textapi_go"
	"io/ioutil"
	"net/http"
)

var TextApiClient *textapi.Client

type WatsonToneResponse struct {
	DocumentTone struct {
		ToneCategories []struct {
			CategoryName string `json:"category_name"`
			CategoryId   string `json:"category_id"`
			Tones        []struct {
				Score    float64 `json:"score"`
				ToneId   string  `json:"tone_id"`
				ToneName string  `json:"tone_name"`
			} `json:"tones"`
		} `json:"tone_categories"`
	} `json:"document_tone"`
}

type SummaryResponse struct {
	Keyword          string              `json:"keyword"`
	Text             string              `json:"text"`
	KeywordSentiment float32             `json:"keyword_sentiment"`
	Reviews          *GooglePlaceReviews `json:"google_place_review"`
}

type WatsonUnderstandRequest struct {
	Text     string `json:"text"`
	Features struct {
		Keywords struct {
			Emotion   bool  `json:"emotion"`
			Sentiment bool  `json:"sentiment"`
			Limit     int32 `json:"limit"`
		} `json:"keywords"`
	} `json:"features"`
}

type WatsonUnderstandResponse struct {
	Keywords []struct {
		Text      string `json:"text"`
		Sentiment struct {
			Score float32 `json:"score"`
		} `json:"sentiment"`
	} `json:"keywords"`
}

func Summarize(reviews string, creds *ApiKeys) (resp *SummaryResponse, err error) { // return summary for business reviews

	if TextApiClient == nil {
		auth := textapi.Auth{creds.TextApi.AppId, creds.TextApi.Key}
		TextApiClient, err = textapi.NewClient(auth, true)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("incoming review text: " + reviews)
	summarizeParams := &textapi.SummarizeParams{
		Title:             "Review",
		Text:              reviews,
		NumberOfSentences: 1,
		Mode:              "short",
	}
	summary, err := TextApiClient.Summarize(summarizeParams)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", TextApiClient.RateLimits)

	resp = new(SummaryResponse)
	summaryText := ""
	for x := range summary.Sentences {
		summaryText = summaryText + summary.Sentences[x]
	}

	understanding, err := FindKeyword(creds, reviews)
	if err != nil {
		fmt.Println(err)
	}

	title := understanding.Keywords[0].Text
	keywordSentiment := understanding.Keywords[0].Sentiment.Score
	fmt.Println("title: " + title)

	resp.Text = summaryText
	resp.Keyword = title
	resp.KeywordSentiment = keywordSentiment

	fmt.Println("outgoing summary" + summaryText)
	return resp, nil
}

func AnalyzeTone(creds *ApiKeys, text string) (*WatsonToneResponse, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://gateway.watsonplatform.net/tone-analyzer/api/v3/tone", nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(creds.Watson.ToneAnalysis.Username, creds.Watson.ToneAnalysis.Password)
	q := req.URL.Query()
	q.Add("version", "2016-05-19")
	q.Add("text", text)
	q.Add("sentences", "false")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r = new(WatsonToneResponse)
	err = json.Unmarshal([]byte(bodyBytes), &r)
	if err != nil {
		fmt.Println(err)
	}
	return r, err
}

func FindKeyword(creds *ApiKeys, text string) (*WatsonUnderstandResponse, error) {

	reqBody := new(WatsonUnderstandRequest)
	reqBody.Text = text
	reqBody.Features.Keywords.Emotion = false
	reqBody.Features.Keywords.Sentiment = true
	reqBody.Features.Keywords.Limit = 1

	reqBodyJson, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://gateway.watsonplatform.net/natural-language-understanding/api/v1/analyze", bytes.NewBuffer(reqBodyJson))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(creds.Watson.LanguageUnderstanding.Username, creds.Watson.LanguageUnderstanding.Password)
	q := req.URL.Query()
	q.Add("version", "2017-02-27")
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r = new(WatsonUnderstandResponse)
	err = json.Unmarshal([]byte(bodyBytes), &r)
	if err != nil {
		fmt.Println(err)
	}
	return r, err
}
