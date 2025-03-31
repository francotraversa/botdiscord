package bot

import (
	"encoding/json" // Asegúrate de importar este paquete
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
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

func ConectarADiscord() {
	token := "MTM0NjE4ODEwMzc4MTQ1MzkxNw.GFuKI2.Fljl5-wNn-JDrI8zuKgkLNLCRwVxIEdO7XfQxA"

	sess, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatalf("Error al iniciar sesión en Discord: %v", err)
	}

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if strings.HasPrefix(m.Content, "./mervaleta/predict") {
			params := strings.Fields(m.Content)[1:]
			ticker := params[0]
			periodo := params[1]
			intervalo := params[2]
			cmd := exec.Command("python3", "python/ml.py", ticker, periodo, intervalo)
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println("Error ejecutando Python:", err)
				return
			}

			outputText := string(output)
			lines := strings.Split(outputText, "\n")
			handleAPIResponse(s, ("Fecha           | Apertura |  Alto      |  Bajo    |  Cierre   | Volumen"), m.ChannelID)

			for _, line := range lines {
				// Ignorar mensajes de TensorFlow que comienzan con la fecha y "I tensorflow"
				if strings.Contains(line, "I tensorflow") || strings.Contains(line, "oneDNN") || strings.Contains(line, "cpu_feature_guard") {
					continue
				}

				// Filtrar solo fechas válidas o el "Proximo Precio"
				if strings.HasPrefix(line, "2025-") || strings.Contains(line, "Proximo Precio: ") {
					handleAPIResponse(s, line, m.ChannelID)
				}
			}

		}

		if strings.HasPrefix(m.Content, "./mervaleta/tickers/") {
			// Extraemos los parámetros después de "/.mervaleta"
			params := strings.Fields(m.Content)[1:]

			// Verificamos si hay parámetros, si no, respondemos pidiendo los parámetros
			if len(params) == 0 {
				s.ChannelMessageSend(m.ChannelID, "Por favor, ingresa los parámetros después de /.mervaleta")
				return
			}

			sendToAPI(params, s, m.ChannelID)

		}

		/*if strings.HasPrefix(m.Content, "./mervaleta/data/") {
			// Extraemos los parámetros después de "/.mervaleta"
			params := strings.Fields(m.Content)[1:]
			ticker := params[0]
			paramss := params[1]

			// Verificamos si hay parámetros, si no, respondemos pidiendo los parámetros
			if len(params) == 0 {
				s.ChannelMessageSend(m.ChannelID, "Por favor, ingresa los parámetros después de /.mervaleta")
				return
			}
			cmd := exec.Command("python3", "python/activation.py", ticker, paramss)
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println("Error ejecutando Python:", err)
				return
			}
			outputText := string(output)
			fmt.Print(outputText)
			lines := strings.Split(outputText, "\n")
			for _, line := range lines {
				handleAPIResponse(s, line, m.ChannelID)
			}

		}*/

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

func sendToAPI(params []string, s *discordgo.Session, channelID string) {
	baseURL := "http://192.168.1.188:5001/tickers/"
	query := params[0]
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
