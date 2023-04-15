package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/easbarba/qas/internal/common"
	"github.com/easbarba/qas/internal/domain"
)

// Parse single configuration file,
// returns either properly parsed config parsed or empty struct.
//
//	TODO: check if the expect syntax is correct TODO: or err.
func ParseConfig(filepath string, filename string) (domain.Config, error) {
	var projects domain.Projects

	file, err := os.ReadFile(filepath)
	if err != nil {
		return domain.Config{}, err
	}

	err = json.Unmarshal(file, &projects)
	if err != nil {
		errMsg := fmt.Sprintf("Configuration file has incorrect syntax \n%s", err.Error())
		return domain.Config{}, errors.New(errMsg)
	}

	config := domain.Config{
		Lang:     common.FileNameWithoutExtension(filename),
		Projects: projects,
	}

	return config, nil
}

// TODO: Check for duplicates in configuration files
func ConfigCheckDuplicates() {
	panic("unimplemented")
}

func CheckConfigSyntax() error {
	panic("Not implemented")
}
