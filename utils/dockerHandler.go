package utils

import (
	"context"
	"github.com/docker/docker/client"
	"log"
)

var cli *client.Client

func GetDockerClient()(*client.Client){

	if cli != nil{
		return cli
	}

	ctx := context.Background()
	var err error
	cli, err = client.NewClientWithOpts(client.FromEnv)

	if err != nil{
		log.Panic(err)
	}
	cli.NegotiateAPIVersion(ctx)

	return cli

}

func GetDockerImages(){

	

}



func GetDockerContainers(){


}