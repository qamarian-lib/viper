package viper

import (
	"gopkg.in/spf13/viper.v1"
	"os"
        "path"
        "path/filepath"
      	"strings"
)

func NewFileViper (filePath string, format string) (*viper.Viper, error) {

	// Verifying if file exists. ..1.. {
        _, errX := os.Open (filePath)
        if errX != nil {
                return nil, errX
        }
        // ..1.. }

        // Extracting the name (without extension) and the directory of the file. ..1.. {
        fileName := filepath.Base (filePath)

        fileDir := strings.TrimSuffix (filePath, dataFileName)
        if dataFileDir != "" {
                fileDir = fileDir [ : len (fileDir)]
        } else {
                fileDir = "."
        }

        fileNameWithoutExt := strings.TrimSuffix (fileName, path.Ext (fileName))
        // ..1.. }

        // Creating viper. ..1.. {
        v := viper.New ()
        v.SetConfigName (fileNameWithoutExt)
        v.SetConfigType (format)
        v.AddConfigPath (fileDir)
        errY := v.ReadInConfig ()
        // ..1.. }

        return v, err2
}
