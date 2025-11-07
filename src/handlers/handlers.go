package handlers

import (
	"encoding/json"
	"net/http"
	"sample-inventory-go/src/models"
	"strings"
	"sample-inventory-go/src/store"
	"sample-inventory-go/src/utils"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type APIHandler struct {
	Store *store.InventoryStore
}

func NewAPIHandler(s *store.InventoryStore) *APIHandler {
	return &APIHandler{Store: s}
}

func (h *APIHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Hardcoded credentials for demonstration
	if req.Username != "admin" || req.Password != "password" {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	sessionToken := uuid.New().String()
	session := models.Session{
		Token:  sessionToken,
		UserID: req.Username, // Using username as a placeholder for UserID
	}
	h.Store.AddSession(session)

	utils.RespondWithJSON(w, http.StatusOK, utils.JSONResponse{Message: "Login successful", Data: map[string]string{"token": sessionToken}})
}

func (h *APIHandler) Logout(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	token := ""
	if authHeader != "" {
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) == 2 && tokenParts[0] == "Bearer" {
			token = tokenParts[1]
		}
	}

	if token == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Session token missing")
		return
	}

	h.Store.DeleteSession(token)
	utils.RespondWithJSON(w, http.StatusOK, utils.JSONResponse{Message: "Logout successful"})
}

func (h *APIHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if item.ItemCode == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "ItemCode is required")
		return
	}

	if _, ok := h.Store.GetItem(item.ItemCode); ok {
		utils.RespondWithError(w, http.StatusConflict, "Item with this ItemCode already exists")
		return
	}

	h.Store.AddItem(item)
	utils.RespondWithJSON(w, http.StatusCreated, utils.JSONResponse{Message: "Item added successfully", Data: item})
}

func (h *APIHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemCode := vars["itemCode"]

	var updatedItem models.Item
	if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if itemCode == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "ItemCode is required in path")
		return
	}

	if _, ok := h.Store.GetItem(itemCode); !ok {
		utils.RespondWithError(w, http.StatusNotFound, "Item not found")
		return
	}

	updatedItem.ItemCode = itemCode // Ensure the item code from path is used
	h.Store.UpdateItem(updatedItem)
	utils.RespondWithJSON(w, http.StatusOK, utils.JSONResponse{Message: "Item updated successfully", Data: updatedItem})
}

func (h *APIHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemCode := vars["itemCode"]

	if itemCode == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "ItemCode is required in path")
		return
	}

	if _, ok := h.Store.GetItem(itemCode); !ok {
		utils.RespondWithError(w, http.StatusNotFound, "Item not found")
		return
	}

	h.Store.DeleteItem(itemCode)
	utils.RespondWithJSON(w, http.StatusOK, utils.JSONResponse{Message: "Item deleted successfully"})
}

func (h *APIHandler) FetchItems(w http.ResponseWriter, r *http.Request) {
	itemCode := r.URL.Query().Get("item-code")
	name := r.URL.Query().Get("name")
	procurementDate := r.URL.Query().Get("procurement-date")
	expiryDate := r.URL.Query().Get("expiry-date")

	items := h.Store.SearchItems(itemCode, name, procurementDate, expiryDate)
	utils.RespondWithJSON(w, http.StatusOK, utils.JSONResponse{Message: "Items fetched successfully", Data: items})
}
