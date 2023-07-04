package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())

}

type Laureate struct {
	Firstname  string `json:"firstname"`
	ID         string `json:"id"`
	Motivation string `json:"motivation"`
	Share      string `json:"share"`
	Surname    string `json:"surname"`
}

type Laureates []Laureate

func (l Laureates) SelectRandom() Laureate {
	n := rand.Int() % len(l)
	return l[n]
}

type Prize struct {
	Category  string    `json:"category"`
	Laureates Laureates `json:"laureates"`
	Year      string    `json:"year"`
}

type Prizes []Prize

func (p Prizes) SelectRandom() Prize {
	n := rand.Int() % len(p)
	return p[n]
}

func parsePrizes() (Prizes, error) {
	dataJSON, err := os.ReadFile("prize.json")
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	data := struct {
		Prizes Prizes
	}{}

	if err := json.Unmarshal(dataJSON, &data); err != nil {
		return nil, fmt.Errorf("error parsing json: %w", err)
	}

	var prizes Prizes
	for _, p := range data.Prizes {
		if len(p.Laureates) == 0 {
			continue
		}
		prizes = append(prizes, p)
	}

	return prizes, nil
}

func nobelPrize(w http.ResponseWriter, req *http.Request) {
	handleErr := func(err error) {
		w.WriteHeader(500)
		w.Write([]byte("error: " + err.Error()))
	}

	prizes, err := parsePrizes()
	if err != nil {
		handleErr(err)
		return
	}

	prize := prizes.SelectRandom()
	laureate := prize.Laureates.SelectRandom()

	fmt.Fprintf(w, "[%s|%s] %s %s %s", prize.Year, prize.Category, laureate.Firstname, laureate.Surname, laureate.Motivation)

}

func main() {
	http.HandleFunc("/", nobelPrize)
	http.ListenAndServe(":8000", nil)
}
