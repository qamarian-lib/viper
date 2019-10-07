package viper

import (
	"gopkg.in/qamarian-lib/str.v3"
	"os"
	"testing"
)

func TestNewFileViper (t *testing.T) {
	str.PrintEtr ("Test started!", "std", "TestNewFileViper ()")

	// Testing if function works with absolute paths. ..1.. {
	pwd, errX := os.Getwd ()
	if errX != nil {
		str.PrintEtr ("Test failed! Ref: 0", "err", "TestNewFileViper ()")
		t.FailNow ()
	}

	if pwd == "/" {
		pwd = ""
	}

	vY, errY := NewFileViper (pwd + "/" + confFileName, "yaml")
	if vY == nil || errY != nil {
		str.PrintEtr ("Test failed! Ref: 1", "err", "TestNewFileViper ()")
		t.FailNow ()
	}
	// ..1.. }

	// Testing if function works with relative paths. ..1.. {
	vZ, errZ := NewFileViper ("../viper/" + confFileName, "yaml")
	if vZ == nil || errZ != nil {
		str.PrintEtr ("Test failed! Ref: 2", "err", "TestNewFileViper ()")
		t.FailNow ()
	}
	// ..1.. }

	// Testing if fuction works with local paths. ..1.. {
	vB, errB := NewFileViper ("./" + confFileName, "yaml")
	if vB == nil || errB != nil {
		str.PrintEtr ("Test failed! Ref: 4", "err", "TestNewFileViper ()")
		t.FailNow ()
	}
	// ..1.. }

	str.PrintEtr ("Test passed!", "std", "TestNewFileViper ()")
}

var (
	confFileName string = "/newFileViper_test_confFile.yml"
)
