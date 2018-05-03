package api

import "time"

type (
	Query struct {
		Granularity string    `json:"granularity,omitempty"`
		Start       time.Time `json:"start,omitempty"`
		End         time.Time `json:"end,omitempty"`
		Saved       *Saved    `json:"saved,omitempty"`
		Filter      *Filter   `json:"filter,omitempty"`
		Metric      string    `json:"metric,omitempty"`
	}

	Saved struct {
		Name string `json:"name,omitempty"`
	}

	Selection struct {
		Aps    []string `json:"aps,omitempty"`
		SSIDs  []string `json:"ssids,omitempty"`
		Radios []string `json:"radios,omitempty"`
	}
)

// Filter is a generic loopback construct
type Filter struct {
	Type      string    `json:"type,omitempty"`
	Dimension string    `json:"dimension,omitempty"`
	Value     string    `json:"value,omitempty"`
	Fields    []*Filter `json:"fields,omitempty"`
}

func NewFilter(operator string) *Filter {
	return &Filter{
		Type:   operator,
		Fields: make([]*Filter, 0),
	}
}

func (f *Filter) AppendFieldFilter(field *Filter) {
	f.Fields = append(f.Fields, field)
}
