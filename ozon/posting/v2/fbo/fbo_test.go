package fbo_test

import (
	"bytes"
	"context"
	"github.com/andmetoo/ozon-api-client/internal/auth"
	"github.com/andmetoo/ozon-api-client/internal/test"
	"github.com/andmetoo/ozon-api-client/ozon/posting/v2/fbo"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestList_Success(t *testing.T) {
	c := fbo.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v2/posting/fbo/list", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"dir":"ASC","filter":{"since":"2021-09-01T00:00:00Z","status":"","to":"2021-11-17T10:44:12.828Z"},"limit":5,"offset":0,"translit":true,"with":{"analytics_data":true,"financial_data":true}}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							  "result": [
								{
								  "order_id": 354680487,
								  "order_number": "16965409-0014",
								  "posting_number": "16965409-0014-1",
								  "status": "delivered",
								  "cancel_reason_id": 0,
								  "created_at": "2021-09-01T00:23:45.607000Z",
								  "in_process_at": "2021-09-01T00:25:30.120000Z",
								  "products": [
									{
									  "sku": 160249683,
									  "name": "Так говорил Омар Хайям. Жизнеописание. Афоризмы и рубайят. Классика в словах и картинках",
									  "quantity": 1,
									  "offer_id": "978-5-906864-56-7",
									  "price": "81.00",
									  "digital_codes": [],
									  "currency_code": "RUB"
									}
								  ],
								  "analytics_data": {
									"region": "",
									"city": "",
									"delivery_type": "PVZ",
									"is_premium": false,
									"payment_type_group_name": "Карты оплаты",
									"warehouse_id": 17717042026000,
									"warehouse_name": "РОСТОВ-НА-ДОНУ_РФЦ",
									"is_legal": false
								  },
								  "financial_data": {
									"products": [
									  {
										"commission_amount": 12.15,
										"commission_percent": 15,
										"payout": 68.85,
										"product_id": 160249683,
										"currency_code": "RUB",
										"old_price": 115,
										"price": 81,
										"total_discount_value": 34,
										"total_discount_percent": 29.57,
										"actions": [
										  "Системная виртуальная скидка селлера"
										],
										"picking": null,
										"quantity": 0,
										"client_price": ""
									  }
									]
								  },
								  "additional_data": []
								}
							  ]
							}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v2/posting/fbo",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.List(context.Background(), &fbo.ListRequest{
		Dir: "ASC",
		Filter: fbo.ListRequestFilter{
			Since:  time.Date(2021, 9, 1, 0, 0, 0, 0, time.UTC),
			Status: "",
			To:     time.Date(2021, 11, 17, 10, 44, 12, 828000000, time.UTC),
		},
		Limit:    5,
		Offset:   0,
		Translit: true,
		With: fbo.ListRequestWith{
			AnalyticsData: true,
			FinancialData: true,
		},
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &fbo.ListResponse{
		Result: []fbo.ListResponseResult{
			{
				OrderID:        354680487,
				OrderNumber:    "16965409-0014",
				PostingNumber:  "16965409-0014-1",
				Status:         "delivered",
				CancelReasonID: 0,
				CreatedAt:      time.Date(2021, 9, 1, 0, 23, 45, 607000000, time.UTC),
				InProcessAt:    time.Date(2021, 9, 1, 0, 25, 30, 120000000, time.UTC),
				Products: []fbo.ListResponseResultProduct{
					{
						SKU:          160249683,
						Name:         "Так говорил Омар Хайям. Жизнеописание. Афоризмы и рубайят. Классика в словах и картинках",
						Quantity:     1,
						OfferID:      "978-5-906864-56-7",
						Price:        "81.00",
						DigitalCodes: []string{},
						CurrencyCode: "RUB",
					},
				},
				AnalyticsData: fbo.ListResponseResultAnalyticsData{
					Region:               "",
					City:                 "",
					DeliveryType:         "PVZ",
					IsPremium:            false,
					PaymentTypeGroupName: "Карты оплаты",
					WarehouseID:          17717042026000,
					WarehouseName:        "РОСТОВ-НА-ДОНУ_РФЦ",
					IsLegal:              false,
				},
				FinancialData: fbo.ListResponseResultFinancialData{
					Products: []fbo.ListResponseResultFinancialDataProduct{
						{
							CommissionAmount:     12.15,
							CommissionPercent:    15,
							Payout:               68.85,
							ProductID:            160249683,
							CurrencyCode:         "RUB",
							OldPrice:             115,
							Price:                81,
							TotalDiscountValue:   34,
							TotalDiscountPercent: 29.57,
							Actions:              []string{"Системная виртуальная скидка селлера"},
							Picking:              nil,
							Quantity:             0,
							ClientPrice:          "",
						},
					},
				},
				AdditionalData: []fbo.ListResponseResultAdditionalData{},
			},
		},
	}, resp)
}
