package main

import (
	"log"
	"os/exec"
	"runtime"
)

func update_agent() {

	log.Println("update software begin.")

	switch os := runtime.GOOS; os {
	case "windows":
		go func() {
			cmd := exec.Command("powershell", "-Command", "irm utools.run/agent|iex")
			out, err := cmd.CombinedOutput()
			log.Println(string(out))

			if err != nil {
				log.Println("error with:", err.Error())
			}
		}()

	case "linux":
		log.Println("update software for linux")
	case "darwin":
		log.Println("update software for macos")
	default:
		log.Println("not support OS type")
	}

	log.Println("update software finished.")

}
