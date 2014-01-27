package main

import (
	"container/list"
	"errors"
	"log"
)

type Game struct {
	Name  string `json:"name"`
	Clubs []Club `json:"clubs"`
}

func (g Game) FetchAll() []Game {
	rows, err := DB.Query("select name from games")
	if err != nil {
		log.Fatal(err)
	}

	result := list.New()
	for rows.Next() {
		var game Game
		if err := rows.Scan(&game.Name); err != nil {
			log.Fatal(err)
		} else {
			result.PushBack(game)
		}

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	games := make([]Game, result.Len())
	for e, i := result.Front(), 0; e != nil; e, i = e.Next(), i+1 {
		games[i] = e.Value.(Game)
	}
	return games
}

func (g Game) FetchByName(name string) (Game, error) {
	rows, err := DB.Query("select c.name, c.league, c.stars from clubs c join games g where c.game = g.id and g.name=?", name)
	if err != nil {
		log.Fatal(err)
	}

	result := list.New()
	count := 0
	for ; rows.Next(); count++ {
		var club Club
		if err := rows.Scan(&club.Name, &club.League, &club.Stars); err != nil {
			log.Fatal(err)
		} else {
			result.PushBack(club)
		}

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		return Game{}, errors.New("Could not find game")
	}

	clubs := make([]Club, count)
	for e, i := result.Front(), 0; e != nil; e, i = e.Next(), i+1 {
		clubs[i] = e.Value.(Club)
	}

	return Game{Name: name, Clubs: clubs}, nil
}

type Club struct {
	Name   string  `json:"name"`
	League string  `json:"league"`
	Stars  float64 `json:"stars"`
}