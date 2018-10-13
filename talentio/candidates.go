package talentio

import (
	"fmt"
	"time"
)

// CandidatesService handles communication with the candidates related
// methods of the Talentio API.
type CandidatesService service

// CandidatesListOptions represents an order to a related resource.
type CandidatesListOptions struct {
	Page   int    `url:"page"`
	Status string `url:"status"`
	Sort   string `url:"sort"`
}

type (
	// Candidate represents candidate resources.
	Candidate struct {
		ID               int         `json:"id"`
		FirstName        string      `json:"firstName"`
		LastName         string      `json:"lastName"`
		Email            string      `json:"email"`
		Description      string      `json:"description"`
		FixedAt          time.Time   `json:"fixedAt"`
		RegisteredAt     time.Time   `json:"registeredAt"`
		Status           string      `json:"status"`
		Priority         int         `json:"priority"`
		ChannelType      string      `json:"channelType"`
		ChannelName      string      `json:"channelName"`
		Stages           []Stage     `json:"stages"`
		ReferrerEmployee Employee    `json:"referrerEmployee,omitempty"`
		AgentCompany     Company     `json:"agentCompany,omitempty"`
		Requisition      Requisition `json:"requisition"`
		Tags             []Tag       `json:"tags"`
	}

	// Employee represents employee resources.
	Employee struct {
		ID        int    `json:"id"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Active    bool   `json:"active"`
	}

	// Company represents company resources.
	Company struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	// Tag represents tag resources.
	Tag struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	// Requisition represents requisition resources.
	Requisition struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Active bool   `json:"active"`
	}

	// Item represents item resources.
	Item struct {
		Name string `json:"name"`
		Type string `json:"type"`
		// Input   bool   `json:"input"`
		Comment string `json:"comment"`
	}

	// Stage represents stage resources.
	Stage struct {
		ID     int    `json:"id"`
		Type   string `json:"type"`
		Step   int    `json:"step"`
		Status string `json:"status"`
		// ScheduledAt time.Time    `json:"scheduledAt,omitempty"`
		Description string       `json:"description,omitempty"`
		Evaluations []Evaluation `json:"evaluations,omitempty"`
	}

	// Evaluation represents evaluation resources.
	Evaluation struct {
		ID       int      `json:"id"`
		Finished bool     `json:"finished"`
		Items    []Item   `json:"items"`
		Employee Employee `json:"employee"`
	}
)

// List returns a slice of Candidate.
func (s CandidatesService) List(opt *CandidatesListOptions) ([]*Candidate, *Response, error) {
	u, err := addOptions("candidates", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	candidates := new([]*Candidate)
	resp, err := s.client.Do(req, candidates)
	if err != nil {
		return nil, resp, err
	}

	return *candidates, resp, err
}

// Get returns a new Candidate.
func (s CandidatesService) Get(id int) (*Candidate, *Response, error) {
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
