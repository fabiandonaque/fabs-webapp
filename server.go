package main

import(
	"net/http"
	"log"
	"strings"
	"github.com/fabiandonaque/go/response"
	"github.com/fabiandonaque/go/methods"
	"github.com/fabiandonaque/fabs-webapp/static/apps"
)

func Router(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		http.ServeFile(w,r,"./static/"+r.URL.Path)
	} else if r.Method == http.MethodPost {
		path = r.URL.Path
		if strings.HasPrefix(path, "/api/apps"){
			apps.Router(w,r)
		} else {
			response.New(w,response.ClientError,"Method does not exist.")
		}
	}
}

func main(){
	log.Println("Serving at port 13000")
	http.HandleFunc("/", Router)
	log.Fatal(http.ListenAndServe(":13000",nil))
}
