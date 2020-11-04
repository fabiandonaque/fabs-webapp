package apps

import(
	"net/http"
	"log"
	"strings"
	"github.com/fabiandonaque/go/response"
	"github.com/fabiandonaque/go/methods"
)

func Router(w http.ResponseWriter, r *http.Request){
	path = r.URL.Path
	if strings.HasPrefix(path,"/api/apps/getAppsTree"){
		getAppsTree(w,r)
	} else {
		response.New(w,response.ClientError,"Method does not exist.")
	}
}

func getAppsTree(w http.ResponseWriter, r *http.Request){
	jsonData,err := methods.BodyToJson(r)
	if err != nil { log.Println("error: "+err.Error()); return }
	log.Println(body)
	response.New(w,response.Ok,"")
}

