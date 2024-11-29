package exec

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// RunCMD runs a command in the shell
//
// Parameters:
//   - fs: the command to run
//
// Returns:
//   - error: an error if the command fails
func RunCMD(fs string) error {
	cmd := exec.Command("bash", "-c", fs)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running command:", err)
	}

	return nil
}

// RunReturnCMD runs a command in the shell and returns the output
//
// Parameters:
//   - fs: the command to run
//
// Returns:
//   - string: the output of the command
//   - error: an error if the command fails
func RunReturnCMD(fs string) (string, error) {
	cmd := exec.Command("bash", "-c", fs)
	if errors.Is(cmd.Err, exec.ErrDot) {
		cmd.Err = nil
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

// RunLooseCMD runs a command in the shell without returning anything
//
// Parameters:
//   - fs: the command to run
func RunLooseCMD(fs string) {
	cmd := exec.Command("bash", "-c", fs)
	_ = cmd.Run()
	return
}

// RunReallyLooseCMD runs a command in the shell without returning anything
// and without waiting for the command to finish
//
// Parameters:
//   - fs: the command to run
func RunReallyLooseCMD(fs string) {
	cmd := exec.Command("bash", "-c", fs)
	go cmd.Run()
	return
}
