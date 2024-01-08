package main

import (
	"log"
	"os/exec"
	"runtime"
)

func check_update(strUpdateCmd string) {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("powershell", "-Command", strUpdateCmd)
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Println("error with:", err.Error())
		} else {
			log.Println(string(out))
			log.Println("Host:" + get_hostname() + "Message:update successfully")
		}
	}

}
