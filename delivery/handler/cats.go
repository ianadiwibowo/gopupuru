package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ianadiwibowo/scout-regiment/delivery/serializer"
	"github.com/ianadiwibowo/scout-regiment/model"
	"github.com/jinzhu/gorm"
)

// CatsHandler ...
type CatsHandler struct {
	DB         *gorm.DB
	Serializer *serializer.Serializer
}

// NewCatsHandler ...
func NewCatsHandler(db *gorm.DB) *CatsHandler {
	return &CatsHandler{
		DB:         db,
		Serializer: serializer.New(),
	}
}

// Index handles GET /cats
func (h *CatsHandler) Index(w http.ResponseWriter, r *http.Request) {}

// Show handles GET /cats/{id}
func (h *CatsHandler) Show(w http.ResponseWriter, r *http.Request) {
	// Request validation
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if id == 0 || err != nil {
		respondWithUnprocessableEntity(w)
		return
	}

	// Process
	cat := model.Cat{}
	h.DB.First(&cat, id)

	// Respond
	if cat.ID == 0 {
		respondWithNotFound(w)
		return
	}

	respondWithSuccess(w, h.Serializer.SimpleCat(cat))
}

// Create handles POST /cats
func (h *CatsHandler) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You called POST /cats")
}

// Update handles PATCH /cats/{id}
func (h *CatsHandler) Update(w http.ResponseWriter, r *http.Request) {}

// Delete handles DELETE /cats/{id}
func (h *CatsHandler) Delete(w http.ResponseWriter, r *http.Request) {}
