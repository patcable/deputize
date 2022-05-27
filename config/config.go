// deputize - Update an LDAP group with info from the PagerDuty API
// config.go: configurator for the deputize tool.
//
// Copyright 2017-2022 F5 Inc.
// Licensed under the BSD 3-clause license; see LICENSE for more information.

package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var Config DeputizeConfig
var ConfigFile string

func init() {
	if os.Getenv("DEPUTIZE_CONFIG") == "" {
		ConfigFile = "config.json"
	} else {
		ConfigFile = os.Getenv("DEPUTIZE_CONFIG")
	}
	if _, err := os.Stat(ConfigFile); err == nil {
		Config = NewConfig(ConfigFile)
	}
}

// DeputizeConfig is our config struct
type DeputizeConfig struct {
	BaseDN                 string
	GrayLogEnabled         bool
	GrayLogAddress         string
	LDAPServer             string
	LDAPPort               int
	MailAttribute          string
	MemberAttribute        string
	ModUserDN              string
	OnCallGroup            string
	OnCallGroupDN          string
	OnCallSchedules        []string
	RootCAFile             string
	RunDuration            string
	SlackChan              []string
	SlackEnabled           bool
	TokenPath              string
	UserAttribute          string
	VaultSecretPath        string
	VaultServer            string
	Quiet                  bool
	GitlabEnabled          bool
	GitlabServer           string
	GitlabToken            string
	GitlabGroup            string
	GitlabApproverSchedule string
}

func NewConfig(fname string) DeputizeConfig {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	config := DeputizeConfig{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	return config
}
