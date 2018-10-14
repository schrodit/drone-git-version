package version

import (
	"io/ioutil"
	"strings"

	"github.com/coreos/go-semver/semver"
	"github.com/sirupsen/logrus"
)

func UpdateVersionFile(input_file, output_file string) {
	versions := readInputFile(input_file)
	version := getLatestVersion(versions)
	writeOutpUtFile(output_file, version)
}

func readInputFile(file string) []string {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		logrus.Panicf("Cannot read file %v \n %v", file, err)
	}

	return strings.Split(string(content), ",")
}

func writeOutpUtFile(file, content string) {
	err := ioutil.WriteFile(file, []byte(content), 644)
	if err != nil {
		logrus.Panicf("Cannot write version %v to file %v \n", content, file, err)
	}
}

func getLatestVersion(versions []string) string {
	var latest *semver.Version
	for _, v := range versions {
		if v == "latest" {
			continue
		}
		version := semver.New(v)
		if latest == nil {
			latest = version
			continue
		}
		if latest.LessThan(*version) {
			latest = version
		}
	}
	return latest.String()
}
