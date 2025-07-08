package cli

import (

	"github.com/sajad-dev/gear/internal/command/migration"
	"github.com/sajad-dev/gear/internal/command/run"
	"github.com/spf13/cobra"
)

var CommandsList []*Commands = []*Commands{
	{
		Command: &cobra.Command{
			Use:   "migration",
			Short: "Run migration list",
			Run: func(cmd *cobra.Command, args []string) {
				migration.Migrate()
			},
		},
		Flags: []any{},
	},
	{
		Command: &cobra.Command{
			Use:   "run",
			Short: "Run http server ",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				run.Http(cmd,args)
			},
		},
		Flags: []any{FlagInt{Flag: "port", ShortFlag: "p", Defualt: 8080, Discription: "Run server in your port"}},
	},
}

type FlagString struct {
	Value       string
	Flag        string
	ShortFlag   string
	Defualt     string
	Discription string
}
type FlagInt struct {
	Value       int
	Flag        string
	ShortFlag   string
	Defualt     int
	Discription string
}
type Commands struct {
	Command *cobra.Command
	Flags   []any
}
