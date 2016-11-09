package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/kaneshin/go-talentio/talentio"
)

var (
	page       = flag.Int("page", 1, "")
	status     = flag.String("status", "ongoing", "")
	configPath = flag.String("config", "$HOME/.talentio.tml", "")
)

// config represents a configuration of commands.
var config = struct {
	AccessToken string `toml:"access_token"`
}{}

func flagParse() {
	flag.Parse()

	if *page < 1 {
		*page = 1
	}

	statuses := [6]string{
		talentio.StatusOngoing,
		talentio.StatusReject,
		talentio.StatusFail,
		talentio.StatusPass,
		talentio.StatusPoolActive,
		talentio.StatusPoolInactive,
	}
	for _, s := range statuses {
		if *status == s {
			goto STATUS_OK
		}
	}
	*status = talentio.StatusOngoing
STATUS_OK:

	fp := os.ExpandEnv(*configPath)
	_, _ = toml.DecodeFile(fp, &config)
}

func run() int {
	flagParse()

	c := talentio.NewConfig().WithHTTPClient(http.DefaultClient)
	str := config.AccessToken
	if str == "" {
		str = os.Getenv("TALENTIO_ACCESS_TOKEN")
	}
	c.WithAccessToken(str)

	client := talentio.NewClient(c)

	opt := talentio.CandidatesListOptions{
		Page:   *page,
		Status: *status,
		Sort:   talentio.SortRegisteredAtDescKey,
	}
	candidates, resp, err := client.Candidates.List(&opt)
	if err != nil {
		log.Printf("error: %v", err.Error())
		return 1
	}

	for _, candidate := range candidates {

		fmt.Printf("ID=%d: %v %v (%v)\n", candidate.ID, candidate.FirstName, candidate.LastName, candidate.Status)
	}
	fmt.Printf("Total=%d, Remaining=%d Reset=%d\n", resp.Total, resp.Remaining, resp.Reset)

	return 0
}

func main() {
	os.Exit(run())
}
