package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KernelGamut32/golang-microservices/inventory/internal/inventory"
	"github.com/gorilla/mux"
)

var inventoryService *InventoryService

func Get() *InventoryService {
	if inventoryService == nil {
		inventoryService = &InventoryService{DB: GetInventoryDataStore()}
		return inventoryService
	}
	return inventoryService
}

type InventoryService struct {
	DB inventory.InventoryDatastore
}

func (is *InventoryService) SetInitial(w http.ResponseWriter, r *http.Request) {
	inventory := &inventory.Inventory{}
	json.NewDecoder(r.Body).Decode(inventory)
	if err := is.DB.CreateInventory(inventory); err != nil {
		log.Print("error occured when creating new inventory ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(inventory)
}

func (is *InventoryService) UpdateCurrent(w http.ResponseWriter, r *http.Request) {
	inventory := inventory.Inventory{}
	params := mux.Vars(r)
	var id = params["id"]

	json.NewDecoder(r.Body).Decode(&inventory)

	if err := is.DB.UpdateInventory(id, inventory); err != nil {
		log.Print("error occured when updating inventory ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&inventory)
}

func (is *InventoryService) GetCurrent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var prod_num = params["productNumber"]

	inventory, err := is.DB.GetInventory(prod_num)

	if err != nil {
		log.Print("error occured when getting inventory ", err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(&inventory)
}
