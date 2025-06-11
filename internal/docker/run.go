package docker

import (
	"bytes"
	"context"
	"strings"

	"github.com/MatusOllah/stripansi"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/edelwei88/bytebuild-go/internal/types"
	"github.com/edelwei88/bytebuild-go/internal/utils"
)

func CompileAndExecute(imageName string, fileExt string, cmd string, sourceCode string) (types.ExecResult, error) {
	ctx := context.Background()
	containerName := utils.RandomName(20)
	var result types.ExecResult

	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return result, err
	}
	defer apiClient.Close()

	resp, err := apiClient.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		Cmd:   []string{"/bin/sh"},
		Tty:   true,
	}, nil, nil, nil, containerName)
	if err != nil {
		return result, err
	}
	defer func() {
		apiClient.ContainerRemove(ctx, resp.ID, container.RemoveOptions{Force: true})
	}()

	if err := apiClient.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		return result, err
	}

	shExec := []string{
		"sh",
		"-c",
	}
	filename := strings.Join([]string{"source", fileExt}, "")
	sanitizedSource := strings.Join([]string{"\"", utils.SanitizeForPrintf(sourceCode), "\""}, "")
	createFile := strings.Join([]string{"printf", sanitizedSource, ">", filename}, " ")
	cmdFinal := strings.Join([]string{createFile, cmd}, " && ")

	execConfig := container.ExecOptions{
		Tty:          true,
		WorkingDir:   "/root",
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          append(shExec, cmdFinal),
	}

	execID, err := apiClient.ContainerExecCreate(ctx, resp.ID, execConfig)
	if err != nil {
		return result, err
	}

	respExec, err := apiClient.ContainerExecAttach(ctx, execID.ID, container.ExecAttachOptions{})
	if err != nil {
		return result, err
	}
	defer respExec.Close()

	var stdout, stderr bytes.Buffer
	_, err = stdcopy.StdCopy(&stdout, &stderr, respExec.Reader)
	if err != nil {
		return result, err
	}

	inspect, err := apiClient.ContainerExecInspect(ctx, execID.ID)
	if err != nil {
		return result, err
	}

	result.ExitCode = inspect.ExitCode
	if result.ExitCode == 0 {
		result.Stdout = stripansi.String(stdout.String())
	} else {
		result.Stderr = stripansi.String(stdout.String())
	}

	return result, nil
}

