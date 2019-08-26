package main

import (
	"fmt"

	"github.com/TuringCom/golang_challenge_template/models"
	"github.com/TuringCom/golang_challenge_template/pkg/util"
	"github.com/TuringCom/golang_challenge_template/routers"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	models.Setup()
}

func main() {
	r := routers.SetupRouter()
	// Listen and Server in 0.0.0.0:8080

	port := util.GetEnv("PORT", "8000")

	r.Run(fmt.Sprintf(":%s", port))
}
