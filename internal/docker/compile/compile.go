package compile

import (
	"errors"
	"fmt"

	"github.com/edelwei88/bytebuild-go/internal/docker"
	"github.com/edelwei88/bytebuild-go/internal/types"
)

func Python(imageName string, fileExtension string, sourceCode string, args string) (types.ExecResult, error) {
	cmd := fmt.Sprintf("python source.py %s", args)

	result, err := docker.CompileAndExecute(
		imageName,
		fileExtension,
		cmd,
		sourceCode)
	if err != nil {
		return types.ExecResult{}, errors.New("failed to compile")
	}

	return result, nil
}

func Cpp(imageName string, fileExtension string, sourceCode string, args string) (types.ExecResult, error) {
	cmd := fmt.Sprintf("g++ -o source.o source.cpp && ./source.o %s", args)

	result, err := docker.CompileAndExecute(
		imageName,
		fileExtension,
		cmd,
		sourceCode)
	if err != nil {
		return types.ExecResult{}, errors.New("failed to compile")
	}

	return result, nil
}

func C(imageName string, fileExtension string, sourceCode string, args string) (types.ExecResult, error) {
	cmd := fmt.Sprintf("gcc -o source.o source.c && ./source.o %s", args)

	result, err := docker.CompileAndExecute(
		imageName,
		fileExtension,
		cmd,
		sourceCode)
	if err != nil {
		return types.ExecResult{}, errors.New("failed to compile")
	}

	return result, nil
}

func Lua(imageName string, fileExtension string, sourceCode string, args string) (types.ExecResult, error) {
	cmd := fmt.Sprintf("lua source.lua %s", args)

	result, err := docker.CompileAndExecute(
		imageName,
		fileExtension,
		cmd,
		sourceCode)
	if err != nil {
		return types.ExecResult{}, errors.New("failed to compile")
	}

	return result, nil
}
