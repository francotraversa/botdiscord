package main

import (
	"botdiscord/bot" // Importa el paquete bot, que está dentro de la carpeta /bot
	//"botdiscord/python"
	"botdiscord/environment"
	"fmt"
)

func main() {
	fmt.Println("Bot and Financies are starting")
	//python.ActivarPython()
	envProvider := environment.EnvProvider{}
	environment.LoadDotEnv(envProvider)
	environment.InitializeEnvVariables(envProvider)
	bot.ConectarADiscord(environment.GetEnv().Token) // Llamamos a la función que conecta a Discord
}
