package handlers

import (
	"fmt"
	"net/http"

	"github.com/Dimix-international/readwise-go/internal/service"
	"github.com/Dimix-international/readwise-go/internal/utils"

	"github.com/gorilla/mux"
)

type FileHandler struct {
	router      *mux.Router
	fileService service.ServiceFile
}

func NewFileHandler(router *mux.Router, fileService service.ServiceFile) *FileHandler {
	return &FileHandler{router: router, fileService: fileService}
}

func (u *FileHandler) RegisterRoutes() {
	u.router.HandleFunc("/users/{userID}/parse-kindle-file", u.handleParseKindleFile).Methods("POST")
}

func (s *FileHandler) handleParseKindleFile(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userID"]

	file, _, err := r.FormFile("file")
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, fmt.Sprintf("Error parsing file: %v", err))
		return
	}
	defer file.Close()

	if err := s.fileService.ParseKindleFile(r.Context(), &file, userID); err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, "Successfully parsed file")
}
