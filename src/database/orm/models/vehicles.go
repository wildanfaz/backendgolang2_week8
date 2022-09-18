package models

import "time"

type Vehicle struct {
	VehicleId   uint      `gorm:"primaryKey" json:"vehicle_id"`
	VehicleName string    `json:"vehicle_name"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Status      string    `json:"status"`
	Stock       int       `json:"stock"`
	Category    string    `json:"category"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Rating      float32   `json:"rating"`
}

type Vehicles []Vehicle
