package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
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
	return hostname + "[" + runtime.GOOS + "]"
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
