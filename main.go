package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/krognol/go-wolfram"
	"github.com/shomali11/slacker"
	"github.com/tidwall/gjson"
	witai "github.com/wit-ai/wit-go/v2"
	"log"
	"os"
)

var wolframClient *wolfram.Client

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {

	godotenv.Load(".env")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	client := witai.NewClient(os.Getenv("WIT_AI_TOKEN"))

	wolframClient := &wolfram.Client{
		AppID: os.Getenv("WOLFRAM_APP_ID"),
	}

	go printCommandEvents(bot.CommandEvents())

	var theArray []string = []string{"Who is the president of Ukraine"}

	bot.Command("query for bot - <message>", &slacker.CommandDefinition{
		Description: "Send any question to wolfram",
		Examples:    theArray,
		Handler: func(botContext slacker.BotContext, request slacker.Request, writer slacker.ResponseWriter) {
			query := request.Param("message")
			fmt.Println(query)
			msg, _ := client.Parse(&witai.MessageRequest{
				Query: query,
			})
			data, _ := json.MarshalIndent(msg, "", "    ")
			rough := string(data[:])
			value := gjson.Get(rough, "entities.wit$wolfram_search_query:wolfram_search_query.0.value")
			question := value.String()
			fmt.Println(question)
			res, err := wolframClient.GetSpokentAnswerQuery(question, wolfram.Metric, 1000)
			if err != nil {
				fmt.Println("Wolfram error")
			}
			writer.Reply(res)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)

	if err != nil {
		log.Fatal(err)
	}
}
