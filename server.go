package main

import(
	"net/http"
	"log"
	"github.com/fabiandonaque/go/response"
	"github.com/fabiandonaque/go/methods"
)

func Router(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		http.ServeFile(w,r,"."+r.URL.Path)
	} else if r.Method == http.MethodPost {
		body,err := methods.BodyToJson(r)
		if err != nil { log.Println("error: "+err.Error()); return }
		log.Println(body)
		response.New(w,response.Ok,"")
	}
}

func main(){
	log.Println("Serving at port 13000")
	http.HandleFunc("/", Router)
	log.Fatal(http.ListenAndServe(":13000",nil))
}
