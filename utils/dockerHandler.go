package utils

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
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

func CreateDockerContainer(cli *client.Client, directorypath string, imagename string, ports ...string) container.ContainerCreateCreatedBody {

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
		ExposedPorts: exports,
		Cmd: []string{"./train.sh"},
	}, &container.HostConfig{
		PortBindings: portMap,
		Mounts: []mount.Mount{
			{
				Type: mount.TypeBind,
				Source: directorypath, //"/root/MachineLearning"
				Target: "/MachineLearning",
			},
		},
	}, nil, "")

	if err != nil {
		log.Println("创建容器出错: ",err)
	}

	return resp
}

func StartDockerContainer(cli *client.Client, resp container.ContainerCreateCreatedBody) error {

	err := cli.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{})

	return err
}

func WaitForContainer(cli *client.Client, containerid string)(){

	resp, errchan := cli.ContainerWait(context.Background(), containerid, container.WaitConditionNotRunning)

	select{
	case err:= <- errchan:
		if err != nil {
			log.Panic("wait for container fail: ", err)
		}

	case <-resp:
	}

}

func GetDockerImages() {

}

func GetDockerContainers() {

}
