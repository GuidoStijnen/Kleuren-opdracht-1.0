package WTF

import (
	"fmt"
	"log"
	"os"
)

func WriteToFile(color string, outputFile string) {
	file, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) //0644 geeft premissies voor bewerken aan de owner en degene die txt leest.
	if err != nil {
		log.Fatalf("Error opening file: %v", err) //log.Fatal logged de error en verlaat het programma als de file niet geopend kan worden, %vfromateer het agrument in default
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, color)			 
	if err != nil {
		log.Printf("Error writing to file: %v", err) //error handeling voor als er een fout is van het schrijven naar de txt file 
		return
	}

	fmt.Println("Wrote", color, "to file.") // valideerd de output voor de gerbuiker 
}
