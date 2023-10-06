package utils

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

func WriteFile(filePath string, content []byte) {
	err := ioutil.WriteFile(filePath, content, 0666)
	if err != nil {
		logrus.Error(err)
	}
}
