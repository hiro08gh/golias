package external

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

// Target shell file
const (
	BASH_FILE         = ".bashrc"
	BASH_PROFILE_FILE = ".bashrc_profile"
	ZSH_FILE          = ".zshrc"
)

// Using shell
const (
	BASH_PATH = "/usr/bin/bash"
	ZSH_PATH  = "/usr/bin/zsh"
)

const (
	WRITE_PERM = 0666
	READ_PERM  = 0600
)

// Check Exists file
func Exists(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

func getRootDir() string {
	env := "HOME"
	if runtime.GOOS == "windows" {
		env = "USERPROFILE"
	} else if runtime.GOOS == "plan9" {
		env = "home"
	}
	return os.Getenv(env)
}

// Create shell path ex: /usr/name/.bashrc
func createPath(currentDir, shell string) string {
	t := fmt.Sprintf("%s/%s", currentDir, shell)
	return t
}

// Check using shell and create path
func getShellPath() string {
	shell := os.Getenv("SHELL")

	// Get Home Directory path
	currentDir := getRootDir()

	// Default path
	path := createPath(currentDir, BASH_FILE)

	switch shell {
	case BASH_PATH:
		path := createPath(currentDir, BASH_FILE)
		return path
	case ZSH_PATH:
		path := createPath(currentDir, ZSH_FILE)
		return path
	}
	return path
}

// Add set alias
func Add(args []string) {
	path := getShellPath()
	fmt.Println(path)

	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, READ_PERM)

	if err != nil {
		log.Fatal(err)
	}

	defer fp.Close()

	content := fmt.Sprintf("alias %s=%s", args[0], args[1])

	fmt.Fprintln(fp, content)

	fmt.Println(args[0])
}
