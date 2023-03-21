package main

//
//import (
//	"errors"
//	"fmt"
//	"os"
//
//	"github.com/sirupsen/logrus"
//	"github.com/spf13/cobra"
//
//	"github.com/moadqassem/kubetink/pkg/util/yamled"
//)
//
//type cobraFuncE func(cmd *cobra.Command, args []string) error
//
//func handleErrors(logger *logrus.Logger, action cobraFuncE) cobraFuncE {
//	return func(cmd *cobra.Command, args []string) error {
//		err := action(cmd, args)
//		if err != nil {
//			logger.Errorf("Operation failed: %v.", err)
//		}
//
//		return err
//	}
//}
//
//func loadHelmValues(filename string) (*yamled.Document, error) {
//	if filename == "" {
//		return nil, errors.New("no file specified via --helm-values flag")
//	}
//
//	f, err := os.Open(filename)
//	if err != nil {
//		return nil, err
//	}
//	defer f.Close()
//
//	values, err := yamled.Load(f)
//	if err != nil {
//		return nil, fmt.Errorf("failed to decode %s: %w", filename, err)
//	}
//
//	return values, nil
//}
