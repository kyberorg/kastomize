package config

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

const BinDir = "bin"
const JLinkCmd = "jlink"
const JavacCmd = "javac"

//internal vars
var (
	args *appArgs
)

type appArgs struct {
	ModulesFile string
	OutputDir   string
	JavaHome    string
}

func init() {
	args = &appArgs{}

	kingpin.Flag("modules-file", "Path to file with list of modules").Required().
		ExistingFileVar(&args.ModulesFile)

	kingpin.Flag("output", "Output directory (full path) - should not exist").Required().
		StringVar(&args.OutputDir)

	kingpin.Flag("java-home", "Path to JDK").Required().Envar("JAVA_HOME").
		ExistingDirVar(&args.JavaHome)

	kingpin.Version("0.0.1")
	kingpin.Parse()

}

func Args() *appArgs {
	return args
}
