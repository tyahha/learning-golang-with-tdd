package main

import (
	"github.com/tyahha/learning-golang-with-tdd/math/clockface"
	"os"
	"time"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
