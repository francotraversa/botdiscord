package pipi

import (
	"fmt"
	"os"
	"os/exec"
)

var cmd *exec.Cmd // Variable global para el proceso

func Execute(point string) {
	if point == "ON" {
		script := "../sapistock/python/main.py"

		cmd = exec.Command("cmd", "/C", "start", "cmd", "/K", "python", script)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Start()
		if err != nil {
			fmt.Println("Error al ejecutar el script:", err)
			return
		}

		fmt.Println("Python ejecutándose en nueva ventana")
	} else {
		fmt.Println("Cerrando la ventana de cmd...")
		exec.Command("taskkill", "/F", "/IM", "cmd.exe").Run()
		fmt.Println("Ventana cerrada correctamente")
	}
}
