// Copyright (c) 2023, SailPoint Technologies, Inc. All rights reserved.
package sdk

import (
	"embed"

	"github.com/sailpoint-oss/sailpoint-cli/internal/initialize"
	"github.com/spf13/cobra"
)

//go:embed typescript/*
var tsTemplateContents embed.FS

const tsTemplateDirName = "typescript"

func newTypescriptCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "typescript",
		Short:   "Perform Search operations in IdentityNow using a predefined search template",
		Long:    "\nPerform Search operations in IdentityNow using a predefined search template\n\n",
		Example: "sail search template",
		Aliases: []string{"ts"},
		Args:    cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var projName string
			var err error

			if len(args) > 0 {
				projName = args[0]
			} else {
				projName = "typescript-template"
			}

			err = initialize.InitializeProject(tsTemplateContents, tsTemplateDirName, projName)
			if err != nil {
				return err
			}

			return nil
		},
	}
	return cmd
}
