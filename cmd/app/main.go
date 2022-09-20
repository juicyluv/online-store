package main

import (
	"github.com/juicyluv/online-store/internal/app"
)

func main() {
	a := app.New()

	if err := a.Run(); err != nil {
		panic(err)
	}
}
