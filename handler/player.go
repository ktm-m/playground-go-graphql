package handler

import (
	"github.com/google/uuid"
	"github.com/ktm-m/playground-go-graphql/graph/model"
)

func GetPlayers() []*model.Player {
	return players
}

func CreatePlayer(req *model.NewPlayer) *model.Player {
	player := &model.Player{
		ID:    uuid.New().String(),
		Name:  req.Name,
		Level: 1,
		Class: &req.Class,
		Items: []*model.Item{
			{
				ID:   "I000001",
				Name: "Item 1",
			},
		},
	}

	players = append(players, player)
	return player
}

var players = []*model.Player{
	{
		ID:    "P000001",
		Name:  "Player 1",
		Level: 1,
		Class: &model.AllPlayerClass[1],
		Items: []*model.Item{
			{
				ID:   "I000001",
				Name: "Item 1",
			},
			{
				ID:   "I000002",
				Name: "Item 2",
			},
		},
	},
	{
		ID:    "P000002",
		Name:  "Player 2",
		Level: 1,
		Class: &model.AllPlayerClass[0],
		Items: []*model.Item{
			{
				ID:   "I000002",
				Name: "Item 2",
			},
		},
	},
}
