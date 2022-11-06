package main

import (
	"bufio"
	"fmt"
	"math"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var count = 0

func stringToInt(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

func intToString(i int) string {
	val := strconv.Itoa(i)
	return val
}

func verifyPerfectSquare(i int) bool {
	return math.Sqrt(float64(i))*math.Sqrt(float64(i)) == float64(i)
}

var numericRegex = regexp.MustCompile(`[^0-9 ]+`)

func clearString(str string) string {
	return numericRegex.ReplaceAllString(str, "")
}

// function, which takes a string as
// argument and return the reverse of string.
func reverseString(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

func handleConnection(c net.Conn) {
	// Generating the name of the incoming client
	count++
	client := "cmd" + intToString(count)

	// Response from server
	fmt.Println("Clientul " + client + " s-a conectat!")

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))

		if temp == "STOP" {
			break
		}

		// Response from client
		fmt.Println("Clientul " + client + " a trimis request cu datele: " + temp)

		// Processing the request from client
		parametersFromClient := strings.Split(temp, ",")
		numberExercise, _ := strconv.Atoi(parametersFromClient[0])

		// Calculating the result
		result := ""

		// Check the number of the exercise
		if numberExercise == 2 {
			// Count the perfect squares
			counterPerfectSquare := 0

			// Iterate through the parameters
			for i, val := range parametersFromClient {
				// The first index is the number of the exericise
				if i != 0 {
					// Remove all spaces
					val = strings.TrimSpace(val)

					numberToBeVerified := stringToInt(clearString(val))
					if verifyPerfectSquare(numberToBeVerified) == true {
						counterPerfectSquare++
					}
				}
			}

			// Getting the final result
			result = intToString(counterPerfectSquare)

		} else if numberExercise == 3 {
			// The sum of the reversed array
			sum := 0

			// Iterate through the parameters
			for i, val := range parametersFromClient {
				fmt.Println(val)
				// The first index is the number of the exericise
				if i != 0 {
					// Reverse the string
					sum += stringToInt(reverseString(strings.TrimSpace(val)))
				}
			}

			// Getting the final result
			result = intToString(sum)
		} else if numberExercise == 12 {
			// The sum of the reversed array
			sum := 0

			// Iterate through the parameters
			for i, val := range parametersFromClient {
				// The first index is the number of the exericise
				if i != 0 {
					// Remove all spaces
					val = strings.TrimSpace(val)

					// Add the first letter to the start of the string
					tempVal := string(val[0]) + val
					sum += stringToInt(tempVal)
				}
			}

			// Getting the final result
			result = intToString(sum)
		}

		// Sending the result
		fmt.Print("Clientul " + client + " a primit raspunsul ")
		fmt.Printf("%q\n\n", result)

		c.Write([]byte(result + "\n"))
	}
	c.Close()
}

func main() {
	// Open the file
	dat, err := os.ReadFile("config.txt")
	if err != nil {
		fmt.Println("Error in reading the file!")
	}

	// Read from the file
	arguments := string(dat)

	// Check if the port is valid
	if _, err := strconv.Atoi(arguments); err != nil {
		fmt.Printf("%q is not a valid port number! \n\n", arguments)
	}

	// Configure the server
	PORT := ":" + arguments

	// Start the server
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	fmt.Println("The server is now ON!")

	// Connect the clients
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}

}
