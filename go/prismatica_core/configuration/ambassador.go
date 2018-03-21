package configuration

import (
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func HandleAmbassadorConfiguration(sourceDir string, destinationDir string) {

	cmd := exec.Command("cp", "-rv", sourceDir + "/",
		destinationDir + "/")

	log.WithFields(log.Fields{"command": cmd.Path, "args": cmd.Args,
		"configuration_source": sourceDir,
		"configuration_dest": destinationDir}).
		Info("copying ambassador configuration")

	err := cmd.Start()
	if err != nil {
		output, err := cmd.Output()
		log.WithFields(log.Fields{"error": err, "output": string(output)}).
			Error("could not copy configuration")
	} else {
		log.Info("copied ambassador configuration")
	}


}
