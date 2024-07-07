package main

import (
	"fmt"
	"github.com/kokodak/go-racingcar/view"
)

func main() {
	carNames := view.GetCarNames()
	fmt.Println(carNames)
}
