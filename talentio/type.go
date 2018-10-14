package talentio

import "time"

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

type (
	AmbiguousTime struct {
		Raw interface{}
	}
)

func (t AmbiguousTime) Time() time.Time {
	if v, ok := t.Raw.(string); ok {
		u, _ := time.Parse(time.RFC3339, v)
		return u
	}
	return time.Time{}
}

func (t AmbiguousTime) After(u time.Time) bool {
	v := t.Time()
	if v.IsZero() {
		return false
	}
	return v.After(u)
}

func (t AmbiguousTime) Before(u time.Time) bool {
	v := t.Time()
	if v.IsZero() {
		return false
	}
	return v.Before(u)
}
