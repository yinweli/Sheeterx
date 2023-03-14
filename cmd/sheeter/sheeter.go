package main

import (
	"github.com/spf13/cobra"

	"github.com/yinweli/Sheeterx/cmd/sheeter/build"
	"github.com/yinweli/Sheeterx/cmd/sheeter/version"
	"github.com/yinweli/Sheeterx/sheeter"
)

func main() {
	rootCommand := cobra.Command{
		Use:     sheeter.AppName,
		Long:    "Sheeter used to generate reader, sheeter, data from excel & sheet",
		Version: sheeter.Version,
	}
	rootCommand.AddCommand(build.NewCommand())
	rootCommand.AddCommand(version.NewCommand())
	rootCommand.CompletionOptions.HiddenDefaultCmd = true // 隱藏cobra提供的預設命令
	_ = rootCommand.Execute()
}
