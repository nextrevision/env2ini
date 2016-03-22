package main

import (
	"regexp"

	"github.com/go-ini/ini"
)

// Config : struct for defining OpenStack backed configs
type Config struct {
	EnvPrefix string
	Filename  string
}

// matchRegex is a named regex that matches an env key in the following format:
// PREFIX__SECTION__KEY
var matchRegex = regexp.MustCompile(`^(?P<Prefix>[A-Za-z0-9_]+)__(?P<Section>[A-Za-z0-9_]+)__(?P<Key>[A-Za-z0-9_]+)$`)

// mapMatch converts and returns the named groups as a map of strings
func (c *Config) mapMatch(key string) (map[string]string, bool) {
	setting := make(map[string]string)

	match := matchRegex.FindStringSubmatch(key)
	if match == nil {
		return setting, false
	}

	for i, name := range matchRegex.SubexpNames() {
		if i == 0 || name == "" {
			continue
		}
		setting[name] = match[i]
	}
	return setting, true
}

// updateSetting writes the desired setting to the destination file
func (c *Config) updateSetting(section string, key string, value string) error {
	ini.DefaultHeader = true
	cfg, err := ini.Load(c.Filename)
	if err != nil {
		return err
	}
	setting := cfg.Section(section).Key(key)
	setting.SetValue(value)
	return cfg.SaveTo(c.Filename)
}
