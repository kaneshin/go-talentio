package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/k0kubun/pp"
	"github.com/kaneshin/go-talentio/cmd/internal"
	"github.com/kaneshin/go-talentio/talentio"
)

var (
	color = flag.String("color", "", "")
	image = flag.String("image", "", "")

	channel  = flag.String("channel", "", "")
	username = flag.String("username", "", "")
	emoji    = flag.String("emoji", "", "")

	code = flag.Bool("code", false, "")
)

func main() {
	internal.ParseFlag()

	c := talentio.NewConfig().WithHTTPClient(http.DefaultClient)

	{
		var str string
		if str = internal.Config.AccessToken; str == "" {
			str = os.Getenv("TALENTIO_ACCESS_TOKEN")
		}
		c.WithAccessToken(str)
	}

	client := talentio.NewClient(c)

	opt := talentio.CandidatesListOptions{
		Page:   1,
		Status: talentio.StatusOngoing,
		Sort:   talentio.RegisteredAtDescKey,
	}
	v, _, err := client.Candidates.List(&opt)
	if err != nil {
		panic(err)
	}
	pp.Println(v)
}
