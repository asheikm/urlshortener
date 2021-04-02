// Package has functions that can shorten the given url
package shrink

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"time"
	"urlshortener/model"

	"github.com/sirupsen/logrus"
	"github.com/speps/go-hashids"
)

var UrlLookup map[string]string
var Urls []model.UrlStruct

// Generate Id to match for the given url
func GenerateID() string {
	hd := hashids.NewData()
	h, _ := hashids.NewWithData(hd)
	now := time.Now()
	id, _ := h.Encode([]int{int(now.Unix())})
	return id
}

// This function will add shortened url map to long url, it stores the data into flatfile
// which can be later used to get the data when application restart or on failure
func AddDataToFile(incomingUrls model.UrlStruct) (bool, error) {
	// Data will be overwritten everytime to make sure the json file has the proper data, but restarting app will lose the data
	file, err := os.OpenFile("url_data.out", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		logrus.Info(err)
	}
	defer file.Close()
	var Url model.UrlStruct
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(incomingUrls)
	err = json.Unmarshal(reqBodyBytes.Bytes(), &Url)
	if err != nil {
		logrus.Debug(err)
		return false, err
	}
	Urls = append(Urls, model.UrlStruct{LongUrl: Url.LongUrl, ShortUrl: Url.ShortUrl})
	result, error := json.Marshal(Urls)

	// Write to json file
	encoder := json.NewEncoder(file)
	_ = encoder.Encode(Urls)

	logrus.Info(result, error)
	return true, error
}

// This function will add shortened url map to long url, can be used as lookup to get the
// data based on key value pair, it does not persist data upon restart of the application
func AddDataToMap(incomingUrl model.UrlStruct) (bool, error) {
	if len(UrlLookup) == 0 {
		UrlLookup = make(map[string]string)
	}
	if incomingUrl.LongUrl == "" {
		return false, errors.New("Invalid Long Url")
	}
	UrlLookup[incomingUrl.LongUrl] = incomingUrl.ShortUrl
	return IsURLExists(incomingUrl.LongUrl)
}

// Check if url exists inmemory
func IsURLExists(url string) (bool, error) {
	_, ok := UrlLookup[url]
	if !ok {
		return ok, errors.New("Url not present in the memory")
	}
	return ok, nil
}
