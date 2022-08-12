package tmio

import (
	"encoding/json"
	"fmt"
	"time"
)

type Leaderboard struct {
	Tops []Top `json:"tops"`
	// cspell: ignore playercount
	PlayerCount int `json:"playercount"`
}

type Top struct {
	Player    *Player    `json:"player"`
	Position  int        `json:"position"`
	Time      int        `json:"time"`
	Filename  string     `json:"filename"`
	Timestamp *time.Time `json:"timestamp"`
	URL       string     `json:"url"`
}

func (c *client) GetLeaderboard(mapUID string) (*Leaderboard, error) {
	buf, err := c.nadeoClient.Get("https://trackmania.io/api/leaderboard/map/"+mapUID, true)
	if err != nil {
		return nil, fmt.Errorf("failed to get leaderboard for map %s: %w", mapUID, err)
	}

	l := Leaderboard{}
	err = json.Unmarshal(buf, &l)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal leaderboard response: %w", err)
	}
	return &l, err

}
