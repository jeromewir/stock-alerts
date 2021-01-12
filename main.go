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
	_ = godotenv.Load()

	err := config.ReadFromEnvironment()

	if err != nil {
		log.Fatal(err)
	}

	messengerIDs := config.GetMessengerRecipientIDs()

	ticker := time.NewTicker(5 * time.Minute)

	parsers := []Parser{
		// stockparsers.NewFnacParser(),
		stockparsers.NewAmazonFRParser(),
		stockparsers.NewAmazonESParser(),
		stockparsers.NewAmazonDEParser(),
		stockparsers.NewAmazonCOUKParser(),
		stockparsers.NewAmazonITParser(),
		stockparsers.NewAuchanParser(),
		stockparsers.NewCDiscountParser(),
		stockparsers.NewCarrefourParser(),
		stockparsers.NewLeclercParser(),
	}

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
				fmt.Println("Checking for availabilities")
				for _, p := range parsers {
					jobs <- Job{Parser: p}
				}

				break
			case job := <- jobs:
				go func(job Job) {
					isAvailable, err := job.Parser.IsAvailable()

					for _, mID := range messengerIDs {
						if err != nil {
							fmt.Println(job.Parser.GetName(), err)
							m.SendSimpleMessage(mID, fmt.Sprintf("Impossible de verifier les stocks pour %s: %s", job.Parser.GetName(), err.Error()))
				
							return
						}
	
						fmt.Printf("%s: %t\n", job.Parser.GetName(), isAvailable)
	
						if isAvailable == true {
							m.SendSimpleMessage(mID, fmt.Sprintf("Duuuude, PS5 dispo chez %s! 🏃‍♂️\n%s", job.Parser.GetName(), job.Parser.GetURL()))
						}
					}
				}(job)
				break
			}
		}
	}()

	http.HandleFunc("/webhook", m.Handler)

	mr := func(event messenger.Event, opts messenger.MessageOpts, msg messenger.ReceivedMessage) {
		fmt.Println(fmt.Sprintf("Received a message from %s: %s", opts.Sender.ID, msg.Text))

		_, err := m.SendSimpleMessage(opts.Sender.ID, "Je verifie les dispos 👇")

		if err != nil {
			fmt.Println(err)
		}

		for _, p := range parsers {
			go func(p Parser) {
				isAvailable, err := p.IsAvailable()

				if err != nil {
					m.SendSimpleMessage(opts.Sender.ID, fmt.Sprintf("Impossible de verifier les stocks pour %s: %s", p.GetName(), err.Error()))
		
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

	http.HandleFunc("/", func (res http.ResponseWriter, req *http.Request) {
		res.Header().Set("content-type", "application/json")
		res.WriteHeader(200)
		res.Write([]byte("{\"message\": \"ok\"}"))
	})

	http.ListenAndServe(":5646", nil)
}