package main

import (
	"fmt"
	"os/exec"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func handler(message Message, client mqtt.Client) {
	if message.OS == "win" && message.Type == "powershell" {
		cmd := exec.Command("powershell", "-Command", message.Command)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("error with:", err.Error())
			publish_log(client, err.Error())
		} else {
			fmt.Println(string(out))
			publish_log(client, string(out))
		}
	} else if message.OS == "win" && message.Type == "pwsh" {
		cmd := exec.Command("pwsh", "-Command", message.Command)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("error with:", err.Error())
			publish_log(client, err.Error())
		} else {
			fmt.Println(string(out))
			publish_log(client, string(out))
		}
	} else if message.OS == "win" && message.Type == "bat" {
		cmd := exec.Command("cmd", "/C", message.Command)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("error with:", err.Error())
			publish_log(client, err.Error())
		} else {
			fmt.Println(string(out))
			publish_log(client, string(out))
		}
	} else if message.OS == "win" && message.Type == "python_script" {
		cmd := exec.Command("cmd", "/C", message.Command)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("error with:", err.Error())
			publish_log(client, err.Error())
		} else {
			fmt.Println(string(out))
			publish_log(client, string(out))
		}
	} else if message.OS == "win" && message.Type == "powershell_script" {
		cmd := exec.Command("cmd", "/C", message.Command)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("error with:", err.Error())
			publish_log(client, err.Error())
		} else {
			fmt.Println(string(out))
			publish_log(client, string(out))
		}

		// {
		// 	"type": "powershell_script",
		// 	"command":"curl -s https://it2u.oss-cn-shenzhen.aliyuncs.com/scripts/test.ps1 | powershell -NoLogo"
		// }
	} else if message.OS == "win" && message.Type == "bat_script" {
		cmd := exec.Command("cmd", "/C", message.Command)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("error with:", err.Error())
			publish_log(client, err.Error())
		} else {
			fmt.Println(string(out))
			publish_log(client, string(out))
		}

		// {
		// 	"type": "bat_script",
		// 	"command":"curl -s https://it2u.oss-cn-shenzhen.aliyuncs.com/scripts/test.bat -o a.bat && a.bat"
		// }

	} else {
		publish_log(client, "command type error!")
	}

}
