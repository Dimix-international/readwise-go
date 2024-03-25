package handlers

import (
	"fmt"
	"net/http"

	"github.com/Dimix-international/readwise-go/internal/service"
	"github.com/Dimix-international/readwise-go/internal/utils"
	"github.com/gorilla/mux"
)

type CloudHandler struct {
	router       *mux.Router
	serviceCloud service.ServiceCloud
}

func NewCloudHandler(router *mux.Router, serviceCloud service.ServiceCloud) *CloudHandler {
	return &CloudHandler{router: router, serviceCloud: serviceCloud}
}

func (c *CloudHandler) RegisterRoutes() {
	c.router.HandleFunc("/cloud/send-daily-insights", c.handleSendDailyInsights).Methods("GET")
}

func (c *CloudHandler) handleSendDailyInsights(w http.ResponseWriter, r *http.Request) {
	if err := c.serviceCloud.SendInsightsEmails(r.Context()); err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, fmt.Sprintf("Error get users: %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}
