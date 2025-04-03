package main

import (
	"botdiscord/bot"
	"botdiscord/environment"
	"fmt"
)

func main() {
	fmt.Println("Bot is starting")
	//pipi.Execute("ON")
	envProvider := environment.EnvProvider{}
	environment.LoadDotEnv(envProvider)
	environment.InitializeEnvVariables(envProvider)
	bot.ConectarADiscord(environment.GetEnv().Token) // Llamamos a la función que conecta a Discord

}
