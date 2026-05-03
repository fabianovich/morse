package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
	"go.bug.st/serial"
)

var morseMap = map[string]string{
	".-":    "A", "-...":  "B", "-.-.":  "C", "-..":   "D",
	".":     "E", "..-.":  "F", "--.":   "G", "....":  "H",
	"..":    "I", ".---":  "J", "-.-":   "K", ".-..":  "L",
	"--":    "M", "-.":    "N", "---":   "O", ".--.":  "P",
	"--.-":  "Q", ".-.":   "R", "...":   "S", "-":     "T",
	"..-":   "U", "...-":  "V", ".--":   "W", "-..-":  "X",
	"-.--":  "Y", "--..":  "Z",
	"-----": "0", ".----": "1", "..---": "2", "...--": "3",
	"....-": "4", ".....": "5", "-....": "6", "--...": "7",
	"---..": "8", "----.": "9",
}

func decode(morse string) string {
	letter, ok := morseMap[morse]
	if !ok {
		return "?"
	}
	return letter
}

func main(){
	mode := &serial.Mode{
		BaudRate: 115200,
	}
	port, err := serial.Open("/dev/ttyUSB0", mode)
	if err != nil {
			log.Fatal(err)
	}
	defer port.Close()
	
	scanner := bufio.NewScanner(port)
	morse := ""
	message := ""
	
	for scanner.Scan() {
		data := scanner.Text()

		if strings.Contains(data, ".") {
			morse += "."
		} else if strings.Contains(data, "-") {
			morse += "-"
		} else {
			decoded := decode(morse)
			morse = ""
			message += decoded
		}

		fmt.Printf(message, morse)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("serial error: %v", err)
	}
}