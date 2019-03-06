package main

import (
	"./utils"
	"net/http"
)

func StartTrainHandler(w http.ResponseWriter, request *http.Request){






}


func GetDockerStatusHandler(w http.ResponseWriter, request *http.Request){

	cli := utils.GetDockerClient()

	res := cli.ClientVersion()

	w.Write([]byte(res))

}