package config

import (
	"log"
	"os"
	"syscall"
)

func ValidateArgs() {
	isModulesFileReadable()

	isJdk := isJavaHomePointsToJDK()
	if !isJdk {
		log.Fatalf("JavaHome should point to JDK. " +
			"Please adjust JAVA_HOME env variable or set JavaHome using --java-home=/path/to/jdk \n")
	}

	jLinkExists := checkIfJLinkExists()
	if !jLinkExists {
		log.Fatalf("there no %s binary found in JDK %s\n", JLinkCmd, args.JavaHome)
	}
}

func isModulesFileReadable() {
	info, err := os.Stat(args.ModulesFile)
	modulesFileExists := !os.IsNotExist(err) && info.Mode().IsRegular()
	if modulesFileExists {
		err = syscall.Access(args.ModulesFile, syscall.O_RDONLY)
		if err != nil {
			log.Fatalf("Cannot read modules file: '%s'. Permission denied.")
		}
	} else {
		log.Fatalf("Modules file '%s'not exist. Please check your configuration \n", args.ModulesFile)
	}

}

func isJavaHomePointsToJDK() bool {
	javac := args.JavaHome + string(os.PathSeparator) + BinDir + string(os.PathSeparator) + JavacCmd
	info, err := os.Stat(javac)
	return !os.IsNotExist(err) && info.Mode().IsRegular()
}

func checkIfJLinkExists() bool {
	jLink := args.JavaHome + string(os.PathSeparator) + BinDir + string(os.PathSeparator) + JLinkCmd
	info, err := os.Stat(jLink)
	jLinkExists := !os.IsNotExist(err) && info.Mode().IsRegular()
	return jLinkExists
}
