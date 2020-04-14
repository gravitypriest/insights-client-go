package main

import (
	"os/exec"
	"fmt"
	"io/ioutil"
	"strings"
)

func readFile(fname string) string {
	data, _ := ioutil.ReadFile(fname)
	return string(data)
}

func biosId() string {
	output, _ := exec.Command("/usr/sbin/dmidecode -s system-uuid").Output()
	return string(output)
}

func subManId() string {
	output, _ := exec.Command("/usr/sbin/subscription-manager identity").Output()
	firstLine := strings.Split(string(output), "\n")[0]
	subManId := strings.Split(firstLine, ":")[1]
	return subManId
}

func getCanonicalFacts() map[string]string {
	facts := map[string]string{
		"insights_id": readFile("/etc/insights-client/machine-id"),
		"machine_id": readFile("/etc/machine-id"),
		"bios_uuid": biosId(),
		"subscription_manager_id": subManId(),
	}
	return facts
}
