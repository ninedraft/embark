package embark

import (
	gpflag "github.com/octago/sflags/gen/gpflag"
	cobra "github.com/spf13/cobra"
)

type Config struct {
	// flag definitions here
	// https://github.com/octago/sflags#flags-based-on-structures------
}

func Command() *cobra.Command {
	var config = Config{}
	var cmd = &cobra.Command{
		Use: "command",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	if err := gpflag.ParseTo(&config, cmd.PersistentFlags()); err != nil {
		panic(err)
	}
	return cmd
}
