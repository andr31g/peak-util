package main

import (
	"github.com/urfave/cli/v2"
)

var (
	InputSamples = cli.IntSliceFlag{
		Name:     "samples",
		Aliases:  []string{"s"},
		Usage:    "List of comma-separated integer samples",
		Required: true,
	}
	IterationCount = cli.UintFlag{
		Name:     "iterations",
		Value:    1,
		Aliases:  []string{"i"},
		Usage:    "Number of iterations to run the algorithm",
		Required: false,
	}
)
