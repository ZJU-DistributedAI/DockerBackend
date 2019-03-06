package utils

import (
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"log"
)

var cli *client.Client

func InitDockerClient(){

	if cli != nil{
		return
	}

	ctx := context.Background()
	var err error
	cli, err = client.NewClientWithOpts(client.FromEnv)

	if err != nil{
		log.Panic(err)
	}
	cli.NegotiateAPIVersion(ctx)


}

func GetDockerClient()(*client.Client){
	if cli == nil {
		InitDockerClient()
	}
	return cli
}


func CreateDockerContainer(cli *client.Client)(container.ContainerCreateCreatedBody){

	//创建容器
	portMap := make(nat.PortMap,0)

	jupyterPort, err := nat.NewPort("tcp", "8888")
	tensorBoardPort, err := nat.NewPort("tcp", "6006")
	backendPort, err := nat.NewPort("tcp", "9091")

	exports := make(nat.PortSet, 10)
	exports[jupyterPort] = struct{}{}
	exports[tensorBoardPort] = struct{}{}
	exports[backendPort] = struct{}{}


	jupyterPortBindings := make([]nat.PortBinding, 0, 1)
	tensorBoardPortBindings := make([]nat.PortBinding, 0, 1)
	backendPortBindings := make([]nat.PortBinding, 0, 1)


	jupyterPortBindings = append(jupyterPortBindings, nat.PortBinding{HostPort: "8888"})
	tensorBoardPortBindings = append(tensorBoardPortBindings, nat.PortBinding{HostPort: "6006"})
	backendPortBindings = append(backendPortBindings, nat.PortBinding{HostPort: "9091"})

	portMap[jupyterPort] = jupyterPortBindings
	portMap[tensorBoardPort] = tensorBoardPortBindings
	portMap[backendPort] = backendPortBindings

	resp, err := cli.ContainerCreate(context.Background(),&container.Config{
		Image: "dash00/tensorflow-python3-jupyter",
		Cmd: []string{""},
		Volumes: map[string]struct{}{
			"C:\\Users\\huyifan01\\Documents\\MachineLearning":struct{}{},
		},
		ExposedPorts:exports,
	}, &container.HostConfig{
		PortBindings: portMap,
	}, nil,"test")

	if err != nil {
		log.Panic(err)
	}

	return resp
}


func StartDockerContainer(cli *client.Client){



}

func GetDockerImages(){



}



func GetDockerContainers(){


}