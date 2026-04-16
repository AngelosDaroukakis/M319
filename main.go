package main

import "fmt"

func main() {

	/*
	 * Fragment 1: Die unendliche Bewässerung
	 * Die Ranken werden ohne Ende mit magischem Wasser überflutet,
	 * da das festgelegte Limit nie erreicht wird.
	 */
	wasser := 0
	limit := 10

	fmt.Println("--- Start: Bewässerung ---")
	for wasser < limit {
		fmt.Println("Gieße magisches Wasser...")
		fmt.Println("Limit: ", limit)
		wasser++ // Inkrement hinzugefügt
	}
	fmt.Println("Bewässerung abgeschlossen! \n")

	/*
	 * Fragment 2: Die fehlerhafte Licht-Zufuhr
	 * In diesem Fall wurde zwar ein Inkrement eingebaut, jedoch ist die
	 * Logik der Abbruchbedingung fehlerhaft, wodurch der Wert sich
	 * immer weiter vom Ziel entfernt.
	 */

	lichtEnergie := 100

	fmt.Println("--- Start: Licht-Zufuhr ---")
	for lichtEnergie > 0 {
		fmt.Println("Lichtstrahl wird fokussiert...")
		lichtEnergie-- // statt ++ muss der Wert kleiner werden
		fmt.Println("Lichtenergie: ", lichtEnergie)
	}
	fmt.Println("Lichtenergie verbraucht! \n")

	/*
	 * Fragment 3: Der blockierte Wächter-Check
	 * Hier wird eine Bedingung geprüft, die innerhalb der Schleife
	 * niemals verändert werden kann. Der Wächter wartet vergeblich auf ein Signal.
	 */

	istGezähmt := false
	versuche := 0

	fmt.Println("--- Start: Wächter-Check ---")
	for !istGezähmt {
		fmt.Printf("Versuch %d: Besänftige die Ranken...\n", versuche)
		fmt.Println("Versuche: ", versuche)
		versuche++

		if versuche >= 10 {
			istGezähmt = true
		}
	}

	fmt.Println("Mission abgeschlossen: Die Ranken sind gezähmt!")
}
