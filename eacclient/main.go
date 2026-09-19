package main

import (
	"github.com/Ambloplites/eacclient/eacclient/cmd"
	_ "github.com/Ambloplites/eacclient/eacclient/cmd/common"
	_ "github.com/Ambloplites/eacclient/eacclient/cmd/infinitas"
	_ "github.com/Ambloplites/eacclient/eacclient/cmd/konasute"
)

func main() {
	cmd.Execute()
}
