package setup

import (
	"log"

	"github.com/edelwei88/bytebuild-go/internal/docker"
	"github.com/edelwei88/bytebuild-go/internal/postgres"
	"github.com/edelwei88/bytebuild-go/internal/postgres/models"
)

func DownloadImages() {
	var images []string
	var compilers []models.Compiler

	result := postgres.Postgres.Find(&compilers)
	if result.Error != nil {
		log.Fatal("failed to pull docker images")
	}

	for _, cmp := range compilers {
		images = append(images, cmp.DockerImageName)
	}

	docker.InitImages(images, true)
}
