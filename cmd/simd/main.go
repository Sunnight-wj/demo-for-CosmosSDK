package main

import (
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"os"

	"simapp/app"
	"simapp/cmd/simd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
