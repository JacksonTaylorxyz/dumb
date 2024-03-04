package handlers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rramiachraf/dumb/views"
	"github.com/sirupsen/logrus"
)

func isValidExt(ext string) bool {
	valid := []string{"jpg", "jpeg", "png", "gif"}
	for _, c := range valid {
		if strings.ToLower(ext) == c {
			return true
		}
	}

	return false
}

func ImageProxy(l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v := mux.Vars(r)
		f := v["filename"]
		ext := v["ext"]

		if !isValidExt(ext) {
			w.WriteHeader(http.StatusBadRequest)
			views.ErrorPage(400, "something went wrong").Render(context.Background(), w)
			return
		}

		// first segment of URL resize the image to reduce bandwith usage.
		url := fmt.Sprintf("https://t2.genius.com/unsafe/300x300/https://images.genius.com/%s.%s", f, ext)

		res, err := sendRequest(url)
		if err != nil {
			l.Errorln(err)
			w.WriteHeader(http.StatusInternalServerError)
			views.ErrorPage(500, "cannot reach Genius servers").Render(context.Background(), w)
			return
		}

		if res.StatusCode != http.StatusOK {
			w.WriteHeader(http.StatusInternalServerError)
			views.ErrorPage(500, "something went wrong").Render(context.Background(), w)
			return
		}

		w.Header().Add("Content-type", fmt.Sprintf("image/%s", ext))
		if _, err = io.Copy(w, res.Body); err != nil {
			l.Errorln("unable to write image", err)
		}
	}
}
