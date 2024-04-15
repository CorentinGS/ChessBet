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
	StartsAt  int64  `json:"startsAt"` // Unix timestamp
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

type Player struct {
	Name   string `json:"name"`
	Title  string `json:"title"`
	Fed    string `json:"fed"`
	Rating int    `json:"rating"`
}

type Game struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Status    string   `json:"status"`
	Players   []Player `json:"players"`
	ThinkTime int      `json:"thinkTime"`
}

type Study struct {
	Writeable bool `json:"writeable"`
}

type RoundDetails struct {
	Tour  Tour   `json:"tour"`
	Games []Game `json:"games"`
	Round Round  `json:"round"`
	Study Study  `json:"study"`
}
