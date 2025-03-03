package v1

//go:generate go run github.com/abice/go-enum -f=$GOFILE --marshal --names

// ListResponseResultStatus
// ENUM(new=1, created, disabled, blocked, disabled_due_to_limit, error)
type ListResponseResultStatus int32
