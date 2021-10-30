package utils

import (
	"flag"
	"fmt"
	"go-count-db-stuff/internal/constants"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func SanitizeCmdArgs(args []string) []string {
	var parsed []string
	for _, a := range args {
		if !strings.HasPrefix(a, "-") {
			parsed = append(parsed, a)
			continue
		}

		parts := strings.Split(a, "=")
		if len(parts) == 1 {
			parsed = append(parsed, a)
			continue
		}

		cmd := strings.ToLower(parts[0])
		value := strings.Join(parts[1:], "=")
		parsed = append(parsed, fmt.Sprintf("%s=%s", cmd, value))

	}

	return parsed
}

func ParseCommandLine() string {
	generalCmd := flag.NewFlagSet("general", flag.ExitOnError)
	readCmdPtr := generalCmd.String("config", constants.DefaultConfigFile, "Which Config file to use.")
	var err error
	var configFile string

	err = generalCmd.Parse(SanitizeCmdArgs(os.Args[1:]))
	ErrorHandler(err, "Failed to parse command line")
	configFile = *readCmdPtr

	fullConfigFilePath, err := filepath.Abs(configFile)

	if fullConfigFilePath != configFile {
		log.Printf("File '%s' was relative and was converted to absolute path: '%s'\n", configFile, fullConfigFilePath)
	}

	return fullConfigFilePath
}
