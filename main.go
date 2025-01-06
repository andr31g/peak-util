package main

import (
	"fmt"
	"github.com/andr31g/peak-detector/peakdetect"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	run()
}

func run() {
	inputSamples := &InputSamples
	iterationCount := &IterationCount
	app := &cli.App{
		Name:  "peakdetect",
		Usage: "Peak Detector Utility",
		Commands: []*cli.Command{
			{
				Name:    "detect",
				Aliases: []string{"d"},
				Usage:   "Detect peaks in the provided list of comma-separated integer samples",
				Flags: []cli.Flag{
					inputSamples,
					iterationCount,
				},
				Action: func(c *cli.Context) error {
					iterations := c.Uint(iterationCount.Name)
					samples := c.IntSlice(inputSamples.Name)
					run0(iterations, samples)
					return nil
				},
			},
			{
				Name:    "test",
				Aliases: []string{"t"},
				Usage:   "Run tests",
				Action: func(c *cli.Context) error {
					test()
					return nil
				},
			},
			{
				Name:    "multipass",
				Aliases: []string{"m"},
				Usage:   "Run the multipass tests",
				Action: func(c *cli.Context) error {
					peakdetect.TestMultipass()
					return nil
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run0(iterations uint, samples []int) {
	if iterations > 0 {
		if iterations == 1 {
			primary := peakdetect.DetectPeaks(samples)
			peaks := primary.InflateWithCount(len(samples), &primary)
			fmt.Println("    Original Samples ", primary.GetSamples())
			fmt.Println("         Peak Values ", peaks)
			fmt.Println("Aligned Peak Indices ", peakdetect.AlignPeaksToSamplePositions(len(samples), primary.GetPeaks()))
			fmt.Println("        Peak Indices ", primary.GetPeaks())
		} else {
			fmt.Printf("Running %d iterations\n", iterations)
			fmt.Println("    Original Samples ", samples)
			if secondary, ok := peakdetect.IteratePeakDetect(iterations, samples); ok {
				peaks := secondary.InflateWithCount(len(samples), peakdetect.PrimaryValuesOnly[int](&secondary))
				fmt.Println("         Peak Values ", peaks)
				fmt.Println("Aligned Peak Indices ", peakdetect.AlignPeaksToSamplePositions(len(samples), secondary.GetPrimaryPeaks()))
				fmt.Println("        Peak Indices ", secondary.GetPrimaryPeaks())
			}
		}
	} else /* iterations == 0 */ {
		fmt.Println("    Original Samples ", samples)
	}
}

func test() {
	peakdetect.TestPeakDetectTriple()
	peakdetect.TestMergeOfTriples()
	peakdetect.TestCountBinary(8)
	peakdetect.TestCountDecimal()
}
