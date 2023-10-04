package routes

import (
	"log"
	"net/http"

	"github.com/EduardoStobaus/GoCalculator/controllers"
)

func CalcHandleRequest() {
	http.HandleFunc("/calculator", controllers.Calculator)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
