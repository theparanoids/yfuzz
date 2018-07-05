// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/yahoo/yfuzz/cmd/yfuzz-cli/api"
	"github.com/yahoo/yfuzz/cmd/yfuzz-cli/config"
	"github.com/yahoo/yfuzz/pkg/version"
	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	// Set up viper
	config.Init()

	ascii()

	// Override the version printer to show client and server version.
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("Client %s\n", version.Version)

		serverVersion, err := api.GetServerVersion()
		if err != nil {
			fmt.Println(err.Error())
			color.Red("Error retrieving server version.")
			return
		}

		fmt.Printf("Server %s\n", serverVersion)
	}

	app := cli.NewApp()
	app.Name = "yfuzz"
	app.Version = version.Version
	app.Usage = "a simple command-line utility for yFuzz"
	app.Commands = []cli.Command{
		{
			Name:      "create",
			ShortName: "c",
			Usage:     "Create a job from a docker image",
			Action: func(c *cli.Context) error {
				image := c.Args().Get(0)
				return wrapError(api.CreateJob(image))
			},
		},
		{
			Name:      "list",
			ShortName: "ls",
			Usage:     "List all jobs",
			Action: func(c *cli.Context) error {
				return wrapError(api.ListJobs())
			},
		},
		{
			Name:      "status",
			ShortName: "st",
			Usage:     "Get the status of the specified job",
			Action: func(c *cli.Context) error {
				job := c.Args().Get(0)
				return wrapError(api.GetJobStatus(job))
			},
		},
		{
			Name:      "logs",
			ShortName: "l",
			Usage:     "Get the logs of the specified job",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "tail, t",
					Usage: "Number of tail lines to truncate logs to",
					Value: 0,
				},
			},
			Action: func(c *cli.Context) error {
				job := c.Args().Get(0)
				return wrapError(api.GetJobLogs(job, c.Int("tail")))
			},
		},
		{
			Name:      "delete",
			ShortName: "d",
			Usage:     "Delete the specified job",
			Action: func(c *cli.Context) error {
				job := c.Args().Get(0)
				return wrapError(api.DeleteJob(job))
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		color.Red("yFuzz encountered a fatal error: %s", err.Error())
		os.Exit(1)
	}
}

// Wrapper function to handle errors that occur and make an exit status
func wrapError(e error) error {
	if e == nil {
		return nil
	}

	color.Red(e.Error())

	return cli.NewExitError(errors.New(""), 1)
}

// Prints a fancy looking yfuzz word to the standard output.
func ascii() {
	color.Cyan("      ______ _   _  ______ ______")
	color.Cyan("      |  ___| | | ||___  /|___  /")
	color.Cyan(" _   _| |_  | | | |   / /    / / ")
	color.Cyan("| | | |  _| | | | |  / /    / /  ")
	color.Cyan("| |_| | |   | |_| |./ /___./ /___")
	color.Cyan(" \\__, \\_|    \\___/ \\_____/\\_____/")
	color.Cyan("  __/ |                          ")
	color.Cyan(" |___/                           \n")
	color.Yellow("+ ----=[ yFuzz %s ]\n\n", version.Version)
}
