package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	"time"
	//"syscall"
	//"os/signal"
	"strconv"

  "github.com/diamondburned/arikawa/v2/gateway"
	"github.com/diamondburned/arikawa/v2/session"
  "github.com/joho/godotenv"
)

func main() {
	//signal.Ignore(syscall.SIGHUP)
	//signal.Ignore(syscall.SIGTERM)
	//signal.Ignore(os.Interrupt)

  err := godotenv.Load("./.env")
  if err != nil {
    log.Fatalln("Could not read .env file...", err)
  }

  var token = os.Getenv("TOKEN")
	//var funcs = [5]string{"ping","uid","help","ar","color"}

  log.Println("Getting gatewayURL...")
  url, err := gateway.URL()
  if err != nil {
    log.Fatalln("Could not get Websocket URL..", err)
  }
  log.Println(string(url))

  log.Println("Making session")
	s, err := session.New(token)
	if err != nil {
		log.Fatalln("Session failed:", err)
	}
	// login
	log.Println("Session made")

	log.Println("Add handler")
	s.AddHandler(func(c *gateway.MessageCreateEvent) {
		if strings.HasPrefix(c.Content, "#") || strings.HasPrefix(c.Content, "~M") {
			//if c.GuildID.String() == "776489572511121465" {
			if c.Author.ID.String() != "813235572424179742" {
				log.Println(c.Content)
				cmd := strings.TrimPrefix(strings.TrimPrefix(c.Content, "~M"), "#")
				a := strings.Split(string(cmd), " ")
				log.Println(a[0])
				switch a[0] {
				case "ping":
					_, err := s.SendMessage(c.ChannelID, "Ping!", nil)
					if err != nil {
						log.Println("ERROR IN PING")
					}
				case "scream":
					switch strings.Count(c.Content, " ") {
					case 1:
						arg := strings.Split(c.Content, " ")
						fmt.Println(arg[0])
						fmt.Println(arg[1])
						if arg[1] != "" {
							count, _ := strconv.Atoi(string(strings.Split(c.Content, " ")[1]))
							for i := 1; i <= count; i++ {
								fmt.Println(i)
								s.SendMessage(c.ChannelID, fmt.Sprintf("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\naaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"), nil)
							}
						}
					case 0:
						for i := 1; i <= 5; i++ {
							fmt.Println(i)
							s.SendMessage(c.ChannelID, fmt.Sprintf("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\naaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"), nil)
						}
					case 2:
						count, _ := strconv.Atoi(string(strings.Split(c.Content, " ")[1]))
						arg2 := string(strings.Split(c.Content, " ")[2])
						for i := 1; i <= count; i++ {
							s.SendMessage(c.ChannelID, fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s", arg2, arg2, arg2, arg2, arg2, arg2, arg2, arg2, arg2, arg2, arg2, arg2, arg2, arg2, arg2, arg2, arg2, arg2, arg2, arg2), nil)
						}
					case 3:
						count, _ := strconv.Atoi(string(strings.Split(c.Content, " ")[1]))
						arg2 := string(strings.Split(c.Content, " ")[2])
						for i := 1; i <= count; i++ {
							s.SendMessage(c.ChannelID, fmt.Sprintf("%s %s %s %s %s %s %s %s %s %s %s %s %s %s %s %s %s %s %s %s", arg2, arg2, arg2, arg2, arg2, arg2, arg2, arg2, arg2, arg2, arg2, arg2, arg2, arg2, arg2, arg2, arg2, arg2, arg2, arg2), nil)
						}
					}
				case "color":
					// Grab Hex
					//	hex := string(strings.Split(c.Content, " ")[1])

					// Make sure # is present
					//if !strings.HasPrefix(string(strings.Split(c.Content, " ")[1]), "#") {
					//	var hex = "#" + string(strings.Split(c.Content, " ")[1])
					//} else {

					hex := string(strings.Split(c.Content, " ")[1])

					roles, err := s.Roles(c.GuildID)
					if err != nil {
						s.SendMessage(c.ChannelID, "!OPPS! List of Roles could not be generated", nil)
						break
					}
					for i, v := range roles {
						if string(hex) == string(v.Name) {
							s.AddRole(c.GuildID, c.Author.ID, v.ID)
							s.SendMessage(c.ChannelID, fmt.Sprintf("Set <@!%s> to color %s", c.Author.ID.String(), hex), nil)
							break
						}
						if i == len(roles)-1 {
							s.SendMessage(c.ChannelID, fmt.Sprintf("~role create %s %s", hex, hex), nil)
							s.SendMessage(c.ChannelID, "Please wait...", nil)
							time.Sleep(10 * time.Second)
							roles, err := s.Roles(c.GuildID)
							if err != nil {
								s.SendMessage(c.ChannelID, "!OPPS! List of Roles could not be generated", nil)
								break
							}
							for i, v := range roles {
								if string(hex) == string(v.Name) {
									s.AddRole(c.GuildID, c.Author.ID, v.ID)
									s.SendMessage(c.ChannelID, fmt.Sprintf("Set <@!%s> to color %s", c.Author.ID.String(), hex), nil)
									break
								}
								if i == len(roles)-1 {
									s.SendMessage(c.ChannelID, "PAIN: !!!OPPS!!! SOMETHING WENT WRONG: FAILED TO FIND NEW ROLE ID", nil)
									break
								}
							}
						}
					}
				case "help":
					s.SendMessage(c.ChannelID, fmt.Sprintf("```\nMiato:\n  A lovely little scripted user running on Mia's laptop!\nUsage:\n[ping] - Return 'Ping!'\n[scream] - (scream <count> [scream content]) Screams for you, duh...\n[color] - (color <#color>) Assign's a hex color role to m.Author.ID - Requires Carl-bot + Admin\n```"), nil)
				}
			}
		}
	})

	// Add the needed Gateway intents.
	log.Println("Adding intents")
	s.Gateway.AddIntents(gateway.IntentGuildMessages)

	log.Println("Connecting..")
	if err := s.Open(); err != nil {
		log.Fatalln("Failed to connect:", err)
	}
	defer s.Close()
	log.Println("Connected!")

	log.Println("Getting myself...")
	u, err := s.Me()
	if err != nil {
		log.Fatalln("Failed to get myself:", err)
	}
	log.Println("Gotten!")

	log.Println("Started as", u.Username)

	// Block forever.
	//log.Println("Blocking forever...")
	select {}
}
