package tmio

import (
	"encoding/json"
	"fmt"
	"time"
)

type Map struct {
	Author         string     `json:"author"`
	Name           string     `json:"name"`
	MapType        string     `json:"mapType"`
	MapStyle       string     `json:"mapStyle"`
	AuthorScore    int        `json:"authorScore"`
	GoldScore      int        `json:"goldScore"`
	SilverScore    int        `json:"silverScore"`
	BronzeScore    int        `json:"bronzeScore"`
	CollectionName string     `json:"collectionName"`
	Filename       string     `json:"filename"`
	IsPlayable     bool       `json:"isPlayable"`
	MapID          string     `json:"mapId"`
	MapUID         string     `json:"mapUid"`
	Submitter      string     `json:"submitter"`
	Timestamp      *time.Time `json:"timestamp"`
	FileURL        string     `json:"fileUrl"`
	ThumbnailURL   string     `json:"thumbnailUrl"`
	// cspell:ignore authorplayer submitterplayer exchangeid
	AuthorPlayer    *Player `json:"authorplayer"`
	SubmitterPlayer *Player `json:"submitterplayer"`
	ExchangeID      int     `json:"exchangeid"`
}

func (c *client) GetMap(mapUID string) (*Map, error) {
	buf, err := c.nadeoClient.Get("https://trackmania.io/api/map/"+mapUID, true)
	if err != nil {
		return nil, fmt.Errorf("failed to get map %s: %w", mapUID, err)
	}

	m := Map{}
	err = json.Unmarshal(buf, &m)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal map response: %w", err)
	}
	return &m, err

}
