package main

import (
	"fmt"
	"log"
	"os/exec"
)

func call_bat(message Message) (output string) {
	cmd := exec.Command("cmd", "/C", message.Command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("error with:", err.Error())
	} else {
		return ("Host:" + get_hostname() + "Message:" + string(out))
	}
	return
}

func call_batscript(message Message) (output string) {
	cmd := exec.Command("cmd", "/C", message.Command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("error with:", err.Error())
	} else {
		return ("Host:" + get_hostname() + "Message:" + string(out))
	}
	return
}

func call_powershell(message Message) (output string) {
	cmd := exec.Command("powershell", "-Command", message.Command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("error with:", err.Error())
	} else {
		return ("Host:" + get_hostname() + "Message:" + string(out))
	}
	return
}

func call_pwsh(message Message) (output string) {
	cmd := exec.Command("pwsh", "-Command", message.Command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("error with:", err.Error())
	} else {
		return ("Host:" + get_hostname() + "Message:" + string(out))
	}
	return
}

func call_powershell_script(message Message) (output string) {
	cmd := exec.Command("cmd", "/C", message.Command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("error with:", err.Error())
	} else {
		return ("Host:" + get_hostname() + "Message:" + string(out))
	}
	return
}

func call_pwsh_script(message Message) (output string) {
	cmd := exec.Command("cmd", "/C", message.Command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("error with:", err.Error())
	} else {
		return ("Host:" + get_hostname() + "Message:" + string(out))
	}
	return
}

func call_bash(message Message) (output string) {
	cmd := exec.Command("bash", "-c", message.Command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("error with:", err.Error())
	} else {
		return ("Host:" + get_hostname() + "Message:" + string(out))
	}
	return
}

func call_bash_script(message Message) (output string) {
	cmd := exec.Command("bash", "-c", message.Command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("error with:", err.Error())
	} else {
		return ("Host:" + get_hostname() + "Message:" + string(out))
	}
	return
}
