package generator

import (
	"os/exec"
)

func InitGoModule(projectName, moduleName string) error {

	cmd := exec.Command("go", "mod", "init", moduleName)

	cmd.Dir = projectName // masuk ke folder project

	return cmd.Run()
}