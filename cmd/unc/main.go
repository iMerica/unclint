package main

import "github.com/iMerica/unclint/internal/cli"

var version = "dev"

func main() {
	cli.Execute(version)
}
