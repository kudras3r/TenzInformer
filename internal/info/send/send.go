package send

import "os/exec"

func JSON(filePath string) error {
	// tenzir may be not in PATH - reason why we run /opt...
	// pipeline definition: from...
	pipeDesc := "/opt/tenzir/bin/tenzir 'from file " + filePath + " read json | import'"

	cmd := exec.Command("sudo", "sh", "-c", pipeDesc)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
