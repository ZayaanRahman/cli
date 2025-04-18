package service

import (
	"github.com/docker/cli/cli"
	"github.com/docker/cli/cli/command"
	"github.com/spf13/cobra"
)

// NewServiceCommand returns a cobra command for `service` subcommands
func NewServiceCommand(dockerCli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "service",
		Short: "Manage Swarm services",
		Args:  cli.NoArgs,
		RunE:  command.ShowHelp(dockerCli.Err()),
		Annotations: map[string]string{
			"version": "1.24",
			"swarm":   "manager",
		},
	}
	cmd.AddCommand(
		newCreateCommand(dockerCli),
		newInspectCommand(dockerCli),
		newPsCommand(dockerCli),
		newListCommand(dockerCli),
		newRemoveCommand(dockerCli),
		newScaleCommand(dockerCli),
		newUpdateCommand(dockerCli),
		newLogsCommand(dockerCli),
		newRollbackCommand(dockerCli),
	)
	return cmd
}
