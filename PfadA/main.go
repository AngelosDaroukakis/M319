package main

import "fmt"

/*
 * TEIL 1: Die Basis-Auftriebsfunktion
 * Ein Falke benötigt eine Grundkraft, um in der Luft zu bleiben.
 * Diese berechnet sich aus Windgeschwindigkeit + Anflugwinkel.
 */
func berechneBasisAuftrieb(windgeschwindigkeit int, anflugwinkel int) int {
	ergebnis := windgeschwindigkeit + anflugwinkel
	return ergebnis
}

/*
 * TEIL 2: Der Energieverbrauch
 * Fliegen verbraucht Energie. Je schwerer der Falke, desto mehr.
 * Formel: (Gewicht * Strecke) / 10
 */
func berechneEnergieVerbrauch(gewicht int, strecke int) int {
	ergebnis := (gewicht * strecke) / 10
	return ergebnis
}

/*
 * TEIL 3: Die finale Flug-Validierung
 */
func checkFlugStatus(auftrieb int, verbrauch int) (int, string) {
	differenz := auftrieb - verbrauch
	status := ""

	if differenz > 0 {
		status = "Stabil"
	} else {
		status = "Absturzgefahr"
	}

	return differenz, status
}

/*
 * TEIL 4: Das freudige Kreischen des Falkens
 */
func kreischen() {
	fmt.Println("KKKRREEEIIIISSSCCCHHHHH")
}

func main() {
	fmt.Println("--- Analyse der Variable-Falken startet ---")

	// 1. Aufruf der Basis-Funktion
	speed := 45
	winkel := 12
	aktuellerAuftrieb := berechneBasisAuftrieb(speed, winkel)
	fmt.Printf("Berechneter Auftrieb: %d Einheiten\n", aktuellerAuftrieb)

	// 2. Aufruf der Energie-Funktion
	gewicht := 20
	distanz := 20
	verbrauch := berechneEnergieVerbrauch(gewicht, distanz)
	fmt.Printf("Voraussichtlicher Verbrauch: %d Einheiten\n", verbrauch)

	// 3. Finale Prüfung
	kraftReserve, flugZustand := checkFlugStatus(aktuellerAuftrieb, verbrauch)

	fmt.Println("-------------------------------------------")
	fmt.Printf("Analyse-Ergebnis: %s (Reserve: %d)\n", flugZustand, kraftReserve)

	if flugZustand == "Stabil" && kraftReserve > 0 {
		fmt.Println("Mission abgeschlossen: Der Falke hält seine Bahn!")
		kreischen()
		kreischen()
		kreischen()
	} else {
		fmt.Println("Fehler: Die mathematische Kapselung ist noch instabil.")
	}
}
