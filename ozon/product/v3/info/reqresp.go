package info

type StocksRequestFilter struct {
	OfferID    []string                      `json:"offer_id"`
	ProductID  []string                      `json:"product_id"`
	Visibility StocksRequestFilterVisibility `json:"visibility"`
}

type StocksRequest struct {
	Filter StocksRequestFilter `json:"filter"`
	LastID string              `json:"last_id"`
	Limit  int64               `json:"limit"`
}

type StocksResponseResultItemStock struct {
	Present  int32  `json:"present"`
	Reserved int32  `json:"reserved"`
	Type     string `json:"type"`
}

type StocksResponseResultItem struct {
	OfferID   string                          `json:"offer_id"`
	ProductID int64                           `json:"product_id"`
	Stocks    []StocksResponseResultItemStock `json:"stocks"`
}

type StocksResponseResult struct {
	Items  []StocksResponseResultItem `json:"items"`
	LastID string                     `json:"last_id"`
	Total  int32                      `json:"total"`
}

type StocksResponse struct {
	Result StocksResponseResult `json:"result"`
}

type ListRequest struct {
	OfferID   []string `json:"offer_id"`
	ProductID []string `json:"product_id"`
	SKU       []string `json:"sku"`
}

type ListResponseItemCommission struct {
	DeliveryAmount float64 `json:"delivery_amount"`
	Percent        float64 `json:"percent"`
	ReturnAmount   float64 `json:"return_amount"`
	SaleSchema     string  `json:"sale_schema"`
	Value          float64 `json:"value"`
}

type ListResponseItemErrorTextsParam struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ListResponseItemErrorTexts struct {
	AttributeName    string                            `json:"attribute_name"`
	Description      string                            `json:"description"`
	HintCode         string                            `json:"hint_code"`
	Message          string                            `json:"message"`
	Params           []ListResponseItemErrorTextsParam `json:"params"`
	ShortDescription string                            `json:"short_description"`
}

type ListResponseItemError struct {
	AttributeID int64                      `json:"attribute_id"`
	Code        string                     `json:"code"`
	Field       string                     `json:"field"`
	Level       ListResponseItemErrorLevel `json:"level"`
	State       string                     `json:"state"`
	Texts       ListResponseItemErrorTexts `json:"texts"`
}

type ListResponseItemModelInfo struct {
	Count   int64 `json:"count"`
	ModelID int64 `json:"model_id"`
}

type ListResponseItemPriceIndexesExternalIndexData struct {
	MinimalPrice         string  `json:"minimal_price"`
	MinimalPriceCurrency string  `json:"minimal_price_currency"`
	PriceIndexValue      float64 `json:"price_index_value"`
}

type ListResponseItemPriceIndexesOzonIndexData struct {
	MinimalPrice         string  `json:"minimal_price"`
	MinimalPriceCurrency string  `json:"minimal_price_currency"`
	PriceIndexValue      float64 `json:"price_index_value"`
}

type ListResponseItemPriceIndexesSelfMarketplacesIndexData struct {
	MinimalPrice         string  `json:"minimal_price"`
	MinimalPriceCurrency string  `json:"minimal_price_currency"`
	PriceIndexValue      float64 `json:"price_index_value"`
}

type ListResponseItemPriceIndexes struct {
	ColorIndex                ListResponseItemPriceIndexesColorIndex                `json:"color_index"`
	ExternalIndexData         ListResponseItemPriceIndexesExternalIndexData         `json:"external_index_data"`
	OzonIndexData             ListResponseItemPriceIndexesOzonIndexData             `json:"ozon_index_data"`
	SelfMarketplacesIndexData ListResponseItemPriceIndexesSelfMarketplacesIndexData `json:"self_marketplaces_index_data"`
}

type ListResponseItemSource struct {
	CreatedAt    string                             `json:"created_at"`
	QuantCode    string                             `json:"quant_code"`
	ShipmentType ListResponseItemSourceShipmentType `json:"shipment_type"`
	SKU          int64                              `json:"sku"`
	Source       string                             `json:"source"`
}

type ListResponseItemStatuses struct {
	IsCreated         bool   `json:"is_created"`
	ModerateStatus    string `json:"moderate_status"`
	Status            string `json:"status"`
	StatusDescription string `json:"status_description"`
	StatusFailed      string `json:"status_failed"`
	StatusName        string `json:"status_name"`
	StatusTooltip     string `json:"status_tooltip"`
	StatusUpdatedAt   string `json:"status_updated_at"`
	ValidationStatus  string `json:"validation_status"`
}

type ListResponseItemStocksStock struct {
	Present  int32  `json:"present"`
	Reserved int32  `json:"reserved"`
	SKU      int64  `json:"sku"`
	Source   string `json:"source"`
}

type ListResponseItemStocks struct {
	HasStock bool                          `json:"has_stock"`
	Stocks   []ListResponseItemStocksStock `json:"stocks"`
}

type ListResponseItemVisibilityDetails struct {
	HasPrice bool `json:"has_price"`
	HasStock bool `json:"has_stock"`
}

type ListResponseItem struct {
	Barcodes              []string                          `json:"barcodes"`
	ColorImage            []string                          `json:"color_image"`
	Commissions           []ListResponseItemCommission      `json:"commissions"`
	CreatedAt             string                            `json:"created_at"`
	CurrencyCode          string                            `json:"currency_code"`
	DescriptionCategoryID int64                             `json:"description_category_id"`
	DiscountedFBOStocks   int32                             `json:"discounted_fbo_stocks"`
	Errors                []ListResponseItemError           `json:"errors"`
	HasDiscountedFboItem  bool                              `json:"has_discounted_fbo_item"`
	ID                    int64                             `json:"id"`
	Images                []string                          `json:"images"`
	Images360             []string                          `json:"images360"`
	IsArchived            bool                              `json:"is_archived"`
	IsAutoArchived        bool                              `json:"is_autoarchived"`
	IsDiscounted          bool                              `json:"is_discounted"`
	IsKGT                 bool                              `json:"is_kgt"`
	IsPrepaymentAllowed   bool                              `json:"is_prepayment_allowed"`
	IsSuper               bool                              `json:"is_super"`
	MarketingPrice        string                            `json:"marketing_price"`
	MinPrice              string                            `json:"min_price"`
	ModelInfo             ListResponseItemModelInfo         `json:"model_info"`
	Name                  string                            `json:"name"`
	OfferID               string                            `json:"offer_id"`
	OldPrice              string                            `json:"old_price"`
	Price                 string                            `json:"price"`
	PriceIndexes          ListResponseItemPriceIndexes      `json:"price_indexes"`
	PrimaryImage          []string                          `json:"primary_image"`
	Sources               []ListResponseItemSource          `json:"sources"`
	Statuses              ListResponseItemStatuses          `json:"statuses"`
	Stocks                ListResponseItemStocks            `json:"stocks"`
	TypeID                int64                             `json:"type_id"`
	UpdatedAt             string                            `json:"updated_at"`
	Vat                   string                            `json:"vat"`
	VisibilityDetails     ListResponseItemVisibilityDetails `json:"visibility_details"`
	VolumeWeight          float64                           `json:"volume_weight"`
}

type ListResponse struct {
	Items []ListResponseItem `json:"items"`
}
