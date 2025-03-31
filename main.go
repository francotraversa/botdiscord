package main

import (
	"botdiscord/bot" // Importa el paquete bot, que está dentro de la carpeta /bot
	//"botdiscord/python"
	"fmt"
)

func main() {
	fmt.Println("Bot and Financies are starting")
	//python.ActivarPython()
	bot.ConectarADiscord() // Llamamos a la función que conecta a Discord
}
