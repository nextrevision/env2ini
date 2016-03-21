package main

import (
	"os"
	"testing"

	"gopkg.in/ini.v1"
)

func TestMain(t *testing.T) {
	os.Setenv("TEST__section_1__key_1", "Value_1")
	os.Setenv("TEST__DEFAULT__Key_2", "Value_2")
	config := Config{
		Filename:  "test.ini",
		EnvPrefix: "TEST",
	}
	parseEnvironment(config)

	cfg, err := ini.Load("test.ini")
	if err != nil {
		t.Error(err)
	}

	if cfg.Section("section_1").Key("key_1").String() != "Value_1" {
		t.Error("Values do not match")
	}
	if cfg.Section("DEFAULT").Key("Key_2").String() != "Value_2" {
		t.Error("Values do not match")
	}
	os.Remove("test.ini")
}

func TestReplaceReserved(t *testing.T) {
	if replaceSpecialChars("TEST_DOT") != "TEST." {
		t.Error("Does not convert: TEST_DOT to TEST.")
	}
	if replaceSpecialChars("TEST_DOT_") != "TEST." {
		t.Error("Does not convert: TEST_DOT_ to TEST.")
	}
	if replaceSpecialChars("DOT_TEST") != ".TEST" {
		t.Error("Does not convert: DOT_TEST to .TEST")
	}
	if replaceSpecialChars("TEST_SLASH") != "TEST/" {
		t.Error("Does not convert: TEST_SLASH to TEST/")
	}
	if replaceSpecialChars("TEST_SLASH_") != "TEST/" {
		t.Error("Does not convert: TEST_SLASH_ to TEST/")
	}
	if replaceSpecialChars("SLASH_TEST") != "/TEST" {
		t.Error("Does not convert: SLASH_TEST to /TEST")
	}
	if replaceSpecialChars("TEST_COLON") != "TEST:" {
		t.Error("Does not convert: TEST_COLON to TEST:")
	}
	if replaceSpecialChars("TEST_COLON_") != "TEST:" {
		t.Error("Does not convert: TEST_COLON_ to TEST:")
	}
	if replaceSpecialChars("COLON_TEST") != ":TEST" {
		t.Error("Does not convert: COLON_TEST to :TEST")
	}
}
