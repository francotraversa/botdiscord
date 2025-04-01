package api

import (
	"botdiscord/discord"
	"botdiscord/types"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

func GetPriceFromAPI(ticker string, s *discordgo.Session, channelID string) {
	baseURL := "http://localhost:5001/tickers"
	url := baseURL + ticker

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error en la solicitud:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var apiResponse types.ApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		log.Println("Error al transcribir el JSON:", err)
		discord.HandleAPIResponse(s, "NO anduvo.", channelID)
		return
	}

	if len(apiResponse.Tickers) > 0 {
		for _, ticker := range apiResponse.Tickers {
			discord.HandleAPIResponse(s, fmt.Sprintf("Ticker: %s | Precio: $%.2f", ticker.TickerName, ticker.CurrentPrice), channelID)
		}
	} else {
		discord.HandleAPIResponse(s, "No se encontraron tickers en la respuesta de la API.", channelID)
	}
}
