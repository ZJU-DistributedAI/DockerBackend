package utils

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"log"
)

var cli *client.Client

func InitDockerClient() {

	if cli != nil {
		return
	}

	ctx := context.Background()
	var err error
	cli, err = client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		log.Panic(err)
	}
	cli.NegotiateAPIVersion(ctx)

}





func GetDockerClient() *client.Client {
	if cli == nil {
		InitDockerClient()
	}
	return cli
}

func CreateDockerContainer(cli *client.Client, imagename string, ports ...string) container.ContainerCreateCreatedBody {

	//创建容器
	portMap := make(nat.PortMap, 0)

	exports := make(nat.PortSet, len(ports)+1)

	for i := range ports {

		newPort, err := nat.NewPort("tcp", ports[i])
		if err != nil {
			log.Panic(err)
		}
		exports[newPort] = struct{}{}

		portBindings := make([]nat.PortBinding, 0, 1)
		portBindings = append(portBindings, nat.PortBinding{HostPort: ports[i]})
		portMap[newPort] = portBindings
	}

	resp, err := cli.ContainerCreate(context.Background(), &container.Config{
		Image: imagename,
		Cmd:   []string{""},
		Volumes: map[string]struct{}{
			"C:\\Users\\huyifan01\\Documents\\MachineLearning": struct{}{},
		},
		ExposedPorts: exports,
	}, &container.HostConfig{
		PortBindings: portMap,
	}, nil, "test")

	if err != nil {
		log.Panic(err)
	}

	return resp
}

func StartDockerContainer(cli *client.Client, resp container.ContainerCreateCreatedBody) error {

	err := cli.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{})

	return err
}

func GetDockerImages() {

}

func GetDockerContainers() {

}
