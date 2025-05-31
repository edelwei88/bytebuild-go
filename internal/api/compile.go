package api

import (
	"net/http"

	"github.com/edelwei88/bytebuild-go/internal/docker/compile"
	"github.com/edelwei88/bytebuild-go/internal/postgres"
	"github.com/edelwei88/bytebuild-go/internal/postgres/models"
	"github.com/edelwei88/bytebuild-go/internal/types"
	"github.com/gin-gonic/gin"
)

func Compile(c *gin.Context) {
	var opts struct {
		Language   string `json:"language" binding:"required"`
		Compiler   string `json:"compiler" binding:"required"`
		SourceCode string `json:"source_code" binding:"required"`
		Args       string `json:"args"`
	}

	err := c.ShouldBind(&opts)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	var language models.Language
	postgres.Postgres.Where(models.Language{
		Name: opts.Language,
	}).First(&language)
	if language.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "language not found",
		})
		return
	}

	var compiler models.Compiler
	postgres.Postgres.Where(models.Compiler{
		DockerImageName: opts.Compiler,
	}).First(&compiler)
	if compiler.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "compiler not found",
		})
		return
	}

	var result types.ExecResult
	switch language.Name {
	case "Python":
		result, err = compile.Python(compiler.DockerImageName, language.FileExtension, opts.SourceCode, opts.Args)
	case "Cpp":
		result, err = compile.Cpp(compiler.DockerImageName, language.FileExtension, opts.SourceCode, opts.Args)
	case "C":
		result, err = compile.C(compiler.DockerImageName, language.FileExtension, opts.SourceCode, opts.Args)
	case "Lua":
		result, err = compile.Lua(compiler.DockerImageName, language.FileExtension, opts.SourceCode, opts.Args)
	case "JavaScript":
		result, err = compile.JavaScript(compiler.DockerImageName, language.FileExtension, opts.SourceCode, opts.Args)
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to compile",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to compile",
		})
		return
	}

	user, _ := c.Get("user")
	compile := models.Compile{
		Compiler:   compiler,
		User:       user.(models.User),
		Arg:        opts.Args,
		ExitCode:   result.ExitCode,
		SourceCode: opts.SourceCode,
		Stdout:     result.Stdout,
		Stderr:     result.Stderr,
	}

	status := postgres.Postgres.Create(&compile)
	if status.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create compile",
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
