package lichess

const (
	// BaseURL is the base URL for the Lichess API.
	BaseURL = "https://lichess.org"

	// BroadcastsURL is the URL for the Lichess broadcasts API.
	BroadcastsURL = BaseURL + "/api/broadcast/"
)

// GetBroadcastURL returns the URL for the Lichess broadcast with the given ID.
func getBroadcastURL(id string) string {
	return BroadcastsURL + id
}

func getRoundURL(id string) string {
	return BroadcastsURL + "-/-/" + id
}
