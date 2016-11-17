# Talentio Library for Golang

## Installation

```shell
go get -d github.com/kaneshin/go-talentio
```

## Usage

### Go code

```go
import (
	"net/http"
	"os"

	"github.com/kaneshin/go-talentio/talentio"
	"github.com/kaneshin/go-talentio/talentio/talentioutil"
)

func main() {

	client := talentio.NewClient(
		talentio.NewConfig().
			WithHTTPClient(http.DefaultClient).
			WithAccessToken(os.Getenv("TALENTIO_ACCESS_TOKEN"))
		)
	}

	// fetch ongoing candidates using talentio/talentioutil.
	candidates, err := talentioutil.Candidates.List(client, &talentioutil.CandidatesListOptions{
		MaxPage: 10,                               // limit of paging
		Status:  talentio.StatusOngoing,           // status=ongoing
		Sort:    talentio.SortRegisteredAtDescKey, // sort=-registeredAt
	})
	if err != nil {
		panic(err)
	}

	for _, candidate := range candidates {
		fmt.Printf("ID=%d, RegisteredAt=%v\n", candidate.ID, candidate.RegisteredAt)
	}

	if len(candidates) == 0 {
		return 0
	}

	// fetch candidate's details using talentio/talentioutil.
	if err = talentioutil.Candidates.ApplyDetails(client, candidates); err != nil {
		panic(err)
	}

	// ...
}
```

## License

[The MIT License (MIT)](http://kaneshin.mit-license.org/)

## Author

Shintaro Kaneko <kaneshin0120@gmail.com>
