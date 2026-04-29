package main

import "github.com/yourname/unc/internal/cli"

var version = "dev"

func main() {
	cli.Execute(version)
}
