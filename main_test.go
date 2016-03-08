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

func TestSanitize(t *testing.T) {
	if sanitize("TEST_DOT") != "TEST." {
		t.Error("Does not convert: TEST_DOT to TEST.")
	}
	if sanitize("TEST_DOT_") != "TEST." {
		t.Error("Does not convert: TEST_DOT_ to TEST.")
	}
	if sanitize("DOT_TEST") != ".TEST" {
		t.Error("Does not convert: DOT_TEST to .TEST")
	}
	if sanitize("TEST_SLASH") != "TEST/" {
		t.Error("Does not convert: TEST_SLASH to TEST/")
	}
	if sanitize("TEST_SLASH_") != "TEST/" {
		t.Error("Does not convert: TEST_SLASH_ to TEST/")
	}
	if sanitize("SLASH_TEST") != "/TEST" {
		t.Error("Does not convert: SLASH_TEST to /TEST")
	}
	if sanitize("TEST_COLON") != "TEST:" {
		t.Error("Does not convert: TEST_COLON to TEST:")
	}
	if sanitize("TEST_COLON_") != "TEST:" {
		t.Error("Does not convert: TEST_COLON_ to TEST:")
	}
	if sanitize("COLON_TEST") != ":TEST" {
		t.Error("Does not convert: COLON_TEST to :TEST")
	}
}
