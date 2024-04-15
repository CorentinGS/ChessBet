package lichess

const (
	// BaseURL is the base URL for the Lichess API.
	BaseURL = "https://lichess.org"

	// BroadcastsURL is the URL for the Lichess broadcasts API.
	BroadcastsURL = BaseURL + "/api/broadcast/"
)

func GetBroadcastsURL() string {
	return BroadcastsURL
}

func GetBaseURL() string {
	return BaseURL
}

// GetBroadcastURL returns the URL for the Lichess broadcast with the given ID.
func GetBroadcastURL(id string) string {
	return BroadcastsURL + id
}

func GetRoundURL(id string) string {
	return BroadcastsURL + id
}
