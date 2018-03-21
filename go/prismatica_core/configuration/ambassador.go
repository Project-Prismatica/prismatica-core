package configuration

import (
	"bytes"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func HandleAmbassadorConfiguration(sourceDir string, destinationDir string) {

	var stdout, stderr bytes.Buffer
	cmd := exec.Command("cp", "-rv", sourceDir, destinationDir)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	log.WithFields(log.Fields{"command": cmd.Path, "args": cmd.Args,
		"configuration_source": sourceDir,
		"configuration_dest": destinationDir}).
		Info("copying ambassador configuration")

	err := cmd.Run()
	if err != nil {
		log.WithFields(log.Fields{"error": err, "stderr": stderr.String(),
			"stdout": stdout.String()}).
			Error("could not copy configuration")
			return
	}

	// NOTE: This doesn't check the return value since it should be replaced

	log.WithFields(log.Fields{"stderr": stderr.String(),
		"stdout": stdout.String()}).
		Info("copied ambassador configuration")

	return
}
