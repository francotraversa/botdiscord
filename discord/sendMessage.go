package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func HandleAPIResponse(s *discordgo.Session, response string, channelID string) {
	message := response
	_, err := s.ChannelMessageSend(channelID, message)
	if err != nil {
		log.Printf("Error al enviar el mensaje al canal %s: %v", channelID, err)
	} else {

		log.Println("Mensaje enviado con éxito al canal", channelID)
	}
}
