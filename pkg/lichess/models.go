package lichess

import (
	"strings"
)

type Tour struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	URL         string `json:"url"`
	CreatedAt   int64  `json:"createdAt"`
}

type Round struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	URL       string `json:"url"`
	CreatedAt int64  `json:"createdAt"`
}

type Broadcast struct {
	Tour   Tour    `json:"tour"`
	Rounds []Round `json:"rounds"`
}

func (b Broadcast) String() string {
	sb := strings.Builder{}
	sb.WriteString("Tour: ")
	sb.WriteString(b.Tour.String())
	sb.WriteString("Rounds: ")
	for _, r := range b.Rounds {
		sb.WriteString(r.String())
	}

	return sb.String()
}

func (t Tour) String() string {
	sb := strings.Builder{}
	sb.WriteString("ID: ")
	sb.WriteString(t.ID)
	sb.WriteString("Name: ")
	sb.WriteString(t.Name)
	sb.WriteString("Slug: ")
	sb.WriteString(t.Slug)
	sb.WriteString("Description: ")
	sb.WriteString(t.Description)
	sb.WriteString("URL: ")
	sb.WriteString(t.URL)

	return sb.String()
}

func (r Round) String() string {
	sb := strings.Builder{}
	sb.WriteString("ID: ")
	sb.WriteString(r.ID)
	sb.WriteString("Name: ")
	sb.WriteString(r.Name)
	sb.WriteString("Slug: ")
	sb.WriteString(r.Slug)
	sb.WriteString("URL: ")
	sb.WriteString(r.URL)

	return sb.String()
}
