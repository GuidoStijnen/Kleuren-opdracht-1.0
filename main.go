package main

import (
	"flag"
	GetColor "kleurenopdracht/Color" //import functie uit andere map
	"log"                            //error loggin
	"os"     
)

func main() {

	outputFile := flag.String("output", "output.txt", "name of the output file") //flag toegevoegd voor het valideren van de input
	flag.Parse()

	logFile, err := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // opend een log file waar de errors van de writetofile functie instaan, 0644 zijn premissie die gegeven worden voor toegang tot het error log file.
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	GetColor.Color(*outputFile) //roept functie voor de input van de code en in de functie roept die ook nog de WTF funtie aan
}
