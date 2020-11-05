package apps

/*
	Created by:   Fabián Doñaque
	Copyright by: Fabs Robotics SLU
	Created on:   2020-11-05
*/

///////////////
//  Imports  //
///////////////

import(
	// Core
	"net/http"
	"log"
	"strings"
	"encoding/json"
	"fmt"
	// External
	"github.com/fabiandonaque/go/response"
	"github.com/fabiandonaque/go/methods"
)

///////////////
//  Structs  //
///////////////

type File struct {
	Name string `json:"name"`
	Id string `json:"id"`
	Folder bool `json:"folder"`
	Content []File `json:"content"`
}

//////////////
//  Router  //
//////////////

func Router(w http.ResponseWriter, r *http.Request){
	path := r.URL.Path
	if strings.HasPrefix(path,"/api/apps/getAppsTree"){
		getAppsTree(w,r)
	} else {
		response.New(w,response.ClientError,"Method does not exist.")
	}
}

/////////////////
//  Functions  //
/////////////////

func getAppsTree(w http.ResponseWriter, r *http.Request){
	jsonData,err := methods.BodyToJson(r)
	if err != nil { log.Println("error: "+err.Error()); return }
	log.Println(jsonData)
	var tree []File
	for k := 0; k < 4; k++ {
		newFolder := File{
			Name: fmt.Sprintf("folder%dName",k),
			Id: fmt.Sprintf("folder%d",k),
			Folder: true,
		}
		for j := 0; j < 4; j++ {
			newSubFolder := File{
				Name: fmt.Sprintf("folder%d/subFolder%dName",k,j),
				Id: fmt.Sprintf("folder%d/subFolder%d",k,j),
				Folder: true,
			}
			for i := 0; i < 4; i++ {
				newApp := File{
					Name: fmt.Sprintf("folder%d/subFolder%d/app%dName",k,j,i),
					Id: fmt.Sprintf("folder%d/subFolder%d/app%d",k,j,i),
					Folder: false,
				}
				newSubFolder.Content = append(newSubFolder.Content,newApp)
			}
			newFolder.Content = append(newFolder.Content,newSubFolder)
		}
		tree = append(tree,newFolder)
	}
	data,err := json.Marshal(tree)
	if err != nil { response.New(w,response.ServerError,err.Error()); return }
	dataRaw := json.RawMessage(data)
	response.NewJSON(w,response.Ok,&dataRaw)
}

