package main

import (
	"./utils"
	"encoding/json"
	"log"
	"net/http"
)


type Data struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}


func TrainRequestHandler(w http.ResponseWriter, request *http.Request){

	var data Data
	cli := utils.GetDockerClient()
	res := cli.ClientVersion()

	if res != "" {
		data = Data{Msg: "收到请求", Code: 200}

	}else{
		data = Data{Msg: "收到请求，但与docker通信失败", Code: 500}
	}
	js, _ := json.Marshal(data)
	w.Write(js)

}

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


