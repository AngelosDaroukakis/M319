package main

import (
	"fmt"
)

func main() {
	// --- 1. SYNTAXFEHLER ("Das zerbrochene Siegel") ---
	// Ein fehlendes Zeichen oder ein Tippfehler verhindert den Start der Heilung.
	fmt.Println("Initialisiere Heilungs-Protokoll...")

	// --- 2. LAUFZEITFEHLER ("Die instabile Energie") ---
	// Szenario: Ein Speicherzugriff auf einen nicht existierenden Datensatz.
	var heilungsKristall *int

	// AUFGABE: Der Kristall ist noch 'nil'. Wenn wir versuchen, seine Energie
	// zu nutzen (*), stürzt die Realität ab.
	// Baue eine Prüfung ein: Wenn nil, gib eine Meldung aus, sonst nutze die Energie.
	if heilungsKristall == nil {
		fmt.Println("HeilungsKristall ist nil.")
	} else {
		fmt.Printf("Energie-Level des Kristalls: %d\n", *heilungsKristall)
	}

	// --- 3. LOGIKFEHLER ("Der verdorrte Wald") ---
	// Szenario: Um ein Tal zu heilen, müssen mindestens 50 Energie-Einheiten
	// fließen. Aktuell ist die Bedingung jedoch falsch programmiert.
	energieFluss := 40
	istHeilungErfolgreich := false

	// AUFGABE: Die Heilung sollte nur erfolgreich sein, wenn der Fluss
	// GRÖSSER ODER GLEICH 50 ist. Korrigiere die Logik.
	if energieFluss >= 50 {
		istHeilungErfolgreich = true
	}

	if istHeilungErfolgreich {
		fmt.Println("✅ Die Aura breitet sich aus! Die Highlands erblühen.")
	} else {
		fmt.Println("❌ Die Energie reicht nicht aus. Die Highlands bleiben grau.")
	}

	// --- 4. SCHLEIFEN-FEHLER ("Die ewige Dürre") ---
	// Szenario: Es müssen 5 Bäume gepflanzt werden, aber die Zählung ist fehlerhaft.
	baeumeGepflanzt := 0
	fmt.Println("Beginne Aufforstung...")

	// AUFGABE: Die Bedingung ist so gewählt, dass die Schleife niemals startet.
	// Korrigiere sie so, dass genau 5 Bäume gepflanzt werden (0 bis 4).
	for baeumeGepflanzt < 5 {
		fmt.Printf("Baum Nr. %d gepflanzt.\n", baeumeGepflanzt)
		baeumeGepflanzt++
	}

	fmt.Println("Mission abgeschlossen: Der Weltencode ist rein!")
}
