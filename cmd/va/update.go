package va

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/sailpoint-oss/sailpoint-cli/internal/terminal"
	"github.com/sailpoint-oss/sailpoint-cli/internal/va"
	"github.com/spf13/cobra"
)

func updateAndRebootVA(endpoint, password string) {
	log.Info("Attempting to update", "VA", endpoint)

	update, updateErr := va.RunVACmd(endpoint, password, UpdateCommand)
	if updateErr != nil {
		log.Error("Problem updating", "VA", endpoint, "err", updateErr, "resp", update)
	} else {
		log.Info("Virtual appliance updating", "VA", endpoint)
		reboot, rebootErr := va.RunVACmd(endpoint, password, RebootCommand)
		if rebootErr != nil && rebootErr.Error() != "wait: remote command exited without exit status or exit signal" {
			log.Error("Problem rebooting", "Server", endpoint, "err", rebootErr, "resp", reboot)
		} else {
			log.Info("Virtual appliance rebooting", "VA", endpoint)
		}
	}

	fmt.Println()
}

func newUpdateCommand(term terminal.Terminal) *cobra.Command {
	var credentials []string
	cmd := &cobra.Command{
		Use:     "update",
		Short:   "Perform update operations on a SailPoint virtual appliance",
		Long:    "\nPerform update operations on a SailPoint virtual appliance\n\n",
		Example: "sail va update 10.10.10.10 10.10.10.11 -c S@ilp0int -c S@ilp0int",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			for i, endpoint := range args {
				var password string

				if len(credentials) > i {
					password = credentials[i]
				}

				if password == "" {
					password, _ = term.PromptPassword("Enter password for " + endpoint + ":")
				}

				updateAndRebootVA(endpoint, password)
			}
			return nil
		},
	}

	cmd.Flags().StringArrayVarP(&credentials, "Passwords", "p", []string{}, "You can enter the passwords for the servers in the same order that the servers are listed as arguments")

	return cmd
}
