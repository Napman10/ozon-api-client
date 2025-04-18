package transaction

import "time"

type ListRequestFilterDate struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

type ListRequestFilter struct {
	Date            ListRequestFilterDate `json:"date"`
	OperationType   []string              `json:"operation_type"`
	PostingNumber   string                `json:"posting_number"`
	TransactionType string                `json:"transaction_type"`
}

type ListRequest struct {
	Filter   ListRequestFilter `json:"filter"`
	Page     int64             `json:"page"`
	PageSize int64             `json:"page_size"`
}

type ListResponseResultOperationPosting struct {
	DeliverySchema ListResponseResultOperationPostingDeliverySchema `json:"delivery_schema"`
	OrderDate      string                                           `json:"order_date"`
	PostingNumber  string                                           `json:"posting_number"`
	WarehouseID    int64                                            `json:"warehouse_id"`
}

type ListResponseResultOperationItem struct {
	Name string `json:"name"`
	SKU  int64  `json:"sku"`
}

type ListResponseResultOperationService struct {
	Name  ListResponseResultOperationServiceName `json:"name"`
	Price float64                                `json:"price"`
}

type ListResponseResultOperation struct {
	AccrualsForSale      float64                              `json:"accruals_for_sale"`
	Amount               float64                              `json:"amount"`
	DeliveryCharge       float64                              `json:"delivery_charge"`
	Items                []ListResponseResultOperationItem    `json:"items"`
	OperationDate        string                               `json:"operation_date"`
	OperationID          int64                                `json:"operation_id"`
	OperationType        string                               `json:"operation_type"`
	OperationTypeName    string                               `json:"operation_type_name"`
	Posting              ListResponseResultOperationPosting   `json:"posting"`
	ReturnDeliveryCharge float64                              `json:"return_delivery_charge"`
	SaleCommission       float64                              `json:"sale_commission"`
	Services             []ListResponseResultOperationService `json:"services"`
	Type                 ListResponseResultOperationType      `json:"type"`
}

type ListResponseResult struct {
	Operations []ListResponseResultOperation `json:"operations"`
	PageCount  int64                         `json:"page_count"`
	RowCount   int64                         `json:"row_count"`
}

type ListResponse struct {
	Result ListResponseResult `json:"result"`
}
