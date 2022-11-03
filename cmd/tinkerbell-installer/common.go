package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type cobraFuncE func(cmd *cobra.Command, args []string) error

func handleErrors(logger *logrus.Logger, action cobraFuncE) cobraFuncE {
	return func(cmd *cobra.Command, args []string) error {
		err := action(cmd, args)
		if err != nil {
			logger.Errorf("Operation failed: %v.", err)
		}

		return err
	}
}
