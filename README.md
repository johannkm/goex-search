# goex-search

A Yelp search app that summarizes reviews using Watson and Aylien Text API

## Install

Install [golang](https://golang.org/doc/install) and set your [gopath](https://github.com/golang/go/wiki/GOPATH).

Get api keys for:
- [yelp fusion](https://www.yelp.com/developers)
- [google places (web service)](https://developers.google.com/places/web-service/)
- [watson language understanding](https://www.ibm.com/watson/developercloud/natural-language-understanding.html)
- [aylien text analysis](http://aylien.com/text-api)

Set your api credentials in `api-config.json`

``` bash
cd $GOPATH/src/
git clone https://github.com/johannkm/goex-search.git

cd goex-search/

npm install
npm run build

go build
go install

PORT=8000 $GOPATH/bin/goex-search
# running at http://localhost:8000
```
