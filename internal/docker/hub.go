package docker

import (
	"context"
	"io"
	"log"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

func IsImageInstalled(imgName string) bool {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}
	defer apiClient.Close()

	if _, err := apiClient.ImageInspect(
		context.Background(),
		imgName,
		client.ImageInspectWithAPIOpts(image.InspectOptions{})); err != nil {
		return false
	}

	return true
}

func InstallImage(imgName string) {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}
	defer apiClient.Close()

	res, err := apiClient.ImagePull(context.Background(), imgName, image.PullOptions{})
	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()

	io.ReadAll(res)
}
