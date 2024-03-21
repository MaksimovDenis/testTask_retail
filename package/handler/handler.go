package handler

import (
	"net/http"
	_ "testTask_retail/docs"
	logger "testTask_retail/logs"
	"testTask_retail/package/service"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *http.ServeMux {

	mux := http.NewServeMux()

	mux.Handle("/swagger/", httpSwagger.Handler())

	api := "/api"

	//GET for /api/

	mux.HandleFunc(api+"/orderdetails", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			h.handleGetOrderDetails(w, r)
		} else {
			logger.Log.Errorf("Method now Allowed")
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
