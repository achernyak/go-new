package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s directory", os.Args[0])
		os.Exit(1)
	}
	project := os.Args[1]
	dir := getDir(project)

	setupProject(dir, project)
}

func getDir(project string) string {
	wd, err := os.Getwd()
	checkError(err)

	gopath := os.Getenv("GOPATH")
	if strings.Contains(wd, gopath) {
		return project
	}

	pathArgs := []string{
		gopath,
		"src/github.com",
		getUser(project),
		project,
	}
	return strings.Join(pathArgs, "/")
}

func getUser(project string) string {
	home := os.Getenv("HOME")
	path := home + "/.gitconfig"

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("No confi file found. Creating locally")
		return ""
	}

	userBlock := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if userBlock {
			if strings.HasPrefix(text, "[") {
				break
			}

			if strings.Contains(text, "name =") {
				user := strings.Replace(text, "name = ", "", 1)
				return strings.TrimSpace(user)
			}
		}

		if text == "[user]" {
			userBlock = true
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading config. Creating locally")
		return ""
	}
	return ""
}

func setupProject(dir, project string) {
	var err error
	err = os.MkdirAll(dir, 0777)
	checkError(err)

	filename := dir + "/" + project
	file := filename + "." + "go"
	_, err = os.Create(file)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
