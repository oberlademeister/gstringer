package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/peterbourgon/ff"
	"github.com/sirupsen/logrus"
)

type context struct {
	Example string
	In      string
	Debug   bool
	Logger  *logrus.Logger
}

const cBinaryName = "gstringer"
const cEnvPrefix = "GSTRINGER"

func getContext() context {
	fs := flag.NewFlagSet(cBinaryName, flag.ExitOnError)
	var (
		debug   = fs.Bool("debug", false, "log debug information")
		example = fs.String("example", "", "if set, output yaml instead of reading")
		in      = fs.String("in", "", "yaml file to digest")
		_       = fs.String("config", "", "config file (optional)")
	)
	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [-in <YAMLFile>] [-example <YAMLFile]\n\n", cBinaryName)
		fs.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\neither in or example have to be set\nAll options can also be set via Environment by\nprefixing variables with %s.\nExample: %sDEBUG=1 eqlink\n", cEnvPrefix, cEnvPrefix)
		os.Exit(-1)
	}

	ff.Parse(fs, os.Args[1:],
		ff.WithConfigFileFlag("config"),
		ff.WithConfigFileParser(ff.PlainParser),
		ff.WithEnvVarPrefix(cEnvPrefix),
	)
	ctx := context{
		In:      *in,
		Example: *example,
		Debug:   *debug,
		Logger:  logrus.New(),
	}

	if ctx.Debug {
		ctx.Logger.SetLevel(logrus.DebugLevel)
	}

	if ctx.In == "" && ctx.Example == "" {
		fs.Usage()
	}

	return ctx
}
