package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/kaneshin/go-talentio/cmd/internal"
	"github.com/kaneshin/go-talentio/talentio"
)

var (
	page   = flag.Int("page", 1, "")
	status = flag.String("status", "ongoing", "")
)

func validate() {
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
			goto done
		}
	}
	*status = talentio.StatusOngoing

done:
}

func main() {
	internal.ParseFlag()
	validate()

	c := talentio.NewConfig().WithHTTPClient(http.DefaultClient)
	str := internal.Config.AccessToken
	if str == "" {
		str = os.Getenv("TALENTIO_ACCESS_TOKEN")
	}
	c.WithAccessToken(str)

	client := talentio.NewClient(c)
	opt := talentio.CandidatesListOptions{
		Page: *page,
		// FIXME:
		// It's not working. API might be broken.
		// Status: *status,
		Sort: talentio.SortRegisteredAtDescKey,
	}

	candidates, _, err := client.Candidates.List(&opt)
	if err != nil {
		panic(err)
	}

	for _, candidate := range candidates {
		if *status != candidate.Status {
			// FIXME:
			// The because of opt.Status is not working.
			continue
		}

		fmt.Printf("ID=%d: %v %v\n", candidate.ID, candidate.FirstName, candidate.LastName)
	}
}
