package versions

import (
	"io/ioutil"

	"github.com/coreos/go-semver/semver"
	"github.com/sirupsen/logrus"
)

const (
	MAJOR = "major"
	MINOR = "minor"
	PATCH = "patch"
)

func UpdateVersionFile(file, deploymentType string) string {
	oldVersion := readInputFile(file)
	version := getNewVersion(oldVersion, deploymentType)
	writeOutputFile(file, version)
	return version
}

func readInputFile(file string) string {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		logrus.Panicf("Cannot read file %v \n %v", file, err)
	}

	return string(content)
}

func writeOutputFile(file, content string) {
	err := ioutil.WriteFile(file, []byte(content), 644)
	if err != nil {
		logrus.Panicf("Cannot write version %v to file %v \n", content, file, err)
	}
}

func getNewVersion(v, d string) string {
	version := semver.New(v)
	if d == MAJOR {
		version.BumpMajor()
	} else if d == MINOR {
		version.BumpMinor()
	} else if d == PATCH {
		version.BumpPatch()
	}
	return version.String()
}
