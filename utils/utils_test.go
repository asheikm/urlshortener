package utils_test

import "testing"
import "urlshortener/utils"

func TestIsValidURLwithnoInput(t *testing.T) {
	result := utils.IsValidURL("")
	if result != false {
		t.Errorf("URL is not validated properly")
	}
}

func TestWithValidUrl(t *testing.T) {
        result := utils.IsValidURL("www.google.com/test")
        if result != true {
                t.Errorf("URL is not validated properly")
        }
}

func TestWithNotValidUrl(t *testing.T) {
        result := utils.IsValidURL("www.googlecom/test")
        if result != true {
                t.Errorf("URL is not validated properly")
        }
}

