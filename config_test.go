package main

import (
	"os"
	"testing"

	"gopkg.in/ini.v1"
)

func TestMapMatch(t *testing.T) {
	c := Config{
		EnvPrefix: "TEST",
		Filename:  "test.ini",
		Defaults:  "DEFAULTS",
	}

	match, ok := c.mapMatch("TEST__SECTION_1__KEY_1")
	if !ok {
		t.Error("Could not create matched map for key")
	}

	if match["Prefix"] != "TEST" {
		t.Error("Prefix does not match")
	}
	if match["Section"] != "SECTION_1" {
		t.Error("Section does not match")
	}
	if match["Key"] != "KEY_1" {
		t.Error("Key does not match")
	}
}

func TestUpdateSetting(t *testing.T) {
	os.Create("test.ini")
	c := Config{
		EnvPrefix: "TEST",
		Filename:  "test.ini",
		Defaults:  "DEFAULTS",
	}

	if err := c.updateSetting("section1", "key1", "value1"); err != nil {
		t.Error(err)
	}

	cfg, err := ini.Load("test.ini")
	if err != nil {
		t.Error(err)
	}

	if cfg.Section("section1").Key("key1").String() != "value1" {
		t.Error("Values do not match")
	}

	os.Remove("test.ini")
}
