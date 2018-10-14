package talentioutil

import (
	"sync"

	"github.com/kaneshin/go-talentio/talentio"
)

type (
	_Candidates struct{}

	// CandidatesListOptions represents list of candidates.
	CandidatesListOptions struct {
		Status  string
		Sort    string
		MaxPage int
	}
)

// Candidates is.
var Candidates _Candidates

func (c _Candidates) List(client *talentio.Client, opt *CandidatesListOptions) ([]*talentio.Candidate, error) {
	if opt == nil {
		opt = &CandidatesListOptions{}
	}
	if opt.MaxPage == 0 {
		// default value of max page.
		opt.MaxPage = 10
	}

	var candidates []*talentio.Candidate
	for page := 1; page <= opt.MaxPage; page++ {
		cands, resp, err := client.Candidates.List(
			&talentio.CandidatesListOptions{
				Page:   page,
				Status: opt.Status,
				Sort:   opt.Sort,
			},
		)
		if err != nil {
			return nil, err
		}
		candidates = append(candidates, cands...)

		if resp.Total <= len(candidates) {
			break
		}
	}

	return candidates, nil
}

func (c _Candidates) ApplyDetails(client *talentio.Client, list []*talentio.Candidate) error {

	var (
		e  error
		wg sync.WaitGroup
	)
	tmp := list
	for i, c := range tmp {
		i := i
		c := c
		wg.Add(1)
		go func(i int, c *talentio.Candidate) {
			defer wg.Done()

			var err error
			c, _, err = client.Candidates.Get(c.ID)
			if err != nil {
				e = err
				return
			}
			list[i] = c
		}(i, c)
	}

	wg.Wait()
	return e
}

func (c _Candidates) Filter(list []*talentio.Candidate, f func(*talentio.Candidate) bool) []*talentio.Candidate {

	var candidates []*talentio.Candidate
	for _, v := range list {
		v := v
		if f(v) {
			candidates = append(candidates, v)
		}
	}

	return candidates
}

// NoResume returns new candidates.
func (c _Candidates) NoResume(list []*talentio.Candidate) []*talentio.Candidate {

	return c.Filter(list, func(c *talentio.Candidate) bool {

		return len(c.Stages) == 0
	})
}

// OngoingResume returns candidates' status is ongoing resume.
func (c _Candidates) OngoingResume(list []*talentio.Candidate) []*talentio.Candidate {

	return c.Filter(list, func(c *talentio.Candidate) bool {

		for _, s := range c.Stages {
			if s.Type != talentio.TypeResume {
				return false
			}

			if s.Status == talentio.StatusOngoing {
				return true
			}
		}

		return false
	})
}

// PassResume returns candidates' status is pass of resume.
func (c _Candidates) PassResume(list []*talentio.Candidate) []*talentio.Candidate {

	return c.Filter(list, func(c *talentio.Candidate) bool {

		if len(c.Stages) == 0 {
			return false
		}

		for _, s := range c.Stages {
			if s.Type != talentio.TypeResume {
				return false
			}

			if s.Status != talentio.StatusPass {
				return false
			}
		}

		return true
	})
}

// OngoingInterview returns candidates' status is ongoing of interview.
func (c _Candidates) OngoingInterview(list []*talentio.Candidate) []*talentio.Candidate {

	return c.Filter(list, func(c *talentio.Candidate) bool {

		for _, s := range c.Stages {
			if s.Type != talentio.TypeInterview {
				continue
			}

			if s.Status == talentio.StatusOngoing {
				return true
			}
		}

		return false
	})
}

// PassInterview returns candidates' status is pass of interview.
func (c _Candidates) PassInterview(list []*talentio.Candidate) []*talentio.Candidate {

	return c.Filter(list, func(c *talentio.Candidate) bool {

		for _, s := range c.Stages {
			if s.Status != talentio.StatusPass {
				return false
			}

			if s.Type == talentio.TypeInterview {
				return true
			}
		}

		return false
	})
}
