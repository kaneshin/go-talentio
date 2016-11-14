package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/k0kubun/pp"
	"github.com/kaneshin/go-talentio/talentio"
	"github.com/kaneshin/go-talentio/talentio/talentioutil"
)

const (
	success = iota
	failure
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

	if *page < 0 {
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

	cands, err := talentioutil.Candidates.List(client, &talentioutil.CandidatesListOptions{
		MaxPage: 10,
		Status:  talentio.StatusOngoing,
		Sort:    talentio.SortRegisteredAtDescKey,
	})

	if err != nil {
		log.Println(err)
		return failure
	}

	for _, candidate := range cands {
		fmt.Printf("ID=%d, RegisteredAt=%v\n", candidate.ID, candidate.RegisteredAt)
	}

	if len(cands) == 0 {
		return 0
	}

	// Show details of candidates.
	if err = talentioutil.Candidates.ApplyDetails(client, cands); err != nil {
		log.Println(err)
		return failure
	}

	for _, cand := range cands {
		pp.Println(cand)
	}

	return 0
}

func main() {
	os.Exit(run())
}
