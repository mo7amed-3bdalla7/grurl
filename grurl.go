package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	config := RemoteConfig{data: map[string]string{}}

	config.ParseFile("./.git/config")

	fmt.Print(config.data)
}

type RemoteConfig struct {
	configPath string
	data       map[string]string
}

func (c *RemoteConfig) ParseFile(path string) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}

	defer file.Close()
	c.configPath = path

	c.Parse(file)

	return nil
}

func (c *RemoteConfig) Parse(file *os.File) {
	const remoteRowStart = `[remote "`
	const remoteUrlRowStart = `url = `
	var remoteName string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.Index(line, remoteRowStart) != -1 {
			remoteName = line[len(remoteRowStart) : len(line)-2]

			c.data[remoteName] = ""
		}

		if strings.Index(line, remoteUrlRowStart) != -1 && remoteName != "" {
			url := line[len(remoteUrlRowStart):]
			url = SshUrlToHttpUrl(url)

			c.data[remoteName] = url
			remoteName = ""
		}
	}
}

//Checks if a url is a valid ssh url and converts it to http url or returns the same url
func SshUrlToHttpUrl(url string) string {
	regex := regexp.MustCompile(
		`^(?P<user>.+)@(?P<host>.+?):(?P<port>[0-9]*)(?P<path>.+)$`,
	)

	groups := regex.FindStringSubmatch(url)

	if len(groups) == 5 {
		url = ""

		url = fmt.Sprintf("https://%s", groups[2])

		if groups[3] != "" {
			url += fmt.Sprintf(":%s", groups[3])
		}

		if strings.Index(groups[4], "/") == 0 {
			url += fmt.Sprintf("%s", groups[4])
		} else {
			url += fmt.Sprintf("/%s", groups[4])
		}

		return url
	}

	return url
}
