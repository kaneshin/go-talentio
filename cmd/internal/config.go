package internal

import (
	"flag"
	"os"

	"github.com/BurntSushi/toml"
)

var (
	configPath = flag.String("config", "$HOME/.talentio.tml", "")

	// Config represents a configuration of commands.
	Config = struct {
		AccessToken string `toml:"access_token"`
	}{}
)

// ParseFlag parses flag options and toml file.
func ParseFlag() {
	flag.Parse()

	fp := os.ExpandEnv(*configPath)
	if _, err := toml.DecodeFile(fp, &Config); err != nil {
		return
	}
}
