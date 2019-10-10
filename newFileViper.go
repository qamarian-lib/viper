package viper

import (
	"gopkg.in/qamarian-dtp/err.v0" //v0.4.0
	"gopkg.in/spf13/viper.v1"
        "path"
        "path/filepath"
      	"strings"
)

// NewFileViper () is a simple interface for creating file-based vipers.
//
// The first parametre of this function should be the path of the file. Relative and
// absolute paths are supported.
//
// This second parametre should be a file format supported by viper
// (github.com/spf13/viper). To put in simpler terms, this value must be a valid parametre
// for method someViper.SetConfigType (), as this parametre would be passed to the method.
//
// Method someViper.ReadInConfig () is not called by this function, so you would have to
// do it yourself.
//
// 	v1, err1 := viper.NewFileViper ("../someDir/file.ext", "json")
// 	v2, err2 := viper.NewFileViper ("/etc/file", "yaml")
// 	v3, err3 := viper.NewFileViper ("file.ext", "yaml")
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
