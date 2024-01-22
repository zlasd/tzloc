package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	tzDataPath := GetTzDataPath()
	tzList := LoadTzList(tzDataPath)
	for _, tz := range tzList {
		fmt.Println(GetIdentifierName(tz), "=", tz)
	}
}

func GetTzDataPath() string {
	out, err := exec.Command("go", "env", "GOROOT").Output()
	if err != nil {
		log.Fatalln(err)
	}
	return strings.Trim(string(out), "\n") + "/lib/time/zoneinfo.zip"
}

func LoadTzList(path string) []string {
	fmt.Println("Exec zipinfo command: zipinfo -1", path)
	out, err := exec.Command("zipinfo", "-1", path).Output()
	if err != nil {
		log.Fatalln(err)
	}
	return strings.Split(strings.Trim(string(out), "\n"), "\n")
}

func GetIdentifierName(tzVal string) string {
	replacer := strings.NewReplacer(
		"/", "_",
		"GMT-", "GMT_Minus_",
		"-", "_",
		"+", "_Plus_",
	)
	return replacer.Replace(tzVal)
}
