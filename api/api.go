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

func init() {
	if os.Getenv("SHORT_DOMAIN") == "" {
		os.Setenv("SHORT_DOMAIN", "http://example.me/")
	}
	shrink.UrlLookup = make(map[string]string)
}

// logRequestData logs the input request data
func logRequestData(r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		logrus.Error("Failed to dump request: ", err)
		return
	}
	logrus.Info("Request dump: ", string(dump))
}

// GetVersion returns the API version
func GetVersion(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("v1.0"))
}

// GetShortenedURL handles GET requests to retrieve a shortened URL
func GetShortenedURL(w http.ResponseWriter, r *http.Request) {
	logRequestData(r)

	var incomingUrl model.InputUrl
	if err := json.NewDecoder(r.Body).Decode(&incomingUrl); err != nil {
		http.Error(w, "Unable to decode JSON input", http.StatusBadRequest)
		return
	}

	logrus.Info("Get shortened url for: ", incomingUrl.Url)
	strippedUrl := utils.StripURL(incomingUrl.Url)
	if !utils.IsValidURL(strippedUrl) {
		http.Error(w, "Invalid input URL", http.StatusBadRequest)
		return
	}

	shortUrl, exists := shrink.UrlLookup[strippedUrl]
	if !exists {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	response := model.UrlStruct{
		LongUrl:  incomingUrl.Url,
		ShortUrl: shortUrl,
	}
	json.NewEncoder(w).Encode(response)
}

// CreateShortenedURL handles POST requests to create a shortened URL
func CreateShortenedURL(w http.ResponseWriter, r *http.Request) {
	logRequestData(r)

	var incomingUrl model.InputUrl
	if err := json.NewDecoder(r.Body).Decode(&incomingUrl); err != nil {
		http.Error(w, "Unable to decode JSON input", http.StatusBadRequest)
		return
	}

	logrus.Info("Received URL: ", incomingUrl.Url)
	strippedUrl := utils.StripURL(incomingUrl.Url)
	if !utils.IsValidURL(strippedUrl) {
		http.Error(w, "Invalid input URL", http.StatusBadRequest)
		return
	}

	exists, err := shrink.IsURLExists(strippedUrl)
	if err != nil {
		http.Error(w, "Error checking URL existence", http.StatusInternalServerError)
		return
	}

	if exists {
		http.Error(w, "URL already exists", http.StatusConflict)
		return
	}

	id := shrink.GenerateID()
	shortUrl := os.Getenv("SHORT_DOMAIN") + id
	response := model.UrlStruct{
		LongUrl:  strippedUrl,
		ShortUrl: shortUrl,
	}

	if _, err := shrink.AddDataToMap(response); err != nil {
		http.Error(w, "Unable to add URL to memory", http.StatusInternalServerError)
		return
	}

	if _, err := shrink.AddDataToFile(response); err != nil {
		http.Error(w, "Unable to add URL to file", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response)
}
