package main

import (
	"embed"

	"github.com/SSH-Management/server/cmd/cli"
)

var Version = "dev"

//go:embed migrations/*
var migrations embed.FS

func main() {
	cli.Execute(Version, migrations)
}
