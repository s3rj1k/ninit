package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/s3rj1k/ninit/pkg/config"
	"github.com/s3rj1k/ninit/pkg/logger"
	"github.com/s3rj1k/ninit/pkg/sysinit"
	"github.com/s3rj1k/ninit/pkg/version"
)

func init() {
	if !config.IsHelpFlag() {
		return // return to main flow when no help was asked
	}

	// print help to stdout and exit
	fmt.Println(
		config.Help(
			config.DefaultEnvPrefix,
			config.GetApplicationName(),
			version.GetVersion(),
			version.GetBuildTime(),
		),
	)
	os.Exit(0)
}

func main() {
	c := config.New(
		logger.Create(
			config.DefaultLogPrefix,
			logger.DefaultFlags, // for debug purposes can be set to: 'log.Lmsgprefix | log.Lshortfile | log.Lmsgprefix'
			logger.DebugLevelLog,
		),
	)

	c.Log.Infof("Application: '%s', Version: '%s', BuildTime: '%s'\n",
		config.GetApplicationName(),
		version.GetVersion(),
		version.GetBuildTime(),
	)

	if err := c.Get(config.DefaultEnvPrefix); err != nil {
		c.Log.Fatalf("%v\n", err)
	}

	var wg sync.WaitGroup

	if err := sysinit.Run(c, &wg); err != nil {
		c.Log.Errorf("%v\n", sysinit.GetErrorMessage(err))
	}

	wg.Wait()
}
