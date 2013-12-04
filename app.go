package gogist

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const t = `
<html>
  <head>
    <meta name="go-import" content="%s git gist.github.com/%s/%s.git" />
    <script>window.location='/';</script>
  </head>
</html>
`

// TODO: Move to a static file.
const explain = `
<html>
  <head>
    <title>Go import redirect for GitHub Gists</title>
    <style>
* {
  font-family: Helvetica, Arial, sans-serif;
  font-size: 16px;
}

h2 {
  font-size: 20px;
  font-weight: bold;
  color: #375EAB;
}

code {
  font-family: Menlo, monospace;
  font-size: 14px;
  background: #e9e9e9;
}
    </style>
  </head>
  <body>
    <h2>Welcome!</h2>

    <p>Ever see some tasty Go code in a GitHub Gist and want to try it out? What are you supposed to do, clone it into your own GitHub repo and <code>go get</code> it? Copy it locally? What are you a farmer?</p>

    <p>Why can't you just <code>import "gist.github.com/UserName/112358"</code>? Because Gists don't support that, unfortunately. :(</p>

    <p>Which is why I made this silly hack. Instead of <code>import "gist.github.com/..."</code>, <code>import "go-gist.appspot.com/..."</code>, and this site will tell the Go tool to go fetch the contents of the Gist instead.</p>

    <p>It's all described <a href="http://golang.org/cmd/go/#hdr-Remote_import_paths">in the docs</a>. Those guys thought of everything.</p>

    <p>By the way, this would be totally unnecessary if the Gist HTML page was served with the necessary &lt;meta&gt; tag. If you work at GitHub or know someone who does, bug them to add a tag to their Gist HTML so I can turn this site off.</p>
  </body>
</html>
`

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/{username}/{gistID:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		v := mux.Vars(r)
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(fmt.Sprintf(t, r.URL.Host+r.URL.Path, v["username"], v["gistID"])))
	}).Methods("GET")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(explain))
	})
	http.Handle("/", r)
}
