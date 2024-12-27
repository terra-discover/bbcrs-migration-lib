package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
)

type ListGenModelMigrationRequest []GenModelMigrationRequest

type GenModelMigrationRequest struct {
	ModelsFilePath string // beware, based on real path
	ImportAlias    string // only alias, can be anything
	ContentHeader  string // header of content
}

func genModelMigrations(listReq ListGenModelMigrationRequest, myPackage, writeFilePath string) {
	wd, err := os.Getwd()
	if err != nil {
		log.Panicln("genModelMigrations, os.Getwd:", err)
		return
	}

	allImportContent := ""
	allModelsContent := ""

	// For comment index in models content
	startIdx := 0

	for _, req := range listReq {
		log.Println("Start genModelMigrations:", req.ContentHeader)

		pkgName := path.Join(myPackage, req.ModelsFilePath)
		allImportContent += fmt.Sprintf("%s \"%s\" \n", req.ImportAlias, pkgName)

		dir := path.Join(wd, req.ModelsFilePath)
		modelContent, lastIdx := getModelContent(dir, req.ImportAlias, startIdx)
		allModelsContent += fmt.Sprintf("// %s \n", req.ContentHeader)
		allModelsContent += modelContent

		// Incremenet startIdx
		startIdx = lastIdx + 1

		log.Println("End genModelMigrations:", req.ContentHeader)
	}

	writePath := path.Join(wd, writeFilePath)

	contentFormat :=
		`package migration

		import (
			%s
		)

		// ModelMigrations models to migrate
		var ModelMigrations []interface{} = []interface{}{
			%s
		}`

	content := fmt.Sprintf(
		contentFormat,
		allImportContent,
		allModelsContent,
	)

	buf := bytes.NewBufferString(content)
	contentByte := buf.Bytes()

	err = os.WriteFile(writePath, contentByte, os.ModeExclusive)
	if err != nil {
		log.Panicln("genModelMigrations, os.WriteFile:", err)
		return
	}
}

func getModelContent(dir, importModelAlias string, startIdx int) (resultContent string, lastIdx int) {
	prefix := "type"
	suffix := "struct"
	// Will return `type ModelName struct`
	rawPattern := fmt.Sprintf(`%s [\s\S]*? %s`, prefix, suffix)

	pattern, err := regexp.Compile(rawPattern)
	if err != nil {
		log.Panicln("genModelMigrations, regexp.Compile:", err)
		return
	}

	listModelName := []string{}

	fileDirs, err := os.ReadDir(dir)
	if err != nil {
		log.Panicln("getModelContent, os.ReadDir:", err)
		return
	}

	for _, fileDir := range fileDirs {
		if fileDir.IsDir() {
			continue
		}

		filePath := path.Join(dir, fileDir.Name())
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			log.Panicln("getModelContent, fileDirs os.ReadFile:", err)
			continue
		}

		listSubmatch := pattern.FindStringSubmatch(string(fileContent))
		for _, submatch := range listSubmatch {
			// Split by space
			listStr := strings.Split(submatch, " ")
			for _, str := range listStr {
				if str != prefix && str != suffix {
					listModelName = append(listModelName, str)
				}
			}
		}
	}

	modelsContent := ""

	for idx, modelName := range listModelName {
		// Will return &ModelName{}
		modelsContent += fmt.Sprintf(`&%s.%s{}`, importModelAlias, modelName)
		lastIdx = startIdx + idx
		modelsContent += fmt.Sprintf(", //index: %d \n", lastIdx)
	}

	resultContent = modelsContent

	return
}
