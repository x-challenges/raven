package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"golang.org/x/exp/slices"
)

// Files contains slice of config file for loading, optional.
var Files = []string{}

// FileType type.
type FileType string

const (
	FileYAML FileType = "yaml"
	FileYML  FileType = "yml"
	FileJSON FileType = "json"
)

// FileDefaultType for configurations.
var FileDefaultType = FileYAML

// Enums returns slice of file type enumerations.
func (FileType) Enums() []FileType {
	return []FileType{
		FileYAML,
		FileYML,
		FileJSON,
	}
}

// Setup config file.
func setupConfigFiles(viper *viper.Viper, files ...string) error {
	for i, file := range files {
		var (
			fileType = FileDefaultType
		)

		// skip if config file is empty
		if file == "" {
			continue
		}

		// check file exists
		if _, err := os.Stat(file); err != nil {
			return err
		}

		if ext := filepath.Ext(file); ext != "" {
			fileType = FileType(ext[1:])
		}

		// check that file type exists
		if !slices.Contains(new(FileType).Enums(), fileType) {
			return fmt.Errorf("unexpected config file type: %s", fileType)
		}

		// setup config
		viper.SetConfigType(string(fileType))
		viper.SetConfigFile(file)

		var err error

		if i == 0 {
			err = viper.ReadInConfig()
		} else {
			err = viper.MergeInConfig()
		}

		if err != nil {
			return err
		}
	}
	return nil
}
