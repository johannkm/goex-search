package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	textapi "github.com/AYLIEN/aylien_textapi_go"
)

const (
	ParallelDotsUrl = "https://apis.paralleldots.com/emotion"
)

var TextApiClient *textapi.Client
var Lexicons map[string]*[10]bool

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

func Summarize(reviews string, creds *ApiKeys) (resp *SummaryResponse, err error) {

	if TextApiClient == nil {
		auth := textapi.Auth{creds.TextApi.AppId, creds.TextApi.Key}
		TextApiClient, err = textapi.NewClient(auth, true)
		if err != nil {
			panic(err)
		}
		// ParseLexicon()
	}

	fmt.Println("incoming review text: " + reviews)
	summarizeParams := &textapi.SummarizeParams{
		Title:             "Review",
		Text:              reviews,
		NumberOfSentences: 1,
	}
	summary, err := TextApiClient.Summarize(summarizeParams)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", TextApiClient.RateLimits)

	resp = new(SummaryResponse)
	summaryText := ""
	for x := range summary.Sentences {
		summaryText = summaryText + summary.Sentences[x]
	}

	understanding, err := FindKeyword(creds, reviews)
	if err != nil {
		panic(err)
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

func ProcessText(text string, creds *ApiKeys) (string, error) {

	resp, err := FindKeyword(creds, text)
	if err != nil {
		panic(err)
	}

	return resp.Keywords[0].Text, nil

	// resp, err := AnalyzeTone(creds, text)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Print("incoming tone: ")
	// fmt.Println(resp)
	//
	// res, err := FindAdjective(resp, Lexicons)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("chosen word: " + res)
	// return res, nil

}

func FindAdjective(resp *WatsonToneResponse, lexicons map[string]*[10]bool) (string, error) {

	toneIndex := make(map[string]int32)
	toneIndex["anger"] = 0
	toneIndex["disgust"] = 2
	toneIndex["fear"] = 3
	toneIndex["joy"] = 4
	toneIndex["sadness"] = 7

	toneCats := resp.DocumentTone.ToneCategories
	for i := 0; i < len(toneCats); i++ {
		if toneCats[i].CategoryId == "emotion_tone" {
			tones := toneCats[i].Tones
			emotions := make([]int32, 0, 5)
			for j := 0; j < len(tones); j++ {
				if tones[j].Score > 0.5 {
					emotions = append(emotions, toneIndex[tones[j].ToneId])
				}
			}
			for k, v := range lexicons {
				if AdjectiveMatches(emotions, v) {
					return k, nil
				}
			}
			for i := range emotions {
				fmt.Println(emotions[i])
			}

			return "", nil
		}
	}

	return "", nil
}

func AdjectiveMatches(revEmotions []int32, lexiEmotios *[10]bool) bool {
	for i := range revEmotions {
		if lexiEmotios[revEmotions[i]] == false {
			return false
		}
	}
	return true
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
		panic(err)
	}
	return r, err
}

func ParseLexicon() error {
	file, err := os.Open("NRC-Emotion-Lexicon-Wordlevel-v0.92.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	var wordEmos = make(map[string]*[10]bool)

	scanner := bufio.NewScanner(file)
	counter := 0
	for scanner.Scan() {
		cols := strings.Fields(scanner.Text())
		if len(cols) < 3 {
			continue
		}
		word := cols[0]
		// emotion := cols[1]
		emotionPresent := cols[2]

		if _, hasKey := wordEmos[word]; !hasKey {
			wordEmos[word] = new([10]bool)
			counter = 0
		}
		var c bool
		c = emotionPresent == "1"
		wordEmos[word][counter] = c
		counter++
	}
	if err = scanner.Err(); err != nil {
		return err
	}

	Lexicons = wordEmos
	return nil
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

func FindKeyword(creds *ApiKeys, text string) (*WatsonUnderstandResponse, error) {

	reqBody := new(WatsonUnderstandRequest)
	reqBody.Text = text
	reqBody.Features.Keywords.Emotion = false
	reqBody.Features.Keywords.Sentiment = true
	reqBody.Features.Keywords.Limit = 1

	reqBodyJson, err := json.Marshal(reqBody)
	if err != nil {
		panic(err)
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
		panic(err)
	}
	return r, err
}

// func ParseLexicon() error {
// 	file, err := os.Open("NRC-Emotion-Lexicon-Wordlevel-v0.92.txt")
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()
//
// 	var wordEmos = make(map[string][]string)
//
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		cols := strings.Fields(scanner.Text())
// 		if len(cols) < 3 {
// 			continue
// 		}
// 		word := cols[0]
// 		emotion := cols[1]
// 		emotionPresent := cols[2]
//
// 		if emotionPresent == "1" {
// 			if _, hasKey := wordEmos[word]; !hasKey {
// 				wordEmos[word] = make([]string, 5, 10)
// 			}
// 			wordEmos[word] = append(wordEmos[word], emotion)
// 		}
// 	}
// 	if err = scanner.Err(); err != nil {
// 		return err
// 	}
//
// 	Lexicons = wordEmos
// 	return nil
// }

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
