package roamer

import (
	"context"

	"github.com/pkg/errors"

	roamerError "github.com/SLIpros/roamer/err"
)

// ContextKey context key.
type ContextKey uint8

const (
	// ContextKeyParsedData  is a parsed data.
	ContextKeyParsedData ContextKey = iota + 1
	// ContextKeyParsingError is a parsing error.
	ContextKeyParsingError
)

// ParsedDataFromContext return parsed data from context.
func ParsedDataFromContext[T any](ctx context.Context, ptr *T) error {
	if ptr == nil {
		return errors.WithMessage(roamerError.NilValue, "context")
	}

	if err, ok := ctx.Value(ContextKeyParsingError).(error); ok {
		return errors.WithStack(err)
	}

	v, ok := ctx.Value(ContextKeyParsedData).(*T)
	if !ok {
		return errors.WithStack(roamerError.NoData)
	}

	*ptr = *v
	return nil
}

// ContextWithParsedData returns a context with parsed data.
func ContextWithParsedData(ctx context.Context, data any) context.Context {
	return context.WithValue(ctx, ContextKeyParsedData, data)
}

// contextWithParsingError returns a context with parsing error.
func contextWithParsingError(ctx context.Context, err error) context.Context {
	return context.WithValue(ctx, ContextKeyParsingError, err)
}
