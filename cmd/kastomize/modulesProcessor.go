package main

import (
	"github.com/kyberorg/kastomize/cmd/kastomize/cmdreader"
	"github.com/kyberorg/kastomize/cmd/kastomize/config"
	"github.com/kyberorg/kastomize/cmd/kastomize/utils"
	"github.com/pieterclaerhout/go-log"
	"os/exec"
	"sort"
	"strings"
)

func ProcessModules() {
	//getting all available modules for this JDK
	rawModules := getAvailableModules()
	availableModules = extractModuleNames(rawModules)

	//getting requested modules
	requestedModules = readModulesFile()

	//check modules if they are valid and available
	validModules = validateModules()
}

func getAvailableModules() []string {
	// The command you want to run along with the argument
	cmd := exec.Command(javaCmd, "--list-modules")
	temporarySlice := cmdreader.ReadOutputToSlice(cmd)
	return temporarySlice
}

func extractModuleNames(rawSlice []string) []string {
	modNamesSlice := make([]string, 0)
	for _, rawModule := range rawSlice {
		parts := strings.Split(rawModule, "@")
		if len(parts) < 2 {
			log.Errorf("Malformed module: '%s' - ignoring \n", rawModule)
		} else {
			if len(strings.TrimSpace(parts[0])) > 0 {
				modNamesSlice = append(modNamesSlice, parts[0])
			} else {
				log.Errorf("Empty module name: %s - ignoring", parts)
			}
		}
	}
	return modNamesSlice
}

func readModulesFile() []string {
	cmd := exec.Command(catCmd, config.Args().ModulesFile)
	requestedModules := cmdreader.ReadOutputToSlice(cmd)
	return requestedModules
}

func validateModules() []string {
	validOnlyModules := make([]string, 0)

	for _, requestedModule := range requestedModules {
		//searching if module available
		moduleFound := utils.SearchString(availableModules, requestedModule)
		moduleNotFound := !moduleFound
		if moduleNotFound {
			log.Errorf("%s - NOT FOUND. Ignoring \n", requestedModule)
			continue
		}

		//searching is module already included
		sort.Strings(validOnlyModules)
		moduleFound = utils.SearchString(validOnlyModules, requestedModule)
		notFound := !moduleFound
		if notFound {
			log.Infof("%s - OK \n", requestedModule)
			validOnlyModules = append(validOnlyModules, requestedModule)
		} else {
			log.Errorf("%s - DUPLICATE. Already included. Please review your modules file.\n", requestedModule)
			continue
		}
	}

	return validOnlyModules
}
