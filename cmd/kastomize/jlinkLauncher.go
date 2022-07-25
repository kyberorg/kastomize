package main

import (
	"bytes"
	"fmt"
	"github.com/kyberorg/kastomize/cmd/kastomize/config"
	"github.com/pieterclaerhout/go-log"
	"io"
	"os"
	"os/exec"
)

func MakeJLinkCommand() *exec.Cmd {
	jlinkCommand := config.Args().JavaHome + string(os.PathSeparator) +
		config.BinDir + string(os.PathSeparator) + config.JLinkCmd

	jLinkArgs := make([]string, 0)
	for _, module := range validModules {
		jLinkArgs = append(jLinkArgs, "--add-modules")
		jLinkArgs = append(jLinkArgs, module)
	}
	jLinkArgs = append(jLinkArgs, "---strip-java-debug-attributes")
	jLinkArgs = append(jLinkArgs, "--no-man-pages")
	jLinkArgs = append(jLinkArgs, "--no-header-files")
	jLinkArgs = append(jLinkArgs, "--compress=2")
	jLinkArgs = append(jLinkArgs, "--output")
	jLinkArgs = append(jLinkArgs, config.Args().OutputDir)

	fmt.Printf("lauching %s %s", jlinkCommand, jLinkArgs)
	cmd := exec.Command(jlinkCommand, jLinkArgs...)
	return cmd
}

func ExecuteJLink(cmd *exec.Cmd) {
	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	cmd.Stdout = mw
	cmd.Stderr = mw

	err := cmd.Run()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to start err=%v\n", err)
		fmt.Println(stdBuffer.String())
		os.Exit(1)
	} else {
		log.Info("Ready")
		log.Infof("You can find your custom JVM at %s", config.Args().OutputDir)
	}
}
