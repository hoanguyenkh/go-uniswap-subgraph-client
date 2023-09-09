package unigraphclient

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/emersonmacro/go-uniswap-subgraph-client/graphql"
)

func constructPoolByIdQuery(id string, opts *RequestOptions) (*graphql.Request, error) {
	if !slices.Contains(opts.IncludeFields, "*") && len(opts.ExcludeFields) > 0 {
		return nil, errors.New("request options error: ExcludeFields can only be provided when IncludeFields is set to '*'")
	}

	if slices.Contains(opts.IncludeFields, "*") {
		opts.IncludeFields = gatherAllPoolByIdFields(opts.ExcludeFields)
	}

	var token0Fields, token1Fields []string = []string{}, []string{}
	parts := []string{
		"query pool($poolAddress: String!) {",
		"	pool(id: $poolAddress) {",
	}

	for _, field := range opts.IncludeFields {
		switch {
		case strings.HasPrefix(field, "token0."):
			tokenField := cutPrefix(field, "token0.")
			if !validateTokenField(tokenField) {
				return nil, fmt.Errorf("unrecognized Token0 field in GetPoolById query: %s", field)
			}
			token0Fields = append(token0Fields, tokenField)
		case strings.HasPrefix(field, "token1."):
			tokenField := cutPrefix(field, "token1.")
			if !validateTokenField(tokenField) {
				return nil, fmt.Errorf("unrecognized Token1 field in GetPoolById query: %s", field)
			}
			token1Fields = append(token1Fields, tokenField)
		default:
			if !validatePoolField(field) {
				return nil, fmt.Errorf("unrecognized Pool field in GetPoolById query: %s", field)
			}
			parts = append(parts, fmt.Sprintf("		%s", field))
		}
	}

	if len(token0Fields) > 0 {
		parts = append(parts, "		token0 {", "			"+strings.Join(token0Fields, "\n			"), "		}")
	}

	if len(token1Fields) > 0 {
		parts = append(parts, "		token1 {", "			"+strings.Join(token1Fields, "\n			"), "		}")
	}

	parts = append(parts, "	}", "}")
	query := strings.Join(parts, "\n")

	req := graphql.NewRequest(query)
	req.Var("poolAddress", id)

	return req, nil
}

func gatherAllPoolByIdFields(excludeFields []string) []string {
	fields := []string{}
	for _, field := range PoolFields {
		if !slices.Contains(excludeFields, field) {
			fields = append(fields, field)
		}
	}
	for _, field := range TokenFields {
		token0Field := fmt.Sprintf("token0.%s", field)
		if !slices.Contains(excludeFields, token0Field) {
			fields = append(fields, token0Field)
		}
		token1Field := fmt.Sprintf("token1.%s", field)
		if !slices.Contains(excludeFields, token1Field) {
			fields = append(fields, token1Field)
		}
	}
	return fields
}

func validatePoolField(field string) bool {
	return slices.Contains(PoolFields, field)
}

func validateTokenField(field string) bool {
	return slices.Contains(TokenFields, field)
}

func cutPrefix(s string, prefix string) string {
	cut, _ := strings.CutPrefix(s, prefix)
	return cut
}
