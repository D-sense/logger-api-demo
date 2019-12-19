package main

import (
	"context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go-apps/ocr-web-service/cmd/api"
	"go-apps/ocr-web-service/ocr"
	"go-apps/ocr-web-service/services"
	"log"
	"net/http"
)

func main(){
	initContext := context.Background()

	router := mux.NewRouter()

	//Image component
	imgApi := ImageComponent()

	setRoutes(initContext, router, imgApi)

	log.Println("Server has started on PORT 5000 ....")
	log.Fatal(http.ListenAndServe("0.0.0.0:5000",
		handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}

func setRoutes(ctx context.Context, router *mux.Router, imageApi *api.ImageController){
	router.HandleFunc("/image-sync", imageApi.ImageAsync(ctx)).Methods("POST")
	router.HandleFunc("/image", imageApi.CreateImage(ctx)).Methods("POST")
	router.HandleFunc("/image", imageApi.GetTextByID(ctx)).Methods("GET")
}

func ImageComponent() *api.ImageController {
	imgRepo := &ocr.ImageRepo{}

	imgService := services.ImageService {
		imgRepo,
	}

	controller := &api.ImageController{
		imgService,
	}

	return controller
}