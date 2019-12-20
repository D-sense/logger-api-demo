package main

import (
	"context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go-apps/ocr-web-service/cmd/api"
	"go-apps/ocr-web-service/uploader"
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

	log.Println("Server has started on PORT 8000 ....")
	log.Fatal(http.ListenAndServe("0.0.0.0:8000",
		handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}

func setRoutes(ctx context.Context, router *mux.Router, imageApi *api.ImageController){
	router.HandleFunc("/image_uploader", imageApi.ImageAsync(ctx)).Methods("POST")
}

func ImageComponent() *api.ImageController {
	imgRepo := &uploader.ImageRepo{}

	imgService := services.ImageService {
		imgRepo,
	}

	controller := &api.ImageController{
		imgService,
	}

	return controller
}