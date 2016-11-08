package talentio

import (
	"fmt"
	"net/http"
	"time"
)

// A CandidatesService is.
type CandidatesService service

// A CandidatesListOptions is.
type CandidatesListOptions struct {
	Page   int    `url:"page"`
	Status string `url:"status"`
	Sort   string `url:"sort"`
}

const (
	StatusOngoing      = "ongoing"
	StatusReject       = "reject"
	StatusFail         = "fail"
	StatusPass         = "pass"
	StatusPoolActive   = "poolActive"
	StatusPoolInactive = "poolInactive"
)

const (
	RegisteredAtAscKey  = "registeredAt"
	RegisteredAtDescKey = "-registeredAt"
)

type (
	// A Candidate is.
	Candidate struct {
		ID               int         `json:"id"`
		FirstName        string      `json:"firstName"`
		LastName         string      `json:"lastName"`
		Email            string      `json:"email"`
		Description      string      `json:"description"`
		RegisteredAt     time.Time   `json:"registeredAt"`
		Status           string      `json:"status"`
		ChannelType      string      `json:"channelType"`
		Priority         int         `json:"priority"`
		Stage            []Stage     `json:"stages"`
		ReferrerEmployee Employee    `json:"referrerEmployee,omitempty"`
		AgentCompany     Company     `json:"agentCompany,omitempty"`
		Requisition      Requisition `json:"requisition"`
		Tags             []Tag       `json:"tags"`
	}

	// An Employee is.
	Employee struct {
		ID        int    `json:"id"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Active    bool   `json:"active"`
	}

	// A Company is.
	Company struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	// A Tag is.
	Tag struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	// A Requisition is.
	Requisition struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Active bool   `json:"active"`
	}

	// A Item is.
	Item struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}

	// A Stage is.
	Stage struct {
		ID          int          `json:"id"`
		Type        string       `json:"type"`
		Step        int          `json:"step"`
		Status      string       `json:"status"`
		ScheduledAt time.Time    `json:"scheduledAt,omitempty"`
		Description string       `json:"description,omitempty"`
		Evaluations []Evaluation `json:"evaluations,omitempty"`
	}

	// An Evaluation is.
	Evaluation struct {
		ID       int      `json:"id"`
		Finished bool     `json:"finished"`
		Items    []Item   `json:"items"`
		Employee Employee `json:"employee"`
	}
)

// List returns a pointer of Candidates.
func (s CandidatesService) List(opt *CandidatesListOptions) ([]*Candidate, *http.Response, error) {
	u, err := addOptions("candidates", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println(req.URL.String())

	candidates := new([]*Candidate)
	resp, err := s.client.Do(req, candidates)
	if err != nil {
		return nil, resp, err
	}

	return *candidates, resp, err
}

// Get returns a pointer of Candidate.
func (s CandidatesService) Get(id int) (*Candidate, *http.Response, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("candidates/%d", id), nil)
	if err != nil {
		return nil, nil, err
	}

	candidate := new(Candidate)
	resp, err := s.client.Do(req, candidate)
	if err != nil {
		return nil, resp, err
	}

	return candidate, resp, err
}
