package lichess

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

func GetQuery(ctx context.Context, url string, answer interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(answer)
	if err != nil {
		slog.Error("Failed to decode response", slog.String("url", url), slog.String("error", err.Error()))
		return err
	}

	return nil
}

func GetBroadcast(ctx context.Context, id string) (Broadcast, error) {
	answer := Broadcast{}
	url := GetBroadcastURL(id)

	err := GetQuery(ctx, url, &answer)
	if err != nil {
		return Broadcast{}, err
	}

	return answer, nil
}
