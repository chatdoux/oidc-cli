package config

import (
	"fmt"
	"os/exec"
	"strings"
)

const keychainServiceName = "oidc"

func setKeychainPassword(name string, password string) error {
	cmd := exec.Command("security", "add-generic-password", "-s", keychainServiceName, "-a", name, "-w", password, "-U")
	_, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("exec: %v", err)
	}
	return nil
}

func getKeychainPassword(name string) (string, error) {
	cmd := exec.Command("security", "find-generic-password", "-s", keychainServiceName, "-a", name, "-w")
	stdout, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("exec: %v", err)
	}
	return strings.TrimSpace(string(stdout)), nil
}
