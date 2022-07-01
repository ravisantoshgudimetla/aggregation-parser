package main

import (
	"os"

	cmd "github.com/ravisantoshgudimetla/aggregation-parser/cmd"
)

func main() {
	url := os.Args[1]
	cmd.Parse(url)
}
