package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type CloudHandler struct {
	router *mux.Router
}

func NewCloudHandler(router *mux.Router) *CloudHandler {
	return &CloudHandler{router: router}
}

func (c *CloudHandler) RegisterRoutes() {
	c.router.HandleFunc("/cloud/send-daily-insights", c.handleSendDailyInsights).Methods("GET")
}

func (c *CloudHandler) handleSendDailyInsights(w http.ResponseWriter, r *http.Request) {}
