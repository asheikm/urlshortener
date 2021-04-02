// package redirect
package redirect

import (
	"net/http"
	"urlshortener/api"
	"urlshortener/shrink"

	"github.com/sirupsen/logrus"
)

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	var url api.InputUrl
	// Check if given url found on the memory map
	err, v := api.JsonDecodeWrapper(r, url)
	if err != nil {
		w.WriteHeader(500)
	}
	rurl, ok := shrink.GetDataFromMap(v.Url)
	if ok {
		logrus.Info("Redirecting url " + v.Url + " to " + rurl)
		http.Redirect(w, r, rurl, http.StatusSeeOther)
	} else {
		w.WriteHeader(500)
	}
}
