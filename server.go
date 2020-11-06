package main

import(
	"net/http"
	"log"
	"strings"
	"io/ioutil"
	"github.com/fabiandonaque/go/response"
	"github.com/fabiandonaque/fabs-webapp/api/apps"
)

func Router(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		http.ServeFile(w,r,"./static/"+r.URL.Path)
	} else if r.Method == http.MethodPost {
		path := r.URL.Path
		if strings.HasPrefix(path, "/api/apps"){
			apps.Router(w,r)
		} else if strings.HasPrefix(path, "/api/size") {
			raw, err := ioutil.ReadAll(r.Body)
			if err != nil { log.Println(err); return }
			log.Println(string(raw))
			response.New(w,response.Ok,"")
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
