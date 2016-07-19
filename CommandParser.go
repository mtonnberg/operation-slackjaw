package main

import (
	"strings"
)

func parseCommand(command string) (ok bool, action string, appName string, env string) {
	deployArgs := strings.Split(command, " ")
	if len(deployArgs) == 3 {
		return true, deployArgs[0], deployArgs[1], deployArgs[2]
	}
	return false, "", "", ""
}
