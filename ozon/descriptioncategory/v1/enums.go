package v1

//go:generate go-enum -f=$GOFILE --marshal --names

// Language
// ENUM(DEFAULT=0, RU, EN, TR, ZH_HANS)
type Language int32
