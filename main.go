package main

import (
	"flag"
	"os"
	"regexp"
	"strings"

	log "github.com/Sirupsen/logrus"
)

func init() {
	debug := flag.Bool("debug", false, "enable debug logging")
	flag.Parse()

	// enable debug log level when debug flag is set
	if *debug {
		log.SetLevel(log.DebugLevel)
	}
}

func main() {
	for _, envVar := range os.Environ() {
		envKey := strings.Split(envVar, "=")[0]
		envValue := strings.Split(envVar, "=")[1]
		matched, err := regexp.MatchString(`^([A-Za-z0-9_]+__){2}[A-Za-z0-9_]+$`, envKey)
		if err != nil {
			log.Fatal(err.Error())
		}

		if matched {
			for _, config := range Configs {

				match, ok := config.mapMatch(envKey)
				if !ok {
					log.WithFields(log.Fields{
						"key": envKey,
					}).Error("Could not map key")
					continue
				}

				if match["Prefix"] == config.EnvPrefix {
					section := sanitize(match["Section"])
					key := strings.ToLower(sanitize(match["Key"]))

					if match["Section"] != config.Defaults {
						section = strings.ToLower(section)
					}

					if err := config.updateSetting(section, key, envValue); err != nil {
						log.Fatal(err.Error())
					}

					log.WithFields(log.Fields{
						"section": section,
						"key":     key,
					}).Info("Updated setting")

					break
				}
			}
			log.Debug("Skipping key: ", envKey)
		} else {
			log.Debug("Skipping key: ", envKey)
		}
	}
}

// sanitize replaces reserved words in variables to their symbol counterpart
func sanitize(value string) string {
	dotRegex := regexp.MustCompile("_?DOT_?")
	slashRegex := regexp.MustCompile("_?SLASH_?")
	colonRegex := regexp.MustCompile("_?COLON_?")
	value = dotRegex.ReplaceAllString(value, ".")
	value = slashRegex.ReplaceAllString(value, "/")
	value = colonRegex.ReplaceAllString(value, ":")
	return value
}
