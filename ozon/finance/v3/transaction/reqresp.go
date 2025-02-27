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
	DeliverySchema string `json:"delivery_schema"`
	OrderDate      string `json:"order_date"`
	PostingNumber  string `json:"posting_number"`
	WarehouseID    int64  `json:"warehouse_id"`
}

type ListResponseResultOperationItem struct {
	Name string `json:"name"`
	SKU  int64  `json:"sku"`
}

type ListResponseResultOperation struct {
	OperationID          int64                              `json:"operation_id"`
	OperationType        string                             `json:"operation_type"`
	OperationDate        string                             `json:"operation_date"`
	OperationTypeName    string                             `json:"operation_type_name"`
	DeliveryCharge       int64                              `json:"delivery_charge"`
	ReturnDeliveryCharge int64                              `json:"return_delivery_charge"`
	AccrualsForSale      int64                              `json:"accruals_for_sale"`
	SaleCommission       int64                              `json:"sale_commission"`
	Amount               float64                            `json:"amount"`
	Type                 string                             `json:"type"`
	Posting              ListResponseResultOperationPosting `json:"posting"`
	Items                []ListResponseResultOperationItem  `json:"items"`
}

type ListResponseResult struct {
	Operations []ListResponseResultOperation `json:"operations"`
	PageCount  int64                         `json:"page_count"`
	RowCount   int64                         `json:"row_count"`
}

type ListResponse struct {
	Result ListResponseResult `json:"result"`
}
