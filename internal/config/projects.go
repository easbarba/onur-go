package config

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

// Raw structure of Configuration files
// log config files found
type Raw struct {
	Lang     string `json:"lang"`
	Projects []struct {
		Name   string `json:"name"`
		Branch string `json:"branch"`
		URL    string `json:"url"`
	} `json:"projects"`
}

// All configuration files unmarshallowed
func All() []Raw {
	var result []Raw
	files := files()

	fmt.Println("Configuration files found: ")

	for _, file := range files {
		p := path.Join(Folder(), file.Name())
		fmt.Println(p)

		fileInfo, err := os.Stat(p)

		// ignore broken symbolic link
		if os.IsNotExist(err) {
			continue
		}

		// ignore directories
		if fileInfo.IsDir() {
			continue
		}

		// ignore csv files (legacy)
		if ext := filepath.Ext(p); ext == ".csv" {
			continue
		}

		parsed := parse(p)
		result = append(result, parsed)
	}

	return result
}

// array of configuratiion file name
func files() []fs.FileInfo {
	files, err := ioutil.ReadDir(Folder())

	if err != nil {
		log.Fatal(err)
	}

	return files
}

func parse(filepath string) Raw {
	file, err := ioutil.ReadFile(filepath)
	var proj Raw

	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(file, &proj)

	if err != nil {
		fmt.Println(err)
	}

	return proj
}

func home() string {
	home, _ := os.UserHomeDir()

	return home
}

// HomeFolder that all projects repositories will be stored at
func HomeFolder() string {
	result := path.Join(home(), "Projects")

	return result
}

// Folder that config files will be looked up for
func Folder() string {
	result := path.Join(home(), ".config", "qas")

	return result
}
