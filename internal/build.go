package internal

import (
	"os"
	"path"
)

// BuildProject packages the project as defined in the configuration file.
func BuildProject(config Config) error {
	var err error
	_ = os.Mkdir("temp", 0755)
	exportsFile, err := os.Open(ExportsFilePath)
	if err != nil {
		return err
	}
	defer exportsFile.Close()

	directory, err := compileCatrinaJS(exportsFile)
	if err != nil {
		return err
	}

	inputFile, err := os.Open(config.MainJS)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	imports, lines := getImportsJSInputFile(inputFile)
	var files []string
	for _, v := range imports {
		for _, imp := range directory {
			for _, name := range imp.names {
				if name == v {
					files = safeAppend(files, imp.path)
				}
			}
		}
	}

	for _, file := range files {
		f, err := os.Open(path.Join("lib", file))
		if err != nil {
			return err
		}
		imports, _ = getImportsJS(f, imports)
		_ = f.Close()
	}

	finalJS := path.Join(config.BuildPath, config.BuildJS)
	_ = os.Remove(finalJS)

	buildFilePath := "./temp/main.build.js"
	_ = os.Remove(buildFilePath)
	buildFile, err := os.OpenFile(buildFilePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		return err
	}
	defer buildFile.Close()

	catrinaJS, err := os.Open(path.Join("temp", CompileFileJs))
	if err != nil {
		return err
	}
	defer catrinaJS.Close()

	err = writeCatrinaCoreJS(buildFile)
	if err != nil {
		return err
	}

	err = writeImportsJS(catrinaJS, buildFile, imports)
	if err != nil {
		return err
	}

	err = writeFinalFileJS(buildFile, config.MainJS, lines)
	if err != nil {
		return err
	}

	err = os.Rename(buildFilePath, finalJS)
	if err != nil {
		return err
	}

	inputFileCSS, err := os.Open(config.MainCSS)
	if err != nil {
		return err
	}
	defer inputFileCSS.Close()

	imports, lines = getImportsCSS(inputFileCSS, []string{})

	buildFilePathCSS := "./temp/styles.build.css"

	buildFileCSS, err := os.OpenFile(buildFilePathCSS, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		return err
	}
	defer buildFileCSS.Close()

	var imp []string
	for _, v := range imports {
		f, err := os.Open(v)
		if err != nil {
			return err
		}

		imp, _ = getImportsCSS(f, imp)

		_ = f.Close()
	}

	for _, v := range imp {
		imports = safeAppend(imports, v)
	}

	err = writeImportsCSSAndCopyFonts(buildFileCSS, imports, config)
	if err != nil {
		return err
	}

	err = writeFinalFileJS(buildFileCSS, config.MainCSS, lines)
	if err != nil {
		return err
	}

	err = os.Rename(buildFilePathCSS, path.Join(config.BuildPath, config.BuildCSS))

	_ = os.RemoveAll("temp")
	return err
}
