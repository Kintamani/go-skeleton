package main

import (
	"github.com/Kintamani/go-skeleton/internal/app"
	"github.com/Kintamani/go-skeleton/internal/config"
)

func init() {
	config.Load()
}

func main() {
	app.RunHTTP()
}
