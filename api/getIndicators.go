package api

import (
	"botdiscord/discord"
	"botdiscord/types"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func GetIndicatorsFromAPI(ticker string, s *discordgo.Session, channelID string) {
	baseURL := "http://localhost:5001/tickers/data/"
	url := baseURL + ticker

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error en la solicitud:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var apiResponse types.StockData
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		log.Println("Error al transcribir el JSON:", err)
		discord.HandleAPIResponse(s, "No se encontraron tickers en la respuesta de la API.", channelID)

		return
	}

	indicators := fmt.Sprintf("%#v", apiResponse.Indicators)

	indicators = strings.ReplaceAll(indicators, ", ", "\n")
	indicators = strings.ReplaceAll(indicators, ":", ": ")
	indicators = strings.ReplaceAll(indicators, "types.Indicators{", "")
	indicators = strings.ReplaceAll(indicators, "}", "")

	discord.HandleAPIResponse(s, fmt.Sprintf("Ticker: %s\nPrecio: $%.2f\nDecision: %s",
		apiResponse.Ticker, apiResponse.Close, apiResponse.Decision), channelID)
	discord.HandleAPIResponse(s, fmt.Sprintf("Data:\n%s", indicators), channelID)
}
