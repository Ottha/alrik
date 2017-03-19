package main

import (
	"flag"
	"fmt"
	"time"
	"strings"
	"strconv"
	"github.com/bwmarrin/discordgo"
	"math/rand"
)

// Variables used for command line parameters
var (
	Token string
	BotID string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(messageCreate)

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Get the account information.
	u, err := dg.User("@me")
	if err != nil {
		fmt.Println("error obtaining account details,", err)
	}

	// Store the account ID for later use.
	BotID = u.ID

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	// Simple way to keep program running until CTRL-C is pressed.
	<-make(chan struct{})
	return
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {


	// Ignore all messages created by the bot itself
	if m.Author.ID == BotID {
		return
	}

	if m.ChannelID != "293007140459773952" {
		fmt.Println(m.ChannelID)
		fmt.Println("Alrik does not care about this channel")
		return
	}

	// Print message to stdout.
	messageBody := m.Message.Content
	if strings.Contains(messageBody, "Alrik") {
		fmt.Printf("%20s %20s %20s > %s\n Hallo, Willkommen im Rusty Peasant!", m.ChannelID, time.Now().Format(time.Stamp), m.Author.Username, m.Content)
	}

	// If the message is "ping" reply with "Pong!"
	if strings.Contains(messageBody, "Alrik") && strings.Contains(messageBody, "hallo") {
		fmt.Printf("trying to send a message")
		_, _ = s.ChannelMessageSend(m.ChannelID, "Hallo, Willkommen im Rusty Peasant!")
	}

	// If the message is "pong" reply with "Ping!"
	if strings.Contains(messageBody, "D") && strings.Contains(messageBody, "Alrik") {
		fmt.Println("Starting to check dice roll")
		var index int = strings.Index(messageBody, "D")
		fmt.Printf("Started finding index of D at %d", index )
		var rolls int
		var err error
		rolls, err = strconv.Atoi(string(messageBody[index-1]))
		fmt.Printf("found rolls: %d ", rolls)
		var faces int
		faces, err = strconv.Atoi(string(messageBody[index+1:len(messageBody)]))
		fmt.Printf("found faces: %d ", faces)
		result := ""

		if err != nil {
			fmt.Println(err)
		}

		for i := 0; i < rolls; i++ {
			var rolled int
			rolled = rand.Intn(faces) + 1
			fmt.Printf("Rolled: %d ", rolled)
			result += strconv.Itoa(rolled)
			result += ", "
		}

		fmt.Printf("trying to send a message")
		_, _ = s.ChannelMessageSend(m.ChannelID, "Du hast " + result + " gewuerfelt")
	}
}