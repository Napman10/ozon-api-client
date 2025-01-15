package info

//go:generate go run github.com/abice/go-enum -f=$GOFILE --marshal --names

// PricesRequestFilterVisibility
// ENUM(ALL=0, VISIBLE, INVISIBLE, EMPTY_STOCK, NOT_MODERATED, MODERATED, DISABLED, STATE_FAILED, READY_TO_SUPPLY, VALIDATION_STATE_PENDING, VALIDATION_STATE_FAIL, VALIDATION_STATE_SUCCESS, TO_SUPPLY, IN_SALE, REMOVED_FROM_SALE, BANNED, OVERPRICED, CRITICALLY_OVERPRICED, EMPTY_BARCODE, BARCODE_EXISTS, QUARANTINE, ARCHIVED, OVERPRICED_WITH_STOCK, PARTIAL_APPROVED)
type PricesRequestFilterVisibility int32

// PricesResponseItemPriceCurrencyCode
// ENUM(RUB=1, BYN, KZT, EUR, USD, CNY)
type PricesResponseItemPriceCurrencyCode int32

// PricesResponseItemPriceIndexesColorIndex
// ENUM(WITHOUT_INDEX=0, GREEN, YELLOW, RED)
type PricesResponseItemPriceIndexesColorIndex int32
