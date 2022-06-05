package html_utils

import (
	"github.com/jakubruminski/golang/golang_webapp_utils/error_utils"
	"net/http"
	"text/template"
)

func ServeHtml(w http.ResponseWriter, pathToHtml string) {
	t, err := template.ParseFiles(pathToHtml)

	error_utils.CheckErrorButDoNotStop(err, nil)

	t.Execute(w, nil)
}
