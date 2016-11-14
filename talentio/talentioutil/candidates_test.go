package talentioutil_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kaneshin/go-talentio/talentio"
	"github.com/kaneshin/go-talentio/talentio/talentioutil"
)

var testCandidates = []*talentio.Candidate{
	&talentio.Candidate{
		ID: 1,
		Stages: []talentio.Stage{
			stageOngoingResume,
		},
	},
	&talentio.Candidate{
		ID: 2,
		Stages: []talentio.Stage{
			stageOngoingResume,
			stagePassResume,
		},
	},
	&talentio.Candidate{
		ID:     3,
		Stages: []talentio.Stage{},
	},
	&talentio.Candidate{
		ID:     4,
		Stages: []talentio.Stage{},
	},
	&talentio.Candidate{
		ID: 5,
		Stages: []talentio.Stage{
			stagePassResume,
		},
	},
	&talentio.Candidate{
		ID: 6,
		Stages: []talentio.Stage{
			stagePassResume,
			stageOngoingInterview,
		},
	},
	&talentio.Candidate{
		ID: 7,
		Stages: []talentio.Stage{
			stagePassResume,
			stagePassResume,
			stageOngoingInterview,
		},
	},
	&talentio.Candidate{
		ID: 8,
		Stages: []talentio.Stage{
			stagePassResume,
			stagePassResume,
			stagePassResume,
		},
	},
	&talentio.Candidate{
		ID: 9,
		Stages: []talentio.Stage{
			stagePassResume,
			stageOngoingInterview,
			stagePassInterview,
		},
	},
	&talentio.Candidate{
		ID: 10,
		Stages: []talentio.Stage{
			stagePassResume,
			stagePassInterview,
			stagePassInterview,
		},
	},
}

var (
	stagePassResume       = talentio.Stage{Type: "resume", Status: "pass"}
	stageOngoingResume    = talentio.Stage{Type: "resume", Status: "ongoing"}
	stagePassInterview    = talentio.Stage{Type: "interview", Status: "pass"}
	stageOngoingInterview = talentio.Stage{Type: "interview", Status: "ongoing"}
)

func TestCandidatesMap(t *testing.T) {

	t.Run("NoResume", func(t *testing.T) {
		var expected []*talentio.Candidate
		expected = append(expected, testCandidates[2])
		expected = append(expected, testCandidates[3])
		assert.Equal(t, expected, talentioutil.Candidates.NoResume(testCandidates))
	})

	t.Run("OngoingResume", func(t *testing.T) {
		var expected []*talentio.Candidate
		expected = append(expected, testCandidates[0])
		expected = append(expected, testCandidates[1])
		assert.Equal(t, expected, talentioutil.Candidates.OngoingResume(testCandidates))
	})

	t.Run("PassResume", func(t *testing.T) {
		var expected []*talentio.Candidate
		expected = append(expected, testCandidates[4])
		expected = append(expected, testCandidates[7])
		assert.Equal(t, expected, talentioutil.Candidates.PassResume(testCandidates))
	})

	t.Run("OngoingInterview", func(t *testing.T) {
		var expected []*talentio.Candidate
		expected = append(expected, testCandidates[5])
		expected = append(expected, testCandidates[6])
		expected = append(expected, testCandidates[8])
		assert.Equal(t, expected, talentioutil.Candidates.OngoingInterview(testCandidates))
	})

	t.Run("PassInterview", func(t *testing.T) {
		var expected []*talentio.Candidate
		expected = append(expected, testCandidates[9])
		assert.Equal(t, expected, talentioutil.Candidates.PassInterview(testCandidates))
	})
}
