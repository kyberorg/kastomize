package main

import (
	"fmt"
	"github.com/kyberorg/kastomize/cmd/kastomize/config"
	"github.com/pieterclaerhout/go-log"
)

const (
	javaCmd = "java"
	catCmd  = "cat"
)

var (
	requestedModules []string
	availableModules []string
	validModules     []string
)

func main() {
	// Print the log timestamps
	log.PrintTimestamp = false

	//validate args - if something wrong program exists
	config.ValidateArgs()

	//if we here - all are valid
	log.Info("Processing modules...")
	log.Info("")

	ProcessModules()

	log.Info("")
	log.Info("Processing modules...DONE")
	fmt.Println("Will create JDK with following modules:", validModules)

	log.Info("")
	log.Info("Creating your custom JVM")
	createCustomJDK()
}

func createCustomJDK() {
	//make command and args
	jlinkCommand := MakeJLinkCommand()
	//launch it
	ExecuteJLink(jlinkCommand)
}
