package main

import (
	"flag"
	"os"
	"regexp"
	"strings"

	log "github.com/Sirupsen/logrus"
)

func main() {
	debug := flag.Bool("debug", false, "enable debug logging")
	filename := flag.String("filename", "", "destination filename for writing settings (required)")
	envPrefix := flag.String("prefix", "", "environment prefix to look for keys (required)")

	flag.Parse()

	// enable debug log level when debug flag is set
	if *debug {
		log.SetLevel(log.DebugLevel)
	}

	if *filename == "" || *envPrefix == "" {
		log.Fatal("Must provide both -filename and -prefix")
	}

	config := Config{
		Filename:  *filename,
		EnvPrefix: *envPrefix,
	}
	parseEnvironment(config)
}

func parseEnvironment(config Config) {
	for _, envVar := range os.Environ() {
		envKey := strings.Split(envVar, "=")[0]
		envValue := strings.Split(envVar, "=")[1]
		matched, err := regexp.MatchString(`^([A-Za-z0-9_]+__){2}[A-Za-z0-9_]+$`, envKey)
		if err != nil {
			log.Fatal(err.Error())
		}

		if matched {
			match, ok := config.mapMatch(envKey)
			if !ok {
				log.WithFields(log.Fields{
					"key": envKey,
				}).Error("Could not map key")
				continue
			}

			if match["Prefix"] == config.EnvPrefix {
				section := replaceSpecialChars(match["Section"])
				key := replaceSpecialChars(match["Key"])

				if _, err := os.Stat(config.Filename); os.IsNotExist(err) {
					if _, err := os.Create(config.Filename); err != nil {
						log.Fatal("Could not create file ", config.Filename)
					} else {
						log.Info("Created file ", config.Filename)
					}
				}

				if err := config.updateSetting(section, key, envValue); err != nil {
					log.Fatal(err.Error())
				}

				log.WithFields(log.Fields{
					"section": section,
					"key":     key,
				}).Info("Updated setting")

				continue
			} else {
				log.Debug("Skipping key: ", envKey)
			}
		} else {
			log.Debug("Skipping key: ", envKey)
		}
	}
}

// replaceSpecialChars replaces reserved words in variables to their symbol counterpart
func replaceSpecialChars(value string) string {
	dotRegex := regexp.MustCompile("_?DOT_?")
	slashRegex := regexp.MustCompile("_?SLASH_?")
	colonRegex := regexp.MustCompile("_?COLON_?")
	value = dotRegex.ReplaceAllString(value, ".")
	value = slashRegex.ReplaceAllString(value, "/")
	value = colonRegex.ReplaceAllString(value, ":")
	return value
}
