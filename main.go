package main

import (
	"log"
	"os"

	"api.namegame.com/api"
	"api.namegame.com/api/controllers"
	"api.namegame.com/api/routes"
	"api.namegame.com/commons"
	"api.namegame.com/domains"
	"api.namegame.com/jobs"
	"api.namegame.com/messaging"
	"api.namegame.com/messaging/emitters"
	"api.namegame.com/repositories"
	"api.namegame.com/services"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	initAPI()
}

func initAPI() {
	firebaseClient, err := messaging.FirebaseConfig{}.CreateClient()

	if err != nil {
		panic(err)
	}

	mongoClient := commons.MongoClient{}
	mongoDatabase := mongoClient.GetDatabase(os.Getenv("MONGO_URI"),
		"name-game")
	nameStatisticsRepository := repositories.NameStatisticsRepository{
		DB: mongoDatabase,
	}

	hunchRoundRepository := repositories.HunchRoundRepository{
		HunchRounds: make(map[string]domains.HunchRound, 0),
	}

	roomStateRepository := repositories.RoomStateRepository{
		RoomStates: make(map[string]domains.RoomState, 0),
	}

	scoreboardRepository := repositories.ScoreboardRepository{
		Scoreboards: make(map[string]domains.Scoreboard, 0),
	}

	roomStateEmitter := emitters.RoomState{
		RoomStateRepository:  roomStateRepository,
		ScoreboardRepository: scoreboardRepository,
		FirebaseClient:       firebaseClient,
	}

	scoreboardEmitter := emitters.Scoreboard{
		ScoreboardRepository: scoreboardRepository,
		FirebaseClient:       firebaseClient,
	}

	roundJob := jobs.RoundJob{HunchRoundRepository: hunchRoundRepository,
		RoomStateEmitter:     roomStateEmitter,
		ScoreboardRepository: scoreboardRepository,
		RoomStateRepository:  roomStateRepository,
	}

	roomService := services.RoomService{
		HunchRoundRepository:     hunchRoundRepository,
		RoomStateRepository:      roomStateRepository,
		ScoreboardRepository:     scoreboardRepository,
		ScoreboardEmitter:        scoreboardEmitter,
		RoomStateEmitter:         roomStateEmitter,
		NameStatisticsRepository: nameStatisticsRepository,
		RoundJob:                 roundJob,
	}

	roomController := controllers.RoomController{RoomService: roomService}

	baseRoutes := make([]routes.BaseRoute, 0)
	roomRoute := routes.RoomRoute{RoomController: roomController}
	baseRoutes = append(baseRoutes, roomRoute)
	server := api.Server{Routes: baseRoutes}

	r := server.SetupRoutes()
	r.Run(":" + os.Getenv("APP_PORT"))
}
