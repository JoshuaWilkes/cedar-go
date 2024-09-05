package parsererror

import (
	"fmt"

	"github.com/cedar-policy/cedar-go/internal/ast"
)

type PositionalError struct {
	Position ast.Position
	Message  string
}

func (pe PositionalError) Error() string {
	return fmt.Sprintf("%v: %v", pe.Position, pe.Message)
}
