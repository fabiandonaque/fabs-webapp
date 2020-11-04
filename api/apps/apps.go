package apps

import(
	"net/http"
	"log"
	"strings"
	"encoding/json"
	"github.com/fabiandonaque/go/response"
	"github.com/fabiandonaque/go/methods"
)

func Router(w http.ResponseWriter, r *http.Request){
	path := r.URL.Path
	if strings.HasPrefix(path,"/api/apps/getAppsTree"){
		getAppsTree(w,r)
	} else {
		response.New(w,response.ClientError,"Method does not exist.")
	}
}

func getAppsTree(w http.ResponseWriter, r *http.Request){
	jsonData,err := methods.BodyToJson(r)
	if err != nil { log.Println("error: "+err.Error()); return }
	log.Println(jsonData)
	var tree = map[string]interface{}{
		"app1": map[string]interface{}{
			"sub1_app1": map[string]interface{}{
				"sub1_sub1_app1": map[string]interface{}{},
				"sub2_sub1_app1": map[string]interface{}{},
				"sub3_sub1_app1": map[string]interface{}{},
				"sub4_sub1_app1": map[string]interface{}{},
			},
			"sub2_app1": map[string]interface{}{
				"sub1_sub2_app1": map[string]interface{}{},
				"sub2_sub2_app1": map[string]interface{}{},
				"sub3_sub2_app1": map[string]interface{}{},
				"sub4_sub2_app1": map[string]interface{}{},
			},
			"sub3_app1": map[string]interface{}{
				"sub1_sub3_app1": map[string]interface{}{},
				"sub2_sub3_app1": map[string]interface{}{},
				"sub3_sub3_app1": map[string]interface{}{},
				"sub4_sub3_app1": map[string]interface{}{},
			},
			"sub4_app1": map[string]interface{}{
				"sub1_sub4_app1": map[string]interface{}{},
				"sub2_sub4_app1": map[string]interface{}{},
				"sub3_sub4_app1": map[string]interface{}{},
				"sub4_sub4_app1": map[string]interface{}{},
			},
		},
		"app2": map[string]interface{}{
			"sub1_app2": map[string]interface{}{
				"sub1_sub1_app2": map[string]interface{}{},
				"sub2_sub1_app2": map[string]interface{}{},
				"sub3_sub1_app2": map[string]interface{}{},
				"sub4_sub1_app2": map[string]interface{}{},
			},
			"sub2_app2": map[string]interface{}{
				"sub1_sub2_app2": map[string]interface{}{},
				"sub2_sub2_app2": map[string]interface{}{},
				"sub3_sub2_app2": map[string]interface{}{},
				"sub4_sub2_app2": map[string]interface{}{},
			},
			"sub3_app2": map[string]interface{}{
				"sub1_sub3_app2": map[string]interface{}{},
				"sub2_sub3_app2": map[string]interface{}{},
				"sub3_sub3_app2": map[string]interface{}{},
				"sub4_sub3_app2": map[string]interface{}{},
			},
			"sub4_app2": map[string]interface{}{
				"sub1_sub4_app2": map[string]interface{}{},
				"sub2_sub4_app2": map[string]interface{}{},
				"sub3_sub4_app2": map[string]interface{}{},
				"sub4_sub4_app2": map[string]interface{}{},
			},
		},
		"app3": map[string]interface{}{
			"sub1_app3": map[string]interface{}{
				"sub1_sub1_app3": map[string]interface{}{},
				"sub2_sub1_app3": map[string]interface{}{},
				"sub3_sub1_app3": map[string]interface{}{},
				"sub4_sub1_app3": map[string]interface{}{},
			},
			"sub2_app3": map[string]interface{}{
				"sub1_sub2_app3": map[string]interface{}{},
				"sub2_sub2_app3": map[string]interface{}{},
				"sub3_sub2_app3": map[string]interface{}{},
				"sub4_sub2_app3": map[string]interface{}{},
			},
			"sub3_app3": map[string]interface{}{
				"sub1_sub3_app3": map[string]interface{}{},
				"sub2_sub3_app3": map[string]interface{}{},
				"sub3_sub3_app3": map[string]interface{}{},
				"sub4_sub3_app3": map[string]interface{}{},
			},
			"sub4_app3": map[string]interface{}{
				"sub1_sub4_app3": map[string]interface{}{},
				"sub2_sub4_app3": map[string]interface{}{},
				"sub3_sub4_app3": map[string]interface{}{},
				"sub4_sub4_app3": map[string]interface{}{},
			},
		},
		"app4": map[string]interface{}{
			"sub1_app4": map[string]interface{}{
				"sub1_sub1_app4": map[string]interface{}{},
				"sub2_sub1_app4": map[string]interface{}{},
				"sub3_sub1_app4": map[string]interface{}{},
				"sub4_sub1_app4": map[string]interface{}{},
			},
			"sub2_app4": map[string]interface{}{
				"sub1_sub2_app4": map[string]interface{}{},
				"sub2_sub2_app4": map[string]interface{}{},
				"sub3_sub2_app4": map[string]interface{}{},
				"sub4_sub2_app4": map[string]interface{}{},
			},
			"sub3_app4": map[string]interface{}{
				"sub1_sub3_app4": map[string]interface{}{},
				"sub2_sub3_app4": map[string]interface{}{},
				"sub3_sub3_app4": map[string]interface{}{},
				"sub4_sub3_app4": map[string]interface{}{},
			},
			"sub4_app4": map[string]interface{}{
				"sub1_sub4_app4": map[string]interface{}{},
				"sub2_sub4_app4": map[string]interface{}{},
				"sub3_sub4_app4": map[string]interface{}{},
				"sub4_sub4_app4": map[string]interface{}{},
			},
		},
	}
	data,err := json.Marshal(tree)
	if err != nil { response.New(w,response.ServerError,err.Error()); return }
	dataRaw := json.RawMessage(data)
	response.NewJSON(w,response.Ok,&dataRaw)
}

