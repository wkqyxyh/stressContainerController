package main 

import (
	"fmt"
	"time"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	
	"stressContainerController/src/util"
)

var ContainerIDList []string

type configuration struct {
    RampUpPeriod int
    RampUpFinalContainerNum int
    DurationPeriod int
    RampDownPeriod int
    ImageName string
}

func stopAllContainers(rampdownSleepTime int) {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	for index, containerID := range ContainerIDList {
		fmt.Println(index, containerID)
		time.Sleep(time.Duration(rampdownSleepTime) * time.Millisecond)
		if err := cli.ContainerStop(ctx, containerID, nil); err != nil {
			panic(err)
		}
	}
	fmt.Println("Success")
}

func startContainer(imgName string) {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	imageName := imgName

	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, out)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println(resp.ID)
	
	ContainerIDList= append(ContainerIDList, resp.ID)
}

func scheduler() {
	conf := util.ReadConfigFile("cfg/config.json")
	
	//rampup schedule
	rampupSleepTime := conf.RampUpPeriod*1000/conf.RampUpFinalContainerNum	
	for i:=0; i<conf.RampUpFinalContainerNum; i++ {
		startContainer(conf.ImageName)
		time.Sleep(time.Duration(rampupSleepTime) * time.Millisecond)
	}
	
	//duration
	time.Sleep(time.Duration(conf.DurationPeriod) * time.Second)
	
	//rampdown schedule	
	rampdownSleepTime := conf.RampDownPeriod*1000/conf.RampUpFinalContainerNum
	stopAllContainers(rampdownSleepTime)
}
func main() {
	scheduler()
}