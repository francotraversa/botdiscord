package bot

import (
	"botdiscord/api"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func ConectarADiscord(token string) {
	sess, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatalf("Error al iniciar sesión en Discord: %v", err)
	}

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}
		if strings.HasPrefix(m.Content, "mervaleta") {
			s.ChannelMessageSend(m.ChannelID, "Si Senor!")
		}

		if strings.HasPrefix(m.Content, "precio") {
			params := strings.ToUpper(strings.Fields(m.Content)[1]) // Extraer solo los tickers

			if checkparametros(params, s, m.ChannelID) {
				return
			}

			api.GetPriceFromAPI(params, s, m.ChannelID)

		}

		if strings.HasPrefix(m.Content, "datos") {
			params := strings.ToUpper(strings.Fields(m.Content)[1]) // Extraer solo los tickers

			if checkparametros(params, s, m.ChannelID) {
				return
			}

			api.GetIndicatorsFromAPI(params, s, m.ChannelID)
		}
	})

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	fmt.Println("The bot is online")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func checkparametros(params string, s *discordgo.Session, channelID string) bool {
	if len(params) == 0 {
		s.ChannelMessageSend(channelID, "Por favor, ingresa los tickers")
		return true
	}
	return false
}
