package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	CONNECT := arguments[1]
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Menu from the client's user interface
	menu := "Te-ai conectat la server!\n"
	menu += "-> Ex2. Serverul returnează către client numărul de numere care sunt pătrate perfecte.\n"
	menu += "-> Ex3. Serverul răspunde către client cu suma numerelor array-ului format prin inversarea" +
		" fiecărui element din array-ul inițial.\n"
	menu += "-> Ex12. Serverul returnează către client suma elementelor array-ului format din dublarea" +
		" primei cifre a fiecărui număr.\n"
	menu += "--> Introduceti nr exercitiului selectat, urmat de lista de parametri necesari.\n"
	fmt.Println(menu)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')

		// Sending the data to the server
		fmt.Fprintf(c, text+"\n")

		fmt.Println("Server-ul a primit request-ul")
		fmt.Println("Server-ul proceseaza datele...")

		// Receving the data from server
		message, _ := bufio.NewReader(c).ReadString('\n')
		message = strings.Replace(message, "\n", "", -1)

		fmt.Print("Server-ul trimite raspunsul ")
		fmt.Printf("%q", message)
		fmt.Print(" catre client.\n")

		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}
