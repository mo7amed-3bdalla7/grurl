package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	listOption := flag.Bool("list", false, "list all remote names with URLs and exit")
	remoteOption := flag.String("remote", "origin", "to set the remote name you for required url")
	pathOption := flag.String("path", ".", "to set the git repository path")

	flag.Parse()

	path := *pathOption + "/.git/config"
	config := NewRemoteConfig()
	err := config.ParseFile(path)

	if err != nil {
		fmt.Println("Check path value.")
		os.Exit(1)
	}

	if *remoteOption == "" {
		fmt.Println("Remote value is required.")
		os.Exit(1)
	}

	if *listOption {
		for key, value := range config.data {
			fmt.Printf("%s \t %s\n", key, value)
		}
		return
	}

	if *remoteOption != "" {
		if url, err := config.data[*remoteOption]; err {
			fmt.Println(url)
			return
		}
		fmt.Println("This remote isn't exist.")
		os.Exit(1)
	}

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

func NewRemoteConfig() *RemoteConfig {
	return &RemoteConfig{data: map[string]string{}}
}
