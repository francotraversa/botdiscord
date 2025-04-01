package api

import (
	"botdiscord/discord"
	"botdiscord/environment"
	"botdiscord/types"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

func GetPriceFromAPI(tickers string, s *discordgo.Session, channelID string) {
	endpoint := "tickers/"
	baseUrl := environment.GetEnv().BaseURL
	url := baseUrl + endpoint + tickers

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
		discord.HandleAPIResponse(s, "No se encontraron tickers en la respuesta de la API.", channelID)
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
