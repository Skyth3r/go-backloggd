package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type currentGame struct {
	Name     string `json:"name"`
	Url      string `json:"url"`
	ImageUrl string `json:"imageUrl"`
}

const baseUrl = "https://backloggd.com"
const username = "BACKLOGGD_USERNAME_HERE"

func main() {

	var currentGames []currentGame
	c := colly.NewCollector()

	c.OnHTML("div.rating-hover", func(e *colly.HTMLElement) {
		game := currentGame{}

		game.Name = e.ChildText("div.game-text-centered")
		partialUrl := e.ChildAttr("a", "href")
		game.Url = baseUrl + partialUrl
		game.ImageUrl = e.ChildAttr("img", "src")

		currentGames = append(currentGames, game)
	})

	c.Visit(baseUrl + "/u/" + username + "/playing/")

	for i := range currentGames {
		fmt.Printf("%v\n", currentGames[i].Name)
		fmt.Printf("%v\n", currentGames[i].Url)
		fmt.Printf("%v\n", currentGames[i].ImageUrl)
	}
}
