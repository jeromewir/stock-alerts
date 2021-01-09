package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	config "github.com/jeromewir/stockalerts/config"
	stockparsers "github.com/jeromewir/stockalerts/parsers"
	"github.com/joho/godotenv"
	"github.com/maciekmm/messenger-platform-go-sdk"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	err = config.ReadFromEnvironment()

	if err != nil {
		log.Fatal(err)
	}

	messengerIDs := config.GetMessengerRecipientIDs()

	ticker := time.NewTicker(5 * time.Minute)
	parsers := make([]Parser, 0)

	parsers = append(parsers, stockparsers.NewFnacParser())
	parsers = append(parsers, stockparsers.NewAuchanParser())
	parsers = append(parsers, stockparsers.NewCDiscountParser())

	jobs := make(chan Job, len(parsers))

	m := &messenger.Messenger {
		AccessToken: config.GetMessengerAccessToken(),
		Debug: messenger.DebugAll, //All,Info,Warning
		VerifyToken: config.GetMessengerVerifyToken(),
	}

	go func() {
		for {
			select {
			case <- ticker.C:
				for _, p := range parsers {
					jobs <- Job{Parser: p}
				}

				break
			case job := <- jobs:
				isAvailable, err := job.Parser.IsAvailable()

				for _, mID := range messengerIDs {
					if err != nil {
						m.SendSimpleMessage(mID, fmt.Sprintf("Impossible de verifier les stocks: %s", err.Error()))
			
						return
					}

					if isAvailable == true {
						m.SendSimpleMessage(mID, fmt.Sprintf("Duuuude, PS5 dispo chez %s! 🏃‍♂️\n%s", job.Parser.GetName(), job.Parser.GetURL()))
					}
				}
				break
			}
		}
	}()

	http.HandleFunc("/webhook", m.Handler)

	mr := func(event messenger.Event, opts messenger.MessageOpts, msg messenger.ReceivedMessage) {
		m.SendSimpleMessage(opts.Sender.ID, "Je verifie les dispos 👇")

		for _, p := range parsers {
			go func(p Parser) {
				isAvailable, err := p.IsAvailable()

				if err != nil {
					m.SendSimpleMessage(opts.Sender.ID, fmt.Sprintf("Impossible de verifier les stocks: %s", err.Error()))
		
					return
				}

				isAvailableS := "🥵"

				if isAvailable == true {
					isAvailableS = "YES! ⚡"
				}

				m.SendSimpleMessage(opts.Sender.ID, fmt.Sprintf("%s: %s (%s)", p.GetName(), isAvailableS, p.GetShortURL()))
			}(p)
		}
	}

	m.MessageReceived = mr

	http.ListenAndServe(":5646", nil)
}