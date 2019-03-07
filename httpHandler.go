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

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

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

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

	var data Data
	cli := utils.GetDockerClient()


	var ports []string = []string{"8888", "6006", "9091"}
	resp := utils.CreateDockerContainer(cli, "//root//MachineLearning",
		"zjudistributeai/images:v0.3", ports...)

	err := utils.StartDockerContainer(cli, resp)

	if err != nil {
		log.Panic("启动容器失败: ", err)
	}

	utils.WaitForContainer(cli, resp.ID)


	data = Data{Msg: "执行训练任务成功", Code: 200}


	js, _ := json.Marshal(data)
	w.Write(js)

}


func GetDockerStatusHandler(w http.ResponseWriter, request *http.Request){

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST,GET")

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


