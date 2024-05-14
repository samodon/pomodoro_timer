package main

import (
	"context"
	"fmt"
	"time"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type Session struct {
	name                string
	lenght, breakLenght int
	songs, dnd          bool
}

func spotifyAuth() *oauth2.Token {
	authConfig := &clientcredentials.Config{
		ClientID:     "",
		ClientSecret: "",
		TokenURL:     spotify.TokenURL,
	}

	accessToken, err := authConfig.Token(context.Background())

	if err != nil {
		fmt.Println("Error getting access token: ", err)
		return nil
	} else {
		return accessToken
	}
}

func startSession(session Session) {
	start_time := time.Now()
	duration := 0

	// DND mode
	if session.dnd {
		fmt.Println("DND mode enabled")
	}

	// Play music
	if session.songs {
		fmt.Println("Playing music")

		client := spotify.Authenticator{}.NewClient(spotifyAuth())
		// begin playing music
		fmt.Println(client)
	}

	// Start session
	show := true
	for duration < (session.lenght * 60) {
		duration = int(time.Since(start_time).Seconds())

		if session.lenght*60-duration < 60 && show {
			fmt.Println("Session is about to end in 1 minute, begin wrapping up")
			show = false
		}
	}

	// Break session
	breaksession(session.breakLenght * 60)
}

func breaksession(seconds int) {
	start_time := time.Now()
	duration := 0

	for duration < (seconds) {
		duration = int(time.Since(start_time).Seconds())
	}
}

func main() {
	newsession := true
	for newsession {

		var session Session
		fmt.Print("Enter the name of this session: ")
		fmt.Scan(&session.name)

		fmt.Print("Do you want to listen to music during this session? (true/false): ")
		fmt.Scan(&session.songs)

		fmt.Print("Do you want to enable DND mode during this session? (true/false): ")
		fmt.Scan(&session.dnd)

		fmt.Print("Enter the lenght of this session in minutes: ")
		fmt.Scan(&session.lenght)

		fmt.Print("Enter the lenght of the break in minutes: ")
		fmt.Scan(&session.breakLenght)

		continueSession := true
		for continueSession {

			startSession(session)

			fmt.Print("Do you want to continue this session? (true/false): ")
			fmt.Scan(&continueSession)
		}

		fmt.Print("Do you want to start a new session? (true/false): ")
		fmt.Scan(&newsession)
	}
}
