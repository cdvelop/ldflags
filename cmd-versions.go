package ldflags

import (
	"log"
	"os/exec"
)

func GoVersion() (version string) {
	return execCommand("go", "version")
}

func TinyGoVersion() (version string) {
	return execCommand("tinygo", "version")
}

func execCommand(command string, args ...string) (result string) {
	cmd := exec.Command(command, args...)
	output, er := cmd.CombinedOutput()
	if er != nil {
		log.Println("execCommand ", command, " error "+er.Error())
		return
	}
	return string(output)

}
