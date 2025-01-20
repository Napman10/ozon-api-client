package info_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/andmetoo/ozon-api-client/ozon/product/v3/info"

	"github.com/andmetoo/ozon-api-client/internal/auth"
	"github.com/andmetoo/ozon-api-client/internal/test"
	"github.com/stretchr/testify/require"
)

func TestStocks_Success(t *testing.T) {
	c := info.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v3/product/info/stocks", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"filter":{"offer_id":["136834"],"product_id":["214887921"],"visibility":"ALL"},"last_id":"","limit":100}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"result": {
								"items": [
									{
										"product_id": 214887921,
										"offer_id": "136834",
										"stocks": [
											{
												"type": "fbs",
												"present": 170,
												"reserved": 0
											},
											{
												"type": "fbo",
												"present": 0,
												"reserved": 0
											}
										]
									}
								],
								"total": 1,
								"last_id": "anVsbA=="
							}
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v3/product/info",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Stocks(context.Background(), &info.StocksRequest{
		Filter: info.StocksRequestFilter{
			OfferID: []string{
				"136834",
			},
			ProductID: []string{
				"214887921",
			},
			Visibility: info.StocksRequestFilterVisibilityALL,
		},
		LastID: "",
		Limit:  100,
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &info.StocksResponse{
		Result: info.StocksResponseResult{
			Items: []info.StocksResponseResultItem{
				{
					ProductID: 214887921,
					OfferID:   "136834",
					Stocks: []info.StocksResponseResultItemStock{
						{
							Type:     "fbs",
							Present:  170,
							Reserved: 0,
						},
						{
							Type:     "fbo",
							Present:  0,
							Reserved: 0,
						},
					},
				},
			},
			Total:  1,
			LastID: "anVsbA==",
		},
	}, resp)
}

func TestList_Success(t *testing.T) {
	c := info.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v3/product/info/list", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"offer_id":["string"],"product_id":["string"],"sku":["string"]}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"items": [
								{
									"barcodes": [
										"string"
									],
									"color_image": [
										"string"
									],
									"commissions": [
										{
											"delivery_amount": 0,
											"percent": 0,
											"return_amount": 0,
											"sale_schema": "string",
											"value": 0
										}
									],
									"created_at": "2019-08-24T14:15:22Z",
									"currency_code": "string",
									"description_category_id": 0,
									"discounted_fbo_stocks": 0,
									"errors": [
										{
											"attribute_id": 0,
											"code": "string",
											"field": "string",
											"level": "ERROR_LEVEL_UNSPECIFIED",
											"state": "string",
											"texts": {
												"attribute_name": "string",
												"description": "string",
												"hint_code": "string",
												"message": "string",
												"params": [
													{
														"name": "string",
														"value": "string"
													}
												],
												"short_description": "string"
											}
										}
									],
									"has_discounted_fbo_item": true,
									"id": 0,
									"images": [
										"string"
									],
									"images360": [
										"string"
									],
									"is_archived": true,
									"is_autoarchived": true,
									"is_discounted": true,
									"is_kgt": true,
									"is_prepayment_allowed": true,
									"is_super": true,
									"marketing_price": "string",
									"min_price": "string",
									"model_info": {
										"count": 0,
										"model_id": 0
									},
									"name": "string",
									"offer_id": "string",
									"old_price": "string",
									"price": "string",
									"price_indexes": {
										"color_index": "COLOR_INDEX_UNSPECIFIED",
										"external_index_data": {
											"minimal_price": "string",
											"minimal_price_currency": "string",
											"price_index_value": 0
										},
										"ozon_index_data": {
											"minimal_price": "string",
											"minimal_price_currency": "string",
											"price_index_value": 0
										},
										"self_marketplaces_index_data": {
											"minimal_price": "string",
											"minimal_price_currency": "string",
											"price_index_value": 0
										}
									},
									"primary_image": [
										"string"
									],
									"sources": [
										{
											"created_at": "2019-08-24T14:15:22Z",
											"quant_code": "string",
											"shipment_type": "SHIPMENT_TYPE_UNSPECIFIED",
											"sku": 0,
											"source": "string"
										}
									],
									"statuses": {
										"is_created": true,
										"moderate_status": "string",
										"status": "string",
										"status_description": "string",
										"status_failed": "string",
										"status_name": "string",
										"status_tooltip": "string",
										"status_updated_at": "2019-08-24T14:15:22Z",
										"validation_status": "string"
									},
									"stocks": {
										"has_stock": true,
										"stocks": [
											{
												"present": 0,
												"reserved": 0,
												"sku": 0,
												"source": "string"
											}
										]
									},
									"type_id": 0,
									"updated_at": "2019-08-24T14:15:22Z",
									"vat": "string",
									"visibility_details": {
										"has_price": true,
										"has_stock": true
									},
									"volume_weight": 0
								}
							]
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v3/product/info",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.List(context.Background(), &info.ListRequest{
		OfferID: []string{
			"string",
		},
		ProductID: []string{
			"string",
		},
		SKU: []string{
			"string",
		},
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &info.ListResponse{
		Items: []info.ListResponseItem{
			{
				Barcodes: []string{
					"string",
				},
				ColorImage: []string{
					"string",
				},
				Commissions: []info.ListResponseItemCommission{
					{
						DeliveryAmount: 0,
						Percent:        0,
						ReturnAmount:   0,
						SaleSchema:     "string",
						Value:          0,
					},
				},
				CreatedAt:             "2019-08-24T14:15:22Z",
				CurrencyCode:          "string",
				DescriptionCategoryID: 0,
				DiscountedFBOStocks:   0,
				Errors: []info.ListResponseItemError{
					{
						AttributeID: 0,
						Code:        "string",
						Field:       "string",
						Level:       info.ListResponseItemErrorLevelERRORLEVELUNSPECIFIED,
						State:       "string",
						Texts: info.ListResponseItemErrorTexts{
							AttributeName: "string",
							Description:   "string",
							HintCode:      "string",
							Message:       "string",
							Params: []info.ListResponseItemErrorTextsParam{
								{
									Name:  "string",
									Value: "string",
								},
							},
							ShortDescription: "string",
						},
					},
				},
				HasDiscountedFboItem: true,
				ID:                   0,
				Images: []string{
					"string",
				},
				Images360: []string{
					"string",
				},
				IsArchived:          true,
				IsAutoArchived:      true,
				IsDiscounted:        true,
				IsKGT:               true,
				IsPrepaymentAllowed: true,
				IsSuper:             true,
				MarketingPrice:      "string",
				MinPrice:            "string",
				ModelInfo: info.ListResponseItemModelInfo{
					Count:   0,
					ModelID: 0,
				},
				Name:     "string",
				OfferID:  "string",
				OldPrice: "string",
				Price:    "string",
				PriceIndexes: info.ListResponseItemPriceIndexes{
					ColorIndex: info.ListResponseItemPriceIndexesColorIndexCOLORINDEXUNSPECIFIED,
					ExternalIndexData: info.ListResponseItemPriceIndexesExternalIndexData{
						MinimalPrice:         "string",
						MinimalPriceCurrency: "string",
						PriceIndexValue:      0,
					},
					OzonIndexData: info.ListResponseItemPriceIndexesOzonIndexData{
						MinimalPrice:         "string",
						MinimalPriceCurrency: "string",
						PriceIndexValue:      0,
					},
					SelfMarketplacesIndexData: info.ListResponseItemPriceIndexesSelfMarketplacesIndexData{
						MinimalPrice:         "string",
						MinimalPriceCurrency: "string",
						PriceIndexValue:      0,
					},
				},
				PrimaryImage: []string{
					"string",
				},
				Sources: []info.ListResponseItemSource{
					{
						CreatedAt:    "2019-08-24T14:15:22Z",
						QuantCode:    "string",
						ShipmentType: info.ListResponseItemSourceShipmentTypeSHIPMENTTYPEUNSPECIFIED,
						SKU:          0,
						Source:       "string",
					},
				},
				Statuses: info.ListResponseItemStatuses{
					IsCreated:         true,
					ModerateStatus:    "string",
					Status:            "string",
					StatusDescription: "string",
					StatusFailed:      "string",
					StatusName:        "string",
					StatusTooltip:     "string",
					StatusUpdatedAt:   "2019-08-24T14:15:22Z",
					ValidationStatus:  "string",
				},
				Stocks: info.ListResponseItemStocks{
					HasStock: true,
					Stocks: []info.ListResponseItemStocksStock{
						{
							Present:  0,
							Reserved: 0,
							SKU:      0,
							Source:   "string",
						},
					},
				},
				TypeID:    0,
				UpdatedAt: "2019-08-24T14:15:22Z",
				Vat:       "string",
				VisibilityDetails: info.ListResponseItemVisibilityDetails{
					HasPrice: true,
					HasStock: true,
				},
				VolumeWeight: 0,
			},
		},
	}, resp)
}
