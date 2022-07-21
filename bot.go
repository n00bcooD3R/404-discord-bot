package main

import (
	"fmt"
	discord "github.com/sleeyax/aternos-discord-bot"
	database "github.com/sleeyax/aternos-discord-bot/database"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Create bot instance with an in-memory database (only supports 1 discord server).
	bot := discord.Bot{
		DiscordToken:  "OTk5NTU1MzcwMTM5Nzg3MzI1.GR_8dR.Ubot6kequ8Or9Wsrl6mTMXY8CBQoYjNbdDvhZc",
		Database: &database.MemoryDatabase{}, // no initial values are provided, so they must be set with `/configure` later on
	}
	
	// Start the bot (errors are omitted for simplicity reasons).
	bot.Start()
	
	// Stop the bot when the main function ends.
	defer bot.Stop()
	
	// You can put some additional code here, for example. 
	//...

	// Block the main thread so the bot keeps running.
	// In this case we wait until 'CTRL + C' or another termination signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	interruptSignal := make(chan os.Signal, 1)
	signal.Notify(interruptSignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-interruptSignal
}