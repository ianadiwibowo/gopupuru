package handler

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ianadiwibowo/scout-regiment/delivery/serializer"
	"github.com/ianadiwibowo/scout-regiment/model"
	"github.com/jinzhu/gorm"
	"github.com/thedevsaddam/govalidator"
)

// CatsHandler ...
type CatsHandler struct {
	DB         *gorm.DB
	Serializer *serializer.Serializer
}

// NewCatsHandler ...
func NewCatsHandler(db *gorm.DB, serializer *serializer.Serializer) *CatsHandler {
	return &CatsHandler{
		DB:         db,
		Serializer: serializer,
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
	err = h.DB.First(&cat, id).Error

	// Response
	if gorm.IsRecordNotFoundError(err) {
		respondWithNotFound(w)
		return
	}

	respondWithOK(w, h.Serializer.SimpleCat(cat))
}

// Create handles POST /cats
func (h *CatsHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Request validation
	var params createParams
	err := validateCreateParams(r, &params)
	if err != nil {
		respondWithBadRequest(w, err)
		return
	}

	// Process

	// Response
}

type createParams struct {
	Name      string `json:"name"`
	Color     string `json:"color"`
	Naughty   bool   `json:"naughty"`
	Dexterity int    `json:"dexterity"`
}

func validateCreateParams(r *http.Request, params *createParams) url.Values {
	rules := govalidator.MapData{
		"name":      []string{"required"},
		"color":     []string{"required", "in:white,black,orange,green"},
		"naughty":   []string{"bool"},
		"dexterity": []string{"required", "numeric_between:1,10"},
	}
	options := govalidator.Options{
		Request: r,
		Data:    params,
		Rules:   rules,
	}
	validator := govalidator.New(options)

	err := validator.ValidateJSON()
	if len(err) != 0 {
		return err
	}

	return nil
}

// Update handles PATCH /cats/{id}
func (h *CatsHandler) Update(w http.ResponseWriter, r *http.Request) {}

// Delete handles DELETE /cats/{id}
func (h *CatsHandler) Delete(w http.ResponseWriter, r *http.Request) {}
