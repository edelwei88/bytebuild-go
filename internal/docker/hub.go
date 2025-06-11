package docker

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/edelwei88/bytebuild-go/internal/postgres"
	"github.com/edelwei88/bytebuild-go/internal/postgres/models"
)

func imageExists(imgName string) bool {
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

func imagePull(imgName string) {
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

func imagesPull(imgNames []string, toLog bool) {
	var imagesToInstall []string
	for _, imgName := range imgNames {
		if !imageExists(imgName) {
			imagesToInstall = append(imagesToInstall, imgName)
			if toLog {
				fmt.Println("Image ", imgName, " is going to be installed.")
			}
		} else {
			if toLog {
				fmt.Println("Image ", imgName, " is installed.")
			}
		}
	}

	for _, imgName := range imagesToInstall {
		if toLog {
			fmt.Println("Installing image ", imgName, "...")
		}
		imagePull(imgName)
	}
}

func PrepareImages() {
	var images []string
	var compilers []models.Compiler

	result := postgres.Postgres.Find(&compilers)
	if result.Error != nil {
		log.Fatal(result.Error.Error())
	}

	for _, cmp := range compilers {
		images = append(images, cmp.DockerImageName)
	}

	imagesPull(images, true)
}
