package docker

import "fmt"

func InitImages(imgNames []string, toLog bool) {
	var imagesToInstall []string
	for _, imgName := range imgNames {
		if !IsImageInstalled(imgName) {
			imagesToInstall = append(imagesToInstall, imgName)
			if toLog {
				fmt.Println("Image ", imgName, " is installed.")
			}
		} else {
			if toLog {
				fmt.Println("Image ", imgName, " is going to be installed.")
			}
		}
	}

	for _, imgName := range imagesToInstall {
		if toLog {
			fmt.Println("Installing image ", imgName, "...")
		}
		InstallImage(imgName)
	}
}
