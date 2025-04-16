package fbo

import "time"

type ListRequestFilter struct {
	Since  time.Time `json:"since"`
	Status string    `json:"status"`
	To     time.Time `json:"to"`
}

type ListRequestWith struct {
	AnalyticsData bool `json:"analytics_data"`
	FinancialData bool `json:"financial_data"`
}

type ListRequest struct {
	Dir      string            `json:"dir"`
	Filter   ListRequestFilter `json:"filter"`
	Limit    int64             `json:"limit"`
	Offset   int64             `json:"offset"`
	Translit bool              `json:"translit"`
	With     ListRequestWith   `json:"with"`
}

type ListResponseResultProduct struct {
	SKU          int64    `json:"sku"`
	Name         string   `json:"name"`
	Quantity     int64    `json:"quantity"`
	OfferID      string   `json:"offer_id"`
	Price        string   `json:"price"`
	DigitalCodes []string `json:"digital_codes"`
	CurrencyCode string   `json:"currency_code"`
}

type ListResponseResultAnalyticsData struct {
	Region               string `json:"region"`
	City                 string `json:"city"`
	DeliveryType         string `json:"delivery_type"`
	IsPremium            bool   `json:"is_premium"`
	PaymentTypeGroupName string `json:"payment_type_group_name"`
	WarehouseID          int64  `json:"warehouse_id"`
	WarehouseName        string `json:"warehouse_name"`
	IsLegal              bool   `json:"is_legal"`
}

type ListResponseResultFinancialDataProductPicking struct {
	Amount float64   `json:"amount"`
	Moment time.Time `json:"moment"`
	Tag    string    `json:"tag"`
}

type ListResponseResultFinancialDataProduct struct {
	CommissionAmount     float64                                        `json:"commission_amount"`
	CommissionPercent    int64                                          `json:"commission_percent"`
	Payout               float64                                        `json:"payout"`
	ProductID            int64                                          `json:"product_id"`
	CurrencyCode         string                                         `json:"currency_code"`
	OldPrice             float64                                        `json:"old_price"`
	Price                float64                                        `json:"price"`
	TotalDiscountValue   float64                                        `json:"total_discount_value"`
	TotalDiscountPercent float64                                        `json:"total_discount_percent"`
	Actions              []string                                       `json:"actions"`
	Picking              *ListResponseResultFinancialDataProductPicking `json:"picking"`
	Quantity             int64                                          `json:"quantity"`
	ClientPrice          string                                         `json:"client_price"`
}

type ListResponseResultFinancialData struct {
	Products []ListResponseResultFinancialDataProduct `json:"products"`
}

type ListResponseResultAdditionalData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ListResponseResult struct {
	OrderID        int64                              `json:"order_id"`
	OrderNumber    string                             `json:"order_number"`
	PostingNumber  string                             `json:"posting_number"`
	Status         string                             `json:"status"`
	CancelReasonID int64                              `json:"cancel_reason_id"`
	CreatedAt      time.Time                          `json:"created_at"`
	InProcessAt    time.Time                          `json:"in_process_at"`
	Products       []ListResponseResultProduct        `json:"products"`
	AnalyticsData  ListResponseResultAnalyticsData    `json:"analytics_data"`
	FinancialData  ListResponseResultFinancialData    `json:"financial_data"`
	AdditionalData []ListResponseResultAdditionalData `json:"additional_data"`
}

type ListResponse struct {
	Result []ListResponseResult `json:"result"`
}
