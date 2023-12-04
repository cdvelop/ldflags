package ldflags

import "os/exec"

func GoVersion() (version, err string) {
	cmd := exec.Command("go", "version")
	output, er := cmd.CombinedOutput()
	if er != nil {
		return "", "GoVersion error " + er.Error()
	}
	return string(output), ""
}
