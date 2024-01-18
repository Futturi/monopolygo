package handler

import (
	"awesomeProject/internal/models"
	"log"
	"strconv"

	"net/http"

	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

func (h *Handler) Game(c *gin.Context) {
	room_id, err := strconv.Atoi(c.Param("room_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	users := h.service.GetUsersByRoomId(room_id)
	usersGame := make([]models.UserGame, 0)
	for _, user := range users {
		username, err := h.service.GetUsernameById(user)
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})
		}
		usersGame = append(usersGame, models.UserGame{
			Userid:    user,
			Username:  username,
			Money:     15000,
			Position:  0,
			Business:  make([]models.Business, 0),
			Monopolys: make([]int, 0),
		})
	}

	businesses := []models.Business{
		{Name: "Home", Price: 0, PriceIfAnotherIn: 0, PriceHome: 0, OneHome: 0, TwoHome: 0, ThreeHome: 0, FourHome: 0, Hotel: 0, NumberOfMonopoly: 1},
		{Name: "Житная улица", Price: 60, Owner: models.UserGame{}, PriceIfAnotherIn: 10, PriceHome: 50, OneHome: 50, TwoHome: 100, ThreeHome: 150, FourHome: 200, Hotel: 250, NumberOfMonopoly: 1},
		{Name: "Нагатинская улица", Price: 120, Owner: models.UserGame{}, PriceIfAnotherIn: 20, PriceHome: 100, OneHome: 100, TwoHome: 200, ThreeHome: 300, FourHome: 400, Hotel: 500, NumberOfMonopoly: 2},
		{Name: "Варшавское шоссе", Price: 160, Owner: models.UserGame{}, PriceIfAnotherIn: 30, PriceHome: 150, OneHome: 150, TwoHome: 300, ThreeHome: 450, FourHome: 600, Hotel: 750, NumberOfMonopoly: 2},
		{Name: "Улица Огарева", Price: 220, Owner: models.UserGame{}, PriceIfAnotherIn: 40, PriceHome: 200, OneHome: 200, TwoHome: 400, ThreeHome: 600, FourHome: 800, Hotel: 1000, NumberOfMonopoly: 2},
		{Name: "Первая Парковая Улица", Price: 240, Owner: models.UserGame{}, PriceIfAnotherIn: 50, PriceHome: 250, OneHome: 250, TwoHome: 500, ThreeHome: 750, FourHome: 1000, Hotel: 1250, NumberOfMonopoly: 3},
		{Name: "Улица Полянка", Price: 300, Owner: models.UserGame{}, PriceIfAnotherIn: 60, PriceHome: 300, OneHome: 300, TwoHome: 600, ThreeHome: 900, FourHome: 1200, Hotel: 1500, NumberOfMonopoly: 3},
		{Name: "Улица Сретенка", Price: 350, Owner: models.UserGame{}, PriceIfAnotherIn: 70, PriceHome: 350, OneHome: 350, TwoHome: 700, ThreeHome: 1050, FourHome: 1400, Hotel: 1750, NumberOfMonopoly: 3},
		{Name: "Ростовская набережная", Price: 400, Owner: models.UserGame{}, PriceIfAnotherIn: 80, PriceHome: 400, OneHome: 400, TwoHome: 800, ThreeHome: 1200, FourHome: 1600, Hotel: 2000, NumberOfMonopoly: 4},
		{Name: "Рязанский проспект", Price: 450, Owner: models.UserGame{}, PriceIfAnotherIn: 90, PriceHome: 450, OneHome: 450, TwoHome: 900, ThreeHome: 1350, FourHome: 1800, Hotel: 2250, NumberOfMonopoly: 4},
		{Name: "Улица Вавилова", Price: 500, Owner: models.UserGame{}, PriceIfAnotherIn: 100, PriceHome: 500, OneHome: 500, TwoHome: 1000, ThreeHome: 1500, FourHome: 2000, Hotel: 2500, NumberOfMonopoly: 4},
		{Name: "Рублевское Шоссе", Price: 550, Owner: models.UserGame{}, PriceIfAnotherIn: 110, PriceHome: 550, OneHome: 550, TwoHome: 1100, ThreeHome: 1650, FourHome: 2200, Hotel: 2750, NumberOfMonopoly: 5},
		{Name: "Улица Тверская", Price: 600, Owner: models.UserGame{}, PriceIfAnotherIn: 120, PriceHome: 600, OneHome: 600, TwoHome: 1200, ThreeHome: 1800, FourHome: 2400, Hotel: 3000, NumberOfMonopoly: 5},
		{Name: "Пушкинская улица", Price: 650, Owner: models.UserGame{}, PriceIfAnotherIn: 130, PriceHome: 650, OneHome: 650, TwoHome: 1300, ThreeHome: 1950, FourHome: 2600, Hotel: 3250, NumberOfMonopoly: 5},
		{Name: "Площадь Маяковского", Price: 700, Owner: models.UserGame{}, PriceIfAnotherIn: 140, PriceHome: 700, OneHome: 700, TwoHome: 1400, ThreeHome: 2100, FourHome: 2800, Hotel: 3500, NumberOfMonopoly: 6},
		{Name: "Улица Грузнский вал", Price: 750, Owner: models.UserGame{}, PriceIfAnotherIn: 150, PriceHome: 750, OneHome: 750, TwoHome: 1500, ThreeHome: 2250, FourHome: 3000, Hotel: 3750, NumberOfMonopoly: 6},
		{Name: "Новинский бульвар", Price: 800, Owner: models.UserGame{}, PriceIfAnotherIn: 160, PriceHome: 800, OneHome: 800, TwoHome: 1600, ThreeHome: 2400, FourHome: 3200, Hotel: 4000, NumberOfMonopoly: 6},
		{Name: "Смоленская площадь", Price: 850, Owner: models.UserGame{}, PriceIfAnotherIn: 170, PriceHome: 850, OneHome: 850, TwoHome: 1700, ThreeHome: 2550, FourHome: 3400, Hotel: 4250, NumberOfMonopoly: 7},
		{Name: "Гоголевский бульвар", Price: 900, Owner: models.UserGame{}, PriceIfAnotherIn: 180, PriceHome: 900, OneHome: 900, TwoHome: 1800, ThreeHome: 2700, FourHome: 3600, Hotel: 4500, NumberOfMonopoly: 7},
		{Name: "Кутузовский проспект", Price: 950, Owner: models.UserGame{}, PriceIfAnotherIn: 190, PriceHome: 950, OneHome: 950, TwoHome: 1900, ThreeHome: 2850, FourHome: 3800, Hotel: 4750, NumberOfMonopoly: 7},
		{Name: "Улица малая бронная", Price: 1000, Owner: models.UserGame{}, PriceIfAnotherIn: 200, PriceHome: 1000, OneHome: 1000, TwoHome: 2000, ThreeHome: 3000, FourHome: 4000, Hotel: 5000, NumberOfMonopoly: 8},
		{Name: "Улица Арбат", Price: 1050, Owner: models.UserGame{}, PriceIfAnotherIn: 210, PriceHome: 1050, OneHome: 1050, TwoHome: 2100, ThreeHome: 3150, FourHome: 4200, Hotel: 5250, NumberOfMonopoly: 8},
	}

	board := models.Board{
		Businesses:     businesses,
		PlayerPosition: []int{0, 0, 0, 0},
	}

	mel := melody.New()
	mel.HandleConnect(func(s *melody.Session) {
		log.Println("connected")
	})

	for {
		for _, user := range usersGame {
			mel.HandleMessage(func(s *melody.Session, msg []byte) {
				switch string(msg) {
				case "roll":
					user.Position += rand.Intn(13)
					board.PlayerPosition[user.Userid] = user.Position
					mel.Broadcast([]byte(string(user.Position)))
				case "buy":
					user.Business = append(user.Business, businesses[user.Position])
					businesses[user.Position].Owner = user
					mel.Broadcast([]byte(user.Username + " bought " + businesses[user.Position].Name))
				}
				if businesses[user.Position].Owner.Userid != user.Userid {
					user.Money -= businesses[user.Position].PriceIfAnotherIn
				}
			})
		}
	}
}
