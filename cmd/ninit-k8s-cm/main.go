package main

import (
	"context"
	"os"
	"sync"

	config "github.com/s3rj1k/ninit/pkg/config/k8s/cm"
	"github.com/s3rj1k/ninit/pkg/k8s/configmap"
	"github.com/s3rj1k/ninit/pkg/log/logger"
	"github.com/s3rj1k/ninit/pkg/log/standart"
	"github.com/s3rj1k/ninit/pkg/sysinit"
	"github.com/s3rj1k/ninit/pkg/utils"
	"github.com/s3rj1k/ninit/pkg/version"
	"k8s.io/klog/v2"
)

func main() {
	log := standart.Create(
		os.Stdout,
		config.DefaultLogPrefix,
		standart.DefaultFlags, // for debug purposes can be set to: 'log.Lmsgprefix | log.Lshortfile | log.Lmsgprefix'
		logger.InfoLevelLog,
	)

	klog.SetLogger(log)

	c := config.New(
		config.DefaultEnvPrefix,
	)

	if utils.IsHelpFlag() {
		c.Help(
			utils.GetApplicationName(),
			version.GetVersion(),
			version.GetBuildTime(),
		)

		os.Exit(0)
	}

	log.Infof("Application: '%s', Version: '%s', BuildTime: '%s'\n",
		utils.GetApplicationName(),
		version.GetVersion(),
		version.GetBuildTime(),
	)

	if err := c.Get(); err != nil {
		log.Fatalf("%v\n", err)
	}

	if c.GetVerboseLogging() {
		log.SetLevel(logger.TraceLevelLog)
	}

	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	defer sysinit.Cleanup(&wg, cancel, log)

	if err := configmap.Run(ctx, &wg, c, log); err != nil {
		log.Errorf("%v\n", err)

		return
	}

	if err := sysinit.Run(ctx, &wg, c, log); err != nil {
		log.Errorf("%v\n", sysinit.GetErrorMessage(err))
	}
}
