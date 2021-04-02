package shrink_test

import (
	"testing"
	"urlshortener/model"
	"urlshortener/shrink"
)

func TestAddDataToMapFailureCase(t *testing.T) {
	var m model.UrlStruct
	m.LongUrl, m.ShortUrl = "", ""
	result, _ := shrink.AddDataToMap(m)
	if result != false {
		t.Errorf("Unable to add to map: case failed")
	}
}

func TestAddDataToMapPassCase(t *testing.T) {
	var m model.UrlStruct
	m.LongUrl, m.ShortUrl = "google.com", shrink.GenerateID()
	result, _ := shrink.AddDataToMap(m)
	if result != true {
		t.Errorf("Unable to add to map: case failed")
	}
}

