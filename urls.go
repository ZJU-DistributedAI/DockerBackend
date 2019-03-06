package main

import (
	"log"
	"net/http"
	"./utils"
)

func init(){
	utils.InitDockerClient()
}

func startDockerBackend(){


	http.HandleFunc("/dockerbackend/dockerstatus", GetDockerStatusHandler)
	http.HandleFunc("/dockerbackend/starttrain", StartTrainHandler)

	// 启动web服务，监听9090端口
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}


func main(){



	startDockerBackend()


}




