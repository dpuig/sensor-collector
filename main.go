package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jacobsa/go-serial/serial"
)

type valueRequest struct {
	Timestamp int64   `json:"timestamp"`
	Terminal  string  `json:"terminal"`
	Sensor    string  `json:"sensor"`
	Value     float64 `json:"value"`
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	CollectorAPI := getEnv("COLLECTOR_API", "")
	ArduinoDevice := getEnv("ARDUINO_DEVICE", "")
	Terminal := getEnv("TERMINAL", "")
	Sensor := getEnv("SENSOR", "")

	if ArduinoDevice != "" {
		if CollectorAPI != "" {
			options := serial.OpenOptions{
				PortName:        ArduinoDevice,
				BaudRate:        9600,
				DataBits:        8,
				StopBits:        1,
				MinimumReadSize: 4,
			}
			port, err := serial.Open(options)
			check(err)
			buf := make([]byte, 100)

			for {
				n, err := port.Read(buf)
				check(err)
				vString := string(buf[:n])
				vString = strings.TrimSuffix(vString, "\r\n")
				log.Println(vString)
				i, err := strconv.ParseFloat(vString, 64)
				if err == nil {
					log.Printf("Value Obtained from Sensor: %v", i)
					sendValue(CollectorAPI, Terminal, Sensor, i)
				} else {
					log.Printf("Invalid value Format: %v", err.Error())
				}
				time.Sleep(time.Second)
			}
		} else {
			check(errors.New("No Collector API"))
		}
	} else {
		check(errors.New("No Sensor Device Connected"))
	}
}

func sendValue(CollectorAPI string, t string, s string, v float64) {
	value := valueRequest{
		Timestamp: time.Now().Unix(),
		Terminal:  t,
		Sensor:    s,
		Value:     v,
	}
	jsonValue, _ := json.Marshal(value)
	log.Printf("%+v\n", value)
	u := bytes.NewReader(jsonValue)

	req, err := http.NewRequest("POST", CollectorAPI, u)
	if err != nil {
		fmt.Println("Error is req: ", err)
	}
	req.Header.Set("Content-Type", "application/json")
	// create a Client
	client := &http.Client{}
	// Do sends an HTTP request and
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error in send req: ", err.Error())
	}
	defer resp.Body.Close()
}
