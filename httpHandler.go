package main

import (
	"./utils"
	"log"
	"net/http"
)

func StartTrainHandler(w http.ResponseWriter, request *http.Request){


	cli := utils.GetDockerClient()


	var ports []string = []string{"8888", "6006", "9091"}
	resp := utils.CreateDockerContainer(cli,"dash00/tensorflow-python3-jupyter", ports...)

	err := utils.StartDockerContainer(cli, resp)

	if err != nil {
		log.Panic(err)
		w.Write([]byte("创建容器失败"))
	}
	w.Write([]byte("创建容器成功"))

}


func GetDockerStatusHandler(w http.ResponseWriter, request *http.Request){

	cli := utils.GetDockerClient()

	res := cli.ClientVersion()

	w.Write([]byte(res))

}