package main

import (
	"fmt"
	"os"
	"time"

	trackmaniaio "github.com/Brent-Jackson/trackmania-dopamine-simulator/tmio"
	gonadeo "github.com/codecat/gonadeo"
	log "github.com/sirupsen/logrus"
)

const (
	mySweetMap = "lrlw2bUE7wisjz1EPXJsYDkP7De"
)

func main() {

	ubiUsername := os.Getenv("UBISOFT_USERNAME")
	if ubiUsername == "" {
		log.Fatal("no UBISOFT_USERNAME env var provided")
	}
	ubiPassword := os.Getenv("UBISOFT_PASSWORD")
	if ubiPassword == "" {
		log.Fatal("no UBISOFT_PASSWORD env var provided")
	}

	nadeo := gonadeo.NewNadeo()
	err := nadeo.AuthenticateUbi(ubiUsername, ubiPassword)
	if err != nil {
		log.Fatal(err)
	}

	// Set User Agent so people know who we are, and can yell at us if we accidentally ddos anything
	nadeo.SetUserAgent("github.com/Brent-Jackson/trackmania-dopamine-simulator")

	tmio := trackmaniaio.NewClient(nadeo)

	track, err := tmio.GetMap(mySweetMap)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("My sweet track's name is: %s\n", track.Name)

	fmt.Printf("My sweet track's Author Medal time is: %s\n", time.Duration(track.AuthorScore*int(time.Millisecond)))

	leaderboard, err := tmio.GetLeaderboard(mySweetMap)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Here's the leaderboard of my sweet track:\n")
	for _, record := range leaderboard.Tops {
		fmt.Printf("\t%2d: %-16s - %-10s (%s)\n", record.Position, record.Player.Name, time.Duration(record.Time*int(time.Millisecond)), record.Timestamp.Format(time.RFC3339))
	}

}
