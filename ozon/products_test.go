package ozon

import (
	"net/http"
	"testing"

	core "github.com/diphantxm/ozon-api-client"
)

func TestGetStocksInfo(t *testing.T) {
	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetStocksInfoParams
		response   string
	}{
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetStocksInfoParams{
				Limit:  100,
				LastId: "",
				Filter: GetStocksInfoFilter{
					OfferId:    "136834",
					ProductId:  214887921,
					Visibility: "ALL",
				},
			},
			`{
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
			}`,
		},
		{
			http.StatusBadRequest,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetStocksInfoParams{
				Limit:  100,
				LastId: "",
				Filter: GetStocksInfoFilter{
					OfferId:    "136834",
					ProductId:  214887921,
					Visibility: "ALL",
				},
			},
			`{
				"code": 0,
				"details": [
				  {
					"typeUrl": "string",
					"value": "string"
				  }
				],
				"message": "string"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.GetStocksInfo(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetProductDetails(t *testing.T) {
	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetProductDetailsParams
		response   string
	}{
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetProductDetailsParams{
				ProductId: 137208233,
			},
			`{
				"result": {
				  "id": 137208233,
				  "name": "Комплект защитных плёнок для X3 NFC. Темный хлопок",
				  "offer_id": "143210586",
				  "barcode": "",
				  "barcodes": [
					"2335900005",
					"7533900005"
				  ],
				  "buybox_price": "",
				  "category_id": 17038062,
				  "created_at": "2021-10-21T15:48:03.529178Z",
				  "images": [
					"https://cdn1.ozone.ru/s3/multimedia-5/6088931525.jpg",
					"https://cdn1.ozone.ru/s3/multimedia-p/6088915813.jpg"
				  ],
				  "has_discounted_item": true,
				  "is_discounted": true,
				  "discounted_stocks": {
					"coming": 0,
					"present": 0,
					"reserved": 0
				  },
				  "currency_code": "RUB",
				  "marketing_price": "",
				  "min_price": "",
				  "old_price": "",
				  "premium_price": "",
				  "price": "590.0000",
				  "recommended_price": "",
				  "sources": [
					{
					  "is_enabled": true,
					  "sku": 522759607,
					  "source": "fbo"
					},
					{
					  "is_enabled": true,
					  "sku": 522759608,
					  "source": "fbs"
					}
				  ],
				  "stocks": {
					"coming": 0,
					"present": 0,
					"reserved": 0
				  },
				  "errors": [],
				  "updated_at": "2023-02-09T06:46:44.152Z",
				  "vat": "0.0",
				  "visible": false,
				  "visibility_details": {
					"has_price": true,
					"has_stock": false,
					"active_product": false
				  },
				  "price_index": "0.00",
				  "commissions": [],
				  "volume_weight": 0.1,
				  "is_prepayment": false,
				  "is_prepayment_allowed": true,
				  "images360": [],
				  "is_kgt": false,
				  "color_image": "",
				  "primary_image": "https://cdn1.ozone.ru/s3/multimedia-p/6088931545.jpg",
				  "status": {
					"state": "imported",
					"state_failed": "imported",
					"moderate_status": "",
					"decline_reasons": [],
					"validation_state": "pending",
					"state_name": "Не продается",
					"state_description": "Не создан",
					"is_failed": true,
					"is_created": false,
					"state_tooltip": "",
					"item_errors": [
					  {
						"code": "error_attribute_values_empty",
						"field": "attribute",
						"attribute_id": 9048,
						"state": "imported",
						"level": "error",
						"description": "Не заполнен обязательный атрибут. Иногда мы обновляем обязательные атрибуты или добавляем новые. Отредактируйте товар или загрузите новый XLS-шаблон с актуальными атрибутами. ",
						"optional_description_elements": {},
						"attribute_name": "Название модели"
					  },
					  {
						"code": "error_attribute_values_empty",
						"field": "attribute",
						"attribute_id": 5076,
						"state": "imported",
						"level": "error",
						"description": "Не заполнен обязательный атрибут. Иногда мы обновляем обязательные атрибуты или добавляем новые. Отредактируйте товар или загрузите новый XLS-шаблон с актуальными атрибутами. ",
						"optional_description_elements": {},
						"attribute_name": "Рекомендовано для"
					  },
					  {
						"code": "error_attribute_values_empty",
						"field": "attribute",
						"attribute_id": 8229,
						"state": "imported",
						"level": "error",
						"description": "Не заполнен обязательный атрибут. Иногда мы обновляем обязательные атрибуты или добавляем новые. Отредактируйте товар или загрузите новый XLS-шаблон с актуальными атрибутами. ",
						"optional_description_elements": {},
						"attribute_name": "Тип"
					  },
					  {
						"code": "error_attribute_values_empty",
						"field": "attribute",
						"attribute_id": 85,
						"state": "imported",
						"level": "error",
						"description": "Не заполнен обязательный атрибут. Иногда мы обновляем обязательные атрибуты или добавляем новые. Отредактируйте товар или загрузите новый XLS-шаблон с актуальными атрибутами. ",
						"optional_description_elements": {},
						"attribute_name": "Бренд"
					  }
					],
					"state_updated_at": "2021-10-21T15:48:03.927309Z"
				  }
				}
			  }`,
		},
		{
			http.StatusBadRequest,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetProductDetailsParams{
				ProductId: 137208233,
			},
			`{
				"code": 0,
				"details": [
				  {
					"typeUrl": "string",
					"value": "string"
				  }
				],
				"message": "string"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.GetProductDetails(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestUpdateStocks(t *testing.T) {
	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *UpdateStocksParams
		response   string
	}{
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&UpdateStocksParams{
				Stocks: []UpdateStocksStock{
					{
						OfferId:   "PG-2404С1",
						ProductId: 55946,
						Stock:     4,
					},
				},
			},
			`{
				"result": [
				  {
					"product_id": 55946,
					"offer_id": "PG-2404С1",
					"updated": true,
					"errors": []
				  }
				]
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.UpdateStocks(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestStocksInSellersWarehouse(t *testing.T) {
	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *StocksInSellersWarehouseParams
		response   string
	}{
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&StocksInSellersWarehouseParams{
				FBSSKU: []string{"123"},
			},
			`{
				"result": [
				  {
					"fbs_sku": 12,
					"present": 34,
					"product_id": 548761,
					"reserved": 5,
					"warehouse_id": 156778,
					"warehouse_name": "something"
				  }
				]
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.StocksInSellersWarehouse(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestUpdatePrices(t *testing.T) {
	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *UpdatePricesParams
		response   string
	}{
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&UpdatePricesParams{
				Prices: []UpdatePricesPrice{
					{
						AutoActionEnabled: "UNKNOWN",
						CurrencyCode:      "RUB",
						MinPrice:          "800",
						OldPrice:          "0",
						Price:             "1448",
						ProductId:         1386,
					},
				},
			},
			`{
				"result": [
				  {
					"product_id": 1386,
					"offer_id": "PH8865",
					"updated": true,
					"errors": []
				  }
				]
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.UpdatePrices(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestUpdateQuantityStockProducts(t *testing.T) {
	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *UpdateQuantityStockProductsParams
		response   string
	}{
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&UpdateQuantityStockProductsParams{
				Stocks: []UpdateQuantityStockProductsStock{
					{
						OfferId:     "PH11042",
						ProductId:   313455276,
						Stock:       100,
						WarehouseId: 22142605386000,
					},
				},
			},
			`{
				"result": [
				  {
					"warehouse_id": 22142605386000,
					"product_id": 118597312,
					"offer_id": "PH11042",
					"updated": true,
					"errors": []
				  }
				]
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.UpdateQuantityStockProducts(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestCreateOrUpdateProduct(t *testing.T) {
	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *CreateOrUpdateProductParams
		response   string
	}{
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&CreateOrUpdateProductParams{
				Items: []CreateOrUpdateProductItem{
					{
						Attributes: []CreateOrUpdateAttribute{
							{
								ComplexId: 0,
								Id:        5076,
								Values: []CreateOrUpdateAttributeValue{
									{
										DictionaryValueId: 971082156,
										Value:             "Стойка для акустической системы",
									},
								},
							},
							{
								ComplexId: 0,
								Id:        9048,
								Values: []CreateOrUpdateAttributeValue{
									{
										Value: "Комплект защитных плёнок для X3 NFC. Темный хлопок",
									},
								},
							},
							{
								ComplexId: 0,
								Id:        8229,
								Values: []CreateOrUpdateAttributeValue{
									{
										DictionaryValueId: 95911,
										Value:             "Комплект защитных плёнок для X3 NFC. Темный хлопок",
									},
								},
							},
							{
								ComplexId: 0,
								Id:        85,
								Values: []CreateOrUpdateAttributeValue{
									{
										DictionaryValueId: 5060050,
										Value:             "Samsung",
									},
								},
							},
							{
								ComplexId: 0,
								Id:        10096,
								Values: []CreateOrUpdateAttributeValue{
									{
										DictionaryValueId: 61576,
										Value:             "серый",
									},
								},
							},
						},
						Barcode:       "112772873170",
						CategoryId:    17033876,
						CurrencyCode:  "RUB",
						Depth:         10,
						DimensionUnit: "mm",
						Height:        250,
						Name:          "Комплект защитных плёнок для X3 NFC. Темный хлопок",
						OfferId:       "143210608",
						OldPrice:      "1100",
						PremiumPrice:  "900",
						Price:         "1000",
						VAT:           "0.1",
						Weight:        100,
						WeightUnit:    "g",
						Width:         150,
					},
				},
			},
			`{
				"result": {
				  "task_id": 172549793
				}
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.CreateOrUpdateProduct(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetListOfProducts(t *testing.T) {
	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetListOfProductsParams
		response   string
	}{
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetListOfProductsParams{
				Filter: GetListOfProductsFilter{
					OfferId:    []string{"136748"},
					ProductId:  []int64{223681945},
					Visibility: "ALL",
				},
				LastId: "",
				Limit:  100,
			},
			`{
				"result": {
				  "items": [
					{
					  "product_id": 223681945,
					  "offer_id": "136748"
					}
				  ],
				  "total": 1,
				  "last_id": "bnVсbA=="
				}
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.GetListOfProducts(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}
