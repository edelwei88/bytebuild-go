package docker

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/edelwei88/bytebuild-go/internal/types"
)

func sanitizeString(str string) string {
	var result string
	for i := range len(str) {
		if string(str[i]) == "\"" && string(str[i-1]) != "\\" {
			result += "\\\""
		} else {
			result += string(str[i])
		}
	}

	return result
}

func CompileAndExecute(imageName string, fileExt string, cmd string, sourceCode string) (types.ExecResult, error) {
	ctx := context.Background()
	containerName := "tmpContainer"
	var result types.ExecResult

	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return result, fmt.Errorf("failed to create Docker client: %v", err)
	}
	defer apiClient.Close()

	resp, err := apiClient.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		Cmd:   []string{"/bin/sh"},
		Tty:   true,
	}, nil, nil, nil, containerName)
	if err != nil {
		return result, fmt.Errorf("failed to create container: %v", err)
	}
	defer func() {
		apiClient.ContainerRemove(ctx, resp.ID, container.RemoveOptions{Force: true})
	}()

	if err := apiClient.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		return result, fmt.Errorf("failed to start container: %v", err)
	}

	shExec := []string{
		"sh",
		"-c",
	}
	filename := strings.Join([]string{"source", fileExt}, "")
	sanitizedSource := strings.Join([]string{"\"", sanitizeString(sourceCode), "\""}, "")
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
		return result, fmt.Errorf("failed to create exec instance: %v", err)
	}

	respExec, err := apiClient.ContainerExecAttach(ctx, execID.ID, container.ExecAttachOptions{})
	if err != nil {
		return result, fmt.Errorf("failed to attach to exec instance: %v", err)
	}
	defer respExec.Close()

	var stdout, stderr bytes.Buffer
	_, err = stdcopy.StdCopy(&stdout, &stderr, respExec.Reader)
	if err != nil {
		return result, fmt.Errorf("failed to read from std: %v", err)
	}

	inspect, err := apiClient.ContainerExecInspect(ctx, execID.ID)
	if err != nil {
		return result, fmt.Errorf("failed to inspect exec instance: %v", err)
	}

	result.ExitCode = inspect.ExitCode
	result.Stdout = stdout.String()
	result.Stderr = stderr.String()

	return result, nil
}
