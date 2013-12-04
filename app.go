package gogist

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const t = `
<html>
  <head>
    <meta name="go-import" content="%s git https://gist.github.com/%s.git" />
    <script>window.location='https://github.com/ImJasonH/go-gist/';</script>
  </head>
</html>
`

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/{username}/{gistID:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(fmt.Sprintf(t, r.URL.Host+r.URL.Path, mux.Vars(r)["gistID"])))
	}).Methods("GET")
	r.Handle("/", http.RedirectHandler("https://github.com/ImJasonH/go-gist", http.StatusSeeOther))
	http.Handle("/", r)
}
