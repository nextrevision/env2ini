package main

import (
	"os"
	"testing"

	"gopkg.in/ini.v1"
)

func TestMain(t *testing.T) {
	os.Create("test.ini")
	Configs = []Config{
		Config{
			EnvPrefix: "TEST",
			Filename:  "test.ini",
			Defaults:  "DEFAULTS",
		},
	}

	os.Setenv("TEST__SECTION_1__KEY_1", "value1")
	os.Setenv("TEST__DEFAULTS__DEFAULTSKEY", "DEFAULTSvalue")
	main()
	cfg, err := ini.Load("test.ini")
	if err != nil {
		t.Error(err)
	}

	if cfg.Section("section_1").Key("key_1").String() != "value1" {
		t.Error("Values do not match")
	}
	if cfg.Section("DEFAULTS").Key("defaultskey").String() != "DEFAULTSvalue" {
		t.Error("Values do not match")
	}
	os.Remove("test.ini")
}
