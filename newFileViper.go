package viper

import (
	"gopkg.in/qamarian-dtp/err.v0" //v0.4.0
	"gopkg.in/spf13/viper.v1"
        "path"
        "path/filepath"
      	"strings"
)

// NewFileViper () is a simple interface for creating file vipers. Argument 0 should be
// the path of the file, while argument 1 should be a file format supported by viper
// (github.com/spf13/viper).
//
// To put in simpler terms, the value of argument 1 should be a valid argument for method
// someViper.SetConfigPath (), as this argument would be passed to the method.
//
// Also, the value argument 0 could be either a relative path or an absolute path. Bare
// file names are not allowd. To use just a file name, do something like this
// "./filename.ext", and not "filename.ext".
//
// Method someViper.ReadInConfig () do not need to be called on vipers created by this
// function, as the method would have already been called by the function.
//
// 	v1, err1 := viper.NewFileViper ("../someDir/file.ext", "json")
// 	v2, err2 := viper.NewFileViper ("/etc/file", "yaml")
//
func NewFileViper (filePath string, format string) (*viper.Viper, error) {
        // Extracting the name (without extension) and the directory of the file. ..1.. {
        fileName := filepath.Base (filePath)

        fileDir := strings.TrimSuffix (filePath, fileName)
        if fileDir != "" {
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
	if errY != nil {
		return nil, err.New ("Unable to create viper.", nil, nil, errY)
	}
        // ..1.. }

        return v, nil
}
