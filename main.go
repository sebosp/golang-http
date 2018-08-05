package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	from := ""
	if r.URL != nil {
		from = r.URL.String()
	}
	log.Printf("Serving: %s\n", from)
	// Some default vars.
	player := &Player{
		name:        "Unset",
		color:       "green",
		environment: "dev",
	}
	err := player.GatherData()
	if err != nil {
		log.Printf("Unable to set Player data: %s\n", err)
		fmt.Fprintf(w, "An error occured loading player data.")
		return
	}
	// No data yet that we should protect.
	log.Printf("%+v\n", player)
	fmt.Fprintf(w, "Hi player")
	fmt.Fprintf(w, "<font color=\""+html.EscapeString(player.color)+"\">\n")
	fmt.Fprintf(w, html.EscapeString(player.name))
	fmt.Fprintf(w, "</font>")
	fmt.Fprintf(w, "<br /> Welcome to env: ")
	fmt.Fprintf(w, html.EscapeString(player.environment))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
