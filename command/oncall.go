// deputize - Update an LDAP group with info from the PagerDuty API
// oncall.go: struct for oncall command
//
// Copyright 2017-2022 F5 Inc.
// Licensed under the BSD 3-clause license; see LICENSE for more information.

package command

import (
	"io"
	"log"
	"os"
	"strings"

	"github.com/Graylog2/go-gelf/gelf"
	"github.com/threatstack/deputize/config"
	"github.com/threatstack/deputize/oncall"
)

// OncallCommand gets the data from the CLI
type OncallCommand struct {
	Meta
}

// Run actually runs oncall/oncall.go
func (c *OncallCommand) Run(args []string) int {
	if _, err := os.Stat(config.ConfigFile); os.IsNotExist(err) {
		log.Fatalf("No config file present (or invalid format). See README.md.\n")
	}
	var conf = config.Config
	if conf.GrayLogEnabled {
		if conf.GrayLogAddress == "" {
			log.Fatalf("GrayLogEnabled is true, and no graylog address was specified\n")
		}
		gelfWriter, err := gelf.NewWriter(conf.GrayLogAddress)
		if err != nil {
			log.Fatalf("gelf.NewWriter: %s", err)
		}
		log.SetOutput(io.MultiWriter(os.Stdout, gelfWriter))
	}
	err := oncall.UpdateOnCallRotation(conf)
	if err != nil {
		log.Fatalf("Oncall update failed: %s", err)
	}
	return 0
}

// Synopsis gives the help output for oncall
func (c *OncallCommand) Synopsis() string {
	return "Gets oncall schedule from PagerDuty and updates LDAP"
}

// Help prints more useful info for oncall
func (c *OncallCommand) Help() string {
	helpText := `
Usage: deputize oncall

  Pull current oncall schedule from PagerDuty and update LDAP.

  This command will connect to PagerDuty using an API key and pull the
  email addresses of the people who are on call. It'll connect to LDAP
  and pull the members of the lg-oncall group. If there's a difference,
  it will replace the lg-oncall group with the members from PagerDuty.
`
	return strings.TrimSpace(helpText)
}
