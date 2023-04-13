package Color

import (
	"bufio"
	"flag"
	"fmt"
	WTF "kleurenopdracht/WTF" //import funtie uit ander map
	"os"
)

func Color(outputFile string) { // case swith functie om de input te matchen met de juiste zin.
	var color string
	flag.StringVar(&color, "color", "", "Enter a color")
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your color:")
	if scanner.Scan() {
		Color := scanner.Text()
		switch Color {
		case "blauw":
			WTF.WriteToFile("Blauw zoals de lucht", outputFile)
		case "rood":
			WTF.WriteToFile("Rood met passie", outputFile)
		case "geel":
			WTF.WriteToFile(" Geel als de stralen van de zon", outputFile)
		case "groen":
			WTF.WriteToFile("Groen van de natuur", outputFile)
		default:
			fmt.Println("Unknown color.") //als de input niet de juiste kleuren bevat retour string
		}
	}
}
