package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type cliExecOutput struct{
	Result int `json:"result"`
	Str string `json:"string"`
	UserStatus int `json:"userStatus"`
	AccountMode int `json:"accountMode"`
	CallMode int `json:"callMode"`
}

type crPostData struct{
	CliExecResult int `json:"cliExecResult"`
	Output cliExecOutput `json:"cliExecOutput"`
}

type crPost struct{
	Status int `json:"status"`
	Data crPostData `json:"data"`

}

func HandleCrWebReuqest(w http.ResponseWriter, r *http.Request){
	post := crPost{
		Status: 0,
		Data: crPostData{
			CliExecResult: 0,
			Output: cliExecOutput{
				Result: 0,
				Str: "test",
			},
		},
	}
	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil{
		return
	}
	fmt.Fprintln(w, string(output))
}

func main(){
	server := http.Server{
		Addr: "172.27.0.3:8776",
	}

	http.HandleFunc("/cr_web/smm/ivr", HandleCrWebReuqest)

	server.ListenAndServe()
}
