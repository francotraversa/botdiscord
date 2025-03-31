package bot

import (
	"encoding/json" // Asegúrate de importar este paquete
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type Ticker struct {
	CurrentPrice float64 `json:"current_price"`
	TickerName   string  `json:"ticker_name"`
}

// Estructura que contiene los tickers (respuesta completa de la API)
type ApiResponse struct {
	Tickers []Ticker `json:"tickers"`
}

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
			params := strings.Fields(m.Content)[1:] // Extraer solo los tickers

			params2 := "tickers/" + strings.ToUpper(m.Content[7:])

			// Verificamos si hay parámetros, si no, respondemos pidiendo los parámetros
			if len(params) == 0 {
				s.ChannelMessageSend(m.ChannelID, "Por favor, ingresa los parámetros después de mervaleta")
				return
			}

			sendToAPI(params2, s, m.ChannelID)

		}

		if strings.HasPrefix(m.Content, "proyeccion") {
			params := strings.Fields(m.Content)[1:] // Extraer solo los tickers

			params2 := "tickers/data/" + strings.ToUpper(m.Content[11:])
			fmt.Print(params2)
			// Verificamos si hay parámetros, si no, respondemos pidiendo los parámetros
			if len(params) == 0 {
				s.ChannelMessageSend(m.ChannelID, "Por favor, ingresa los parámetros después de mervaleta")
				return
			}

			sendToAPI(params2, s, m.ChannelID)

		}

	})

	//permisos del botardo
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

func sendToAPI(params string, s *discordgo.Session, channelID string) {
	baseURL := "http://192.168.1.188:5001/"
	query := params
	url := baseURL + query

	// Hacer la solicitud GET
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error en la solicitud:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var apiResponse ApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		log.Println("Error al transcribir el JSON:", err)
		handleAPIResponse(s, "NO anduvo.", channelID)
		return
	}

	if len(apiResponse.Tickers) > 0 {
		// Formatear y enviar los datos de la API
		for _, ticker := range apiResponse.Tickers {
			handleAPIResponse(s, fmt.Sprintf("Ticker: %s | Precio: %.2f", ticker.TickerName, ticker.CurrentPrice), channelID)
		}
	} else {
		handleAPIResponse(s, "No se encontraron tickers en la respuesta de la API.", channelID)
	}
}

func handleAPIResponse(s *discordgo.Session, response string, channelID string) {
	// Formatear el mensaje con la respuesta de la API
	message := response

	// Intentar enviar el mensaje al canal de Discord
	_, err := s.ChannelMessageSend(channelID, message)
	if err != nil {
		log.Printf("Error al enviar el mensaje al canal %s: %v", channelID, err)
	} else {
		log.Println("Mensaje enviado con éxito al canal", channelID)
	}
}
