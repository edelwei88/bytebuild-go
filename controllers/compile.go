package controllers

import (
	"net/http"
	"path/filepath"

	"github.com/edelwei88/bytebuild-go/initialize"
	"github.com/edelwei88/bytebuild-go/lib"
	"github.com/edelwei88/bytebuild-go/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func CompilePOST(c *gin.Context) {
	language := c.PostForm("language")
	compiler := c.PostForm("compiler")
	if language == "" || compiler == "" {
		c.String(http.StatusBadRequest, "Form error: specify language and compiler")
		return
	}
	var foundLanguage []models.Language
	initialize.DB.Where("name = ?", language).Preload(clause.Associations).Find(&foundLanguage)
	if len(foundLanguage) == 0 {
		c.String(http.StatusBadRequest, "Form error: wrong language")
		return
	}

	var foundCompiler models.Compiler
	foundCompilerBool := false
	for _, item := range foundLanguage[0].Compilers {
		if item.DockerImageName == compiler {
			foundCompiler = item
			foundCompilerBool = true
			break
		}
	}
	if !foundCompilerBool {
		c.String(http.StatusBadRequest, "Form error: wrong compiler")
		return
	}

	file, err := c.FormFile("source")
	if err != nil {
		c.String(http.StatusBadRequest, "Form error: no file with key 'source'")
		return
	}

	fileext := filepath.Ext(file.Filename)
	if fileext != foundLanguage[0].Extenstion {
		c.String(http.StatusBadRequest, "Form error: wrong file extension")
		return
	}

	c.SaveUploadedFile(file, "./files/"+filepath.Base(file.Filename))

	var stdout string
	var stderr error
	switch compiler {
	case "python:3":
		stdout, stderr = lib.Python3(filepath.Base(file.Filename))
	case "python:2":
		stdout, stderr = lib.Python2(filepath.Base(file.Filename))
	case "nickblah/lua:5.4":
		stdout, stderr = lib.Lua54(filepath.Base(file.Filename))
	case "gcc:latest":
		stdout, stderr = lib.CLatest(filepath.Base(file.Filename))
	case "rust:latest":
		stdout, stderr = lib.RustLatest(filepath.Base(file.Filename))
	case "node:latest":
		stdout, stderr = lib.NodeLatest(filepath.Base(file.Filename))

	default:
		c.JSON(http.StatusBadRequest, "Form error: wrong compiler")
		return
	}

	var compile models.Compile
	compile.Stdin = ""
	compile.Compiler = foundCompiler

	if stderr != nil {
		compile.Stderr = stderr.Error()
		compile.Stdout = ""
		initialize.DB.Create(&compile)
		c.JSON(http.StatusOK, gin.H{
			"stdout": nil,
			"stderr": stderr.Error(),
		})
		return
	}

	compile.Compiler = foundCompiler
	compile.Stderr = ""
	compile.Stdout = stdout
	initialize.DB.Create(&compile)
	c.JSON(http.StatusOK, gin.H{
		"stdout": stdout,
		"stderr": nil,
	})
}
