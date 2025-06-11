package docker

import (
	"fmt"

	"github.com/edelwei88/bytebuild-go/internal/types"
)

func Python(imageName string, fileExtension string, sourceCode string, args string) (types.ExecResult, error) {
	cmd := fmt.Sprintf("python source.py %s", args)

	result, err := CompileAndExecute(
		imageName,
		fileExtension,
		cmd,
		sourceCode)
	if err != nil {
		return types.ExecResult{}, err
	}

	return result, nil
}

func Cpp(imageName string, fileExtension string, sourceCode string, args string) (types.ExecResult, error) {
	cmd := fmt.Sprintf("g++ -o source.o source.cpp && ./source.o %s", args)

	result, err := CompileAndExecute(
		imageName,
		fileExtension,
		cmd,
		sourceCode)
	if err != nil {
		return types.ExecResult{}, err
	}

	return result, nil
}

func C(imageName string, fileExtension string, sourceCode string, args string) (types.ExecResult, error) {
	cmd := fmt.Sprintf("gcc -o source.o source.c && ./source.o %s", args)

	result, err := CompileAndExecute(
		imageName,
		fileExtension,
		cmd,
		sourceCode)
	if err != nil {
		return types.ExecResult{}, err
	}

	return result, nil
}

func Lua(imageName string, fileExtension string, sourceCode string, args string) (types.ExecResult, error) {
	cmd := fmt.Sprintf("lua source.lua %s", args)

	result, err := CompileAndExecute(
		imageName,
		fileExtension,
		cmd,
		sourceCode)
	if err != nil {
		return types.ExecResult{}, err
	}

	return result, nil
}

func JavaScript(imageName string, fileExtension string, sourceCode string, args string) (types.ExecResult, error) {
	cmd := fmt.Sprintf("node source.js %s", args)

	result, err := CompileAndExecute(
		imageName,
		fileExtension,
		cmd,
		sourceCode)
	if err != nil {
		return types.ExecResult{}, err
	}

	return result, nil
}

func Rust(imageName string, fileExtension string, sourceCode string, args string) (types.ExecResult, error) {
	cmd := fmt.Sprintf("rustc source.rs && ./source %s", args)

	result, err := CompileAndExecute(
		imageName,
		fileExtension,
		cmd,
		sourceCode)
	if err != nil {
		return types.ExecResult{}, err
	}

	return result, nil
}

func Java(imageName string, fileExtension string, sourceCode string, args string) (types.ExecResult, error) {
	cmd := fmt.Sprintf("javac source.java && java App %s", args)

	result, err := CompileAndExecute(
		imageName,
		fileExtension,
		cmd,
		sourceCode)
	if err != nil {
		return types.ExecResult{}, err
	}

	return result, nil
}

func Go(imageName string, fileExtension string, sourceCode string, args string) (types.ExecResult, error) {
	cmd := fmt.Sprintf("go run source.go %s", args)

	result, err := CompileAndExecute(
		imageName,
		fileExtension,
		cmd,
		sourceCode)
	if err != nil {
		return types.ExecResult{}, err
	}

	return result, nil
}

func Ruby(imageName string, fileExtension string, sourceCode string, args string) (types.ExecResult, error) {
	cmd := fmt.Sprintf("ruby source.rb %s", args)

	result, err := CompileAndExecute(
		imageName,
		fileExtension,
		cmd,
		sourceCode)
	if err != nil {
		return types.ExecResult{}, err
	}

	return result, nil
}

func PHP(imageName string, fileExtension string, sourceCode string, args string) (types.ExecResult, error) {
	cmd := fmt.Sprintf("php source.php %s", args)

	result, err := CompileAndExecute(
		imageName,
		fileExtension,
		cmd,
		sourceCode)
	if err != nil {
		return types.ExecResult{}, err
	}

	return result, nil
}

func Elixir(imageName string, fileExtension string, sourceCode string, args string) (types.ExecResult, error) {
	cmd := fmt.Sprintf("elixir source.exs %s", args)

	result, err := CompileAndExecute(
		imageName,
		fileExtension,
		cmd,
		sourceCode)
	if err != nil {
		return types.ExecResult{}, err
	}

	return result, nil
}
