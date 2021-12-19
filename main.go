package main

import (
	"embed"
	"github.com/SSH-Management/server/cmd/cli"
)

var (
	Version = "dev"

	//go:embed ui/build/*
	Ui embed.FS
)

func main() {
	cli.Execute(Ui)
}
