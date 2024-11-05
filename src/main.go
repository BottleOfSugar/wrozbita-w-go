package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Definiujemy typ FortuneTeller jako strukturę
type FortuneTeller struct {
	fortunes []string
}

// Tworzymy metodę na FortuneTeller, która losuje wróżbę
func (f *FortuneTeller) TellFortune() string {
	rand.Seed(time.Now().UnixNano())    // Ustawiamy ziarno losowości
	index := rand.Intn(len(f.fortunes)) // Losujemy indeks
	return f.fortunes[index]            // Zwracamy losową wróżbę
}

func main() {
	// Tworzymy nowego wróżbitę z listą wróżb
	fortuneTeller := FortuneTeller{
		fortunes: []string{
			"Wielkie zmiany nadchodzą w Twoim życiu.",
			"Czeka Cię szczęście, ale musisz być cierpliwy.",
			"Nie bój się podjąć ryzyka - to będzie Twoja szansa.",
			"Najbliższe dni przyniosą Ci spokój i radość.",
			"Będziesz musiał podjąć trudną decyzję, ale to dla Ciebie dobra okazja.",
		},
	}

	// Zadajemy pytanie i losujemy wróżbę
	fmt.Println("Zapytaj wróżbitę o przyszłość!")
	fortune := fortuneTeller.TellFortune()

	// Wypisujemy wróżbę
	fmt.Println("Twoja wróżba to: ", fortune)
}
