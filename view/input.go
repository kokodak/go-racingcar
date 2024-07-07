package view

import (
	"fmt"
	"strings"
)

func GetCarNames() []string {
	var carNames string
	_, err := fmt.Scanln(&carNames)
	if err != nil {
		panic(err)
	}
	return strings.Split(carNames, ",")
}
