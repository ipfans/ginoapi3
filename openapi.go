package ginoapi3

import "github.com/getkin/kin-openapi/openapi3"

func NewInfo(title, version string) *openapi3.Info {
	return &openapi3.Info{
		Title:   title,
		Version: version,
	}
}

type OperationOption func(*openapi3.Operation)

type Operation func[T, V any](t T, v V) *openapi3.Operation