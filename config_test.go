package main

import (
	"testing"
)

func TestNodeInfo(t *testing.T) {
	nodeInfoTest := NodeInfo{systemID: "system_id"}
	if nodeInfoTest.systemID != "system_id" {
		t.Fail()
	}
}

func TestConfigStructure(t *testing.T) {
	configStructure := Config{APIKey: "api_key", APIURL: "api_url", APIver: "1.0"}
	if _, err := configStructure.Client(); err == nil {
		t.Fail()
	}
}
