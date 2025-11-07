package models

type Item struct {
	ItemCode          string  `json:"item_code"`
	Name              string  `json:"name"`
	Image             string  `json:"image"`
	Description       string  `json:"description"`
	Quantity          int     `json:"quantity"`
	ProcurementDate   *string `json:"procurement_date,omitempty"`
	ManufacturingDate *string `json:"manufacturing_date,omitempty"`
	ExpiryDate        *string `json:"expiry_date,omitempty"`
}

type Session struct {
	Token  string `json:"token"`
	UserID string `json:"user_id"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
