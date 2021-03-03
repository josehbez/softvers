package pm

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//InitCommand ...
type InitCommand struct {
}

// Run ...
func (c InitCommand) Run(ctx *Ctx) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "init",
		Short: "Create an empty .pm.yml or reinitialize an existing one",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			if err := ctx.Manifest.ReadInConfig(); err != nil {
				if _, ok := err.(viper.ConfigFileNotFoundError); ok {
					if errDV := ctx.InitManifest(); errDV != nil {
						ctx.Err.Fatalln(errDV)
					}
					ctx.Out.Printf("Initialized in %s", ctx.WorkingDir)

				} else {
					ctx.Err.Fatalln(err)
				}

			} else {
				ctx.Out.Println(fmt.Sprintf("Reinitialized in %s", ctx.WorkingDir))
			}
		},
	}
	return cmd
}
