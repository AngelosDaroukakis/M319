package main

import (
	"fmt"
)

func main() {
	// --- 1. SYNTAXFEHLER ("Der Grammatikfehler") ---
	// Hinweis: Ein Tippfehler verhindert das Schmieden (Kompilieren) des Codes.
	fmt.Println("Beschwörung des Debugger-Auges gestartet...")

	// --- 2. LAUFZEITFEHLER ("Der Stolperstein") ---
	// Szenario A: Der leere Zeiger (Nil-Pointer)
	// Ein Zeiger auf einen Namen, der aktuell ins Leere (nil) zeigt.
	var DebuggerAuge *string

	// Aufgabe: Hier muss eine Nil-Prüfung eingebaut werden, bevor
	// der Wert gelesen wird (*DebuggerAuge), sonst gibt es eine Panic!
	if DebuggerAuge != nil {
		fmt.Println("Aktiviere Kraft von:", *DebuggerAuge)
	} else {
		fmt.Println("DebuggerAuge ist nil.")
	}

	// Szenario B: Division durch Null
	energieQuelle := 100
	aktiveGatter := 0
	// Aufgabe: Verhindere den Absturz durch eine Prüfung des Divisors (if != 0).
	if aktiveGatter != 0 {
		lastProGatter := energieQuelle / aktiveGatter
		fmt.Println("Lastverteilung:", lastProGatter)
	} else {
		fmt.Println("Division durch Null verhindert.")
	}

	// --- 3. LOGIKFEHLER ("Die Fehlüberlegung") ---
	// Szenario: Es müssen genau 3 Schutzsiegel aktiviert werden.
	fmt.Println("Aktiviere Schutzsiegel...")
	siegelZaehler := 1

	// Aufgabe: Die Schleife läuft aktuell nicht oft genug.
	// Korrigiere die Bedingung, damit 3 Siegel aktiviert werden.
	for siegelZaehler <= 3 {
		fmt.Printf("Siegel Nr. %d fixiert.\n", siegelZaehler)
		siegelZaehler++
	}

	fmt.Println("Die Kathedrale ist gereinigt!")
}
