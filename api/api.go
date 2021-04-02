// Middleware/Backend for rest api
package api

import (
	"encoding/json"
	"net/http"
	"net/http/httputil"
	"os"
	"urlshortener/model"
	"urlshortener/shrink"
	"urlshortener/utils"

	"github.com/sirupsen/logrus"
)

// This init function is by default called only once
func init() {
	if os.Getenv("SHORT_DOMAIN") == "" {
		os.Setenv("SHORT_DOMAIN", "http://localhost:"+os.Getenv("LISTEN_PORT")+"/")
	}
	shrink.UrlLookup = make(map[string]string)
}

type InputUrl struct {
	Url string `json:"url"`
}

// log input request data
func logRequestData(r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		logrus.Error(err, http.StatusInternalServerError)
		return
	}
	logrus.Info("Request dump: ", string(dump))
}

// wrapper function for json decoder
func jsonDecodeWrapper(r *http.Request, iu InputUrl) (error, InputUrl) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&iu)
	if err != nil {
		logrus.Error(err)
	}
	return err, iu
}

// Version Handler
func GetVersion(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("v1.0"))
}

// Handler function implemenation for get call with request url
func GetShortenedURL(w http.ResponseWriter, r *http.Request) {
	var incomingUrl InputUrl
	var urls model.UrlStruct
	logRequestData(r)
	err, incomingUrl := jsonDecodeWrapper(r, incomingUrl)
	if err != nil {
		json.NewEncoder(w).Encode("Unable to decode json input")
	}
	logrus.Info("Get shortened url for : " + incomingUrl.Url)
	strippedUrl := utils.StripURL(incomingUrl.Url)
	if utils.IsValidURL(strippedUrl) != true {
		w.WriteHeader(400)
		w.Write([]byte("Invalid input url"))
	}
	urls.LongUrl = incomingUrl.Url
	urls.ShortUrl = shrink.UrlLookup[strippedUrl]
	json.NewEncoder(w).Encode(urls)
}

// Handler function implementation for post call with request url
func CreateShortenedURL(w http.ResponseWriter, r *http.Request) {
	var urls model.UrlStruct
	var incomingUrl InputUrl
	logRequestData(r)
	err, incomingUrl := jsonDecodeWrapper(r, incomingUrl)
	if utils.IsValidURL(incomingUrl.Url) != true {
		w.WriteHeader(400)
		w.Write([]byte("Invalid input url"))
	}
	logrus.Info(incomingUrl.Url)
	_ = json.NewDecoder(r.Body).Decode(&incomingUrl.Url)
	logrus.Info("Long Url: ", incomingUrl.Url)
	// Remove http or www from request body
	urls.LongUrl = utils.StripURL(incomingUrl.Url)
	// Check if data already present in memory
	ok, err := shrink.IsURLExists(urls.LongUrl)
	if ok {
		json.NewEncoder(w).Encode("Url already exists in memory or file")
	} else {
		id := shrink.GenerateID()
		urls.ShortUrl = os.Getenv("SHORT_DOMAIN") + id
		json.NewEncoder(w).Encode(urls)
		_, err = shrink.AddDataToMap(urls)
		if err != nil {
			json.NewEncoder(w).Encode("Unable to add url into memory")
		}
		_, err = shrink.AddDataToFile(urls)
		if err != nil {
			json.NewEncoder(w).Encode("Unable to add url into memory")
		}
	}
}
