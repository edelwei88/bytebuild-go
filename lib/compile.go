package lib

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func Python3(filename string) (stdout string, stderr error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %v", err)
	}

	volumePath := filepath.Join(homeDir, "Projects/go/src/github.com/edelwei88/bytebuild-go/files")
	if _, err := os.Stat(volumePath); os.IsNotExist(err) {
		return "", fmt.Errorf("volume path does not exist: %s", volumePath)
	}

	args := []string{
		"run",
		"--rm",
		"--name",
		"python3bytebuild",
		"-v",
		volumePath + ":/usr/src/bytebuild",
		"-w",
		"/usr/src/bytebuild",
		"python:3",
		"python",
		filename,
	}

	cmd := exec.Command("docker", args...)

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	err = cmd.Run()
	if err != nil {
		return stdoutBuf.String(), fmt.Errorf("command failed: %v, stderr: %s", err, stderrBuf.String())
	}

	return stdoutBuf.String(), nil
}

func Python2(filename string) (stdout string, stderr error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %v", err)
	}

	volumePath := filepath.Join(homeDir, "Projects/go/src/github.com/edelwei88/bytebuild-go/files")
	if _, err := os.Stat(volumePath); os.IsNotExist(err) {
		return "", fmt.Errorf("volume path does not exist: %s", volumePath)
	}

	args := []string{
		"run",
		"--rm",
		"--name",
		"python3bytebuild",
		"-v",
		volumePath + ":/usr/src/bytebuild",
		"-w",
		"/usr/src/bytebuild",
		"python:2",
		"python",
		filename,
	}

	cmd := exec.Command("docker", args...)

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	err = cmd.Run()
	if err != nil {
		return stdoutBuf.String(), fmt.Errorf("command failed: %v, stderr: %s", err, stderrBuf.String())
	}

	return stdoutBuf.String(), nil
}

func Lua54(filename string) (stdout string, stderr error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %v", err)
	}
	volumePath := filepath.Join(homeDir, "Projects/go/src/github.com/edelwei88/bytebuild-go/files")

	if _, err := os.Stat(volumePath); os.IsNotExist(err) {
		return "", fmt.Errorf("volume path does not exist: %s", volumePath)
	}

	args := []string{
		"run",
		"--rm",
		"--name",
		"lua-bytebuild",
		"-v",
		volumePath + ":/usr/src/bytebuild",
		"-w",
		"/usr/src/bytebuild",
		"nickblah/lua:5.4",
		"lua",
		filename,
	}

	cmd := exec.Command("docker", args...)

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	err = cmd.Run()
	if err != nil {
		return stdoutBuf.String(), fmt.Errorf("command failed: %v, stderr: %s", err, stderrBuf.String())
	}

	return stdoutBuf.String(), nil
}

func CLatest(filename string) (stdout string, stderr error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %v", err)
	}
	volumePath := filepath.Join(homeDir, "Projects/go/src/github.com/edelwei88/bytebuild-go/files")

	if _, err := os.Stat(volumePath); os.IsNotExist(err) {
		return "", fmt.Errorf("volume path does not exist: %s", volumePath)
	}

	args := []string{
		"run",
		"--rm",
		"--name",
		"c-bytebuild",
		"-v",
		volumePath + ":/usr/src/bytebuild",
		"-w",
		"/usr/src/bytebuild",
		"gcc:latest",
		"sh",
		"-c",
		fmt.Sprintf("gcc %s -o program && ./program", filename),
	}

	cmd := exec.Command("docker", args...)

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	err = cmd.Run()
	if err != nil {
		return stdoutBuf.String(), fmt.Errorf("command failed: %v, stderr: %s", err, stderrBuf.String())
	}

	return stdoutBuf.String(), nil
}

func RustLatest(filename string) (stdout string, stderr error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %v", err)
	}
	volumePath := filepath.Join(homeDir, "Projects/go/src/github.com/edelwei88/bytebuild-go/files")

	if _, err := os.Stat(volumePath); os.IsNotExist(err) {
		return "", fmt.Errorf("volume path does not exist: %s", volumePath)
	}

	if _, err := exec.LookPath("docker"); err != nil {
		return "", fmt.Errorf("docker command not found: %v", err)
	}

	args := []string{
		"run",
		"--rm",
		"--name",
		"rust-bytebuild",
		"-v",
		volumePath + ":/usr/src/bytebuild",
		"-w",
		"/usr/src/bytebuild",
		"rust:latest",
		"sh",
		"-c",
		fmt.Sprintf("rustc %s -o program && ./program", filename),
	}

	fmt.Printf("Executing: docker %s\n", args)

	cmd := exec.Command("docker", args...)

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	err = cmd.Run()
	if err != nil {
		return stdoutBuf.String(), fmt.Errorf("command failed: %v, stderr: %s", err, stderrBuf.String())
	}

	return stdoutBuf.String(), nil
}

func NodeLatest(filename string) (stdout string, stderr error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %v", err)
	}
	volumePath := filepath.Join(homeDir, "Projects/go/src/github.com/edelwei88/bytebuild-go/files")

	if _, err := os.Stat(volumePath); os.IsNotExist(err) {
		return "", fmt.Errorf("volume path does not exist: %s", volumePath)
	} else if err != nil {
		return "", fmt.Errorf("volume path access error: %v", err)
	}

	dockerPath, err := exec.LookPath("docker")
	if err != nil {
		return "", fmt.Errorf("docker command not found: %v", err)
	}
	fmt.Printf("Docker path: %s\n", dockerPath)

	if _, err := os.Stat("/var/run/docker.sock"); err != nil {
		return "", fmt.Errorf("cannot access docker socket: %v", err)
	}

	args := []string{
		"run",
		"--rm",
		"--name",
		"nodejs-bytebuild",
		"-v",
		volumePath + ":/usr/src/bytebuild",
		"-w",
		"/usr/src/bytebuild",
		"node:latest",
		"node",
		filename,
	}

	fmt.Printf("Executing: docker %s\n", args)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "docker", args...)

	cmd.Env = os.Environ()

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	err = cmd.Run()
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return stdoutBuf.String(), fmt.Errorf("command timed out after 30 seconds: %v, stderr: %s", err, stderrBuf.String())
		}
		if exitErr, ok := err.(*exec.ExitError); ok {
			return stdoutBuf.String(), fmt.Errorf("command failed with exit code %d, stderr: %s", exitErr.ExitCode(), stderrBuf.String())
		}
		return stdoutBuf.String(), fmt.Errorf("command failed: %v, stderr: %s", err, stderrBuf.String())
	}

	return stdoutBuf.String(), nil
}
