package main

import (
	"github.com/JoelOtter/ggapp/cmd/ggapp/export"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	logger := logrus.StandardLogger()

	cmd := &cobra.Command{
		Use: "ggapp",
	}
	cmd.AddCommand(export.NewCommand(logger))

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
