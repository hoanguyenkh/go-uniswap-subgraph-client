package unigraphclient

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/emersonmacro/go-uniswap-subgraph-client/graphql"
)

func gatherModelFields(model modelFields, excludeFields []string, populateRefs bool) ([]string, error) {
	fields := []string{}
	for _, field := range model.direct {
		if !slices.Contains(excludeFields, field) {
			fields = append(fields, field)
		}
	}
	for k, v := range model.reference {
		if populateRefs {
			refModel, ok := modelMap[v]
			if !ok {
				return nil, fmt.Errorf("reference field not found (%s)", k)
			}
			refFields, err := gatherModelFields(refModel, excludeFields, false)
			if err != nil {
				return nil, err
			}
			for _, refField := range refFields {
				fields = append(fields, fmt.Sprintf("%s.%s", k, refField))
			}
		} else {
			fields = append(fields, fmt.Sprintf("%s.id", k))
		}
	}
	return fields, nil
}

func constructByIdQuery(id string, model modelFields, opts *RequestOptions) (*graphql.Request, error) {
	if opts == nil {
		opts = &RequestOptions{}
	}

	if !slices.Contains(opts.IncludeFields, "*") && len(opts.ExcludeFields) > 0 {
		return nil, errors.New("request options error: ExcludeFields can only be provided when IncludeFields is set to '*'")
	}

	if slices.Contains(opts.IncludeFields, "*") {
		fields, err := gatherModelFields(model, opts.ExcludeFields, true)
		if err != nil {
			return nil, err
		}
		opts.IncludeFields = fields
	}

	parts := []string{
		fmt.Sprintf("query %s($id: String!) {", model.name),
	}

	if opts.Block != 0 {
		parts = append(parts, fmt.Sprintf("	%s(id: $id, block: {number: %d}) {", model.name, opts.Block))
	} else {
		parts = append(parts, fmt.Sprintf("	%s(id: $id) {", model.name))
	}

	var refFieldMap map[string][]string = make(map[string][]string)

	for _, field := range opts.IncludeFields {
		isRef := false
		for k, v := range model.reference {
			prefix := fmt.Sprintf("%s.", k)
			if strings.HasPrefix(field, prefix) {
				isRef = true
				refModel, ok := modelMap[v]
				if !ok {
					return nil, fmt.Errorf("reference field not found (%s)", k)
				}
				directFieldName := cutPrefix(field, prefix)
				if !validateField(refModel, directFieldName) {
					return nil, fmt.Errorf("unrecognized field given in opts.IncludeFields (%s)", field)
				}
				refFieldMap[k] = append(refFieldMap[k], directFieldName)
				break
			}
		}
		if !isRef {
			if !validateField(model, field) {
				return nil, fmt.Errorf("unrecognized field given in opts.IncludeFields (%s)", field)
			}
			parts = append(parts, fmt.Sprintf("		%s", field))
		}
	}

	for k, v := range refFieldMap {
		if len(v) > 0 {
			parts = append(parts, fmt.Sprintf("		%s {", k), "			"+strings.Join(v, "\n			"), "		}")
		}
	}

	parts = append(parts, "	}", "}")
	query := strings.Join(parts, "\n")

	req := graphql.NewRequest(query)
	req.Var("id", id)

	fmt.Println("*** DEBUG ***")
	fmt.Println(req.Query())
	fmt.Println("*************")

	return req, nil
}

func validateField(model modelFields, field string) bool {
	return slices.Contains(model.direct, field)
}

func cutPrefix(s string, prefix string) string {
	cut, _ := strings.CutPrefix(s, prefix)
	return cut
}

// func constructFactoryByIdQuery(id string, opts *RequestOptions) (*graphql.Request, error) {
// 	if !slices.Contains(opts.IncludeFields, "*") && len(opts.ExcludeFields) > 0 {
// 		return nil, errors.New("request options error: ExcludeFields can only be provided when IncludeFields is set to '*'")
// 	}

// 	if slices.Contains(opts.IncludeFields, "*") {
// 		opts.IncludeFields = gatherAllFactoryByIdFields(opts.ExcludeFields)
// 	}

// 	parts := []string{
// 		"query factory($factoryAddress: String!) {",
// 		"	factory(id: $factoryAddress) {",
// 	}

// 	for _, field := range opts.IncludeFields {
// 		parts = append(parts, fmt.Sprintf("		%s", field))
// 	}

// 	parts = append(parts, "	}", "}")
// 	query := strings.Join(parts, "\n")

// 	req := graphql.NewRequest(query)
// 	req.Var("factoryAddress", id)

// 	return req, nil
// }

// func constructPoolByIdQuery(id string, opts *RequestOptions) (*graphql.Request, error) {
// 	if !slices.Contains(opts.IncludeFields, "*") && len(opts.ExcludeFields) > 0 {
// 		return nil, errors.New("request options error: ExcludeFields can only be provided when IncludeFields is set to '*'")
// 	}

// 	if slices.Contains(opts.IncludeFields, "*") {
// 		opts.IncludeFields = gatherAllPoolByIdFields(opts.ExcludeFields)
// 	}

// 	var token0Fields, token1Fields []string = []string{}, []string{}
// 	parts := []string{
// 		"query pool($poolAddress: String!) {",
// 		"	pool(id: $poolAddress) {",
// 	}

// 	for _, field := range opts.IncludeFields {
// 		switch {
// 		case strings.HasPrefix(field, "token0."):
// 			tokenField := cutPrefix(field, "token0.")
// 			if !validateTokenField(tokenField) {
// 				return nil, fmt.Errorf("unrecognized Token0 field in GetPoolById query: %s", field)
// 			}
// 			token0Fields = append(token0Fields, tokenField)
// 		case strings.HasPrefix(field, "token1."):
// 			tokenField := cutPrefix(field, "token1.")
// 			if !validateTokenField(tokenField) {
// 				return nil, fmt.Errorf("unrecognized Token1 field in GetPoolById query: %s", field)
// 			}
// 			token1Fields = append(token1Fields, tokenField)
// 		default:
// 			if !validatePoolField(field) {
// 				return nil, fmt.Errorf("unrecognized Pool field in GetPoolById query: %s", field)
// 			}
// 			parts = append(parts, fmt.Sprintf("		%s", field))
// 		}
// 	}

// 	if len(token0Fields) > 0 {
// 		parts = append(parts, "		token0 {", "			"+strings.Join(token0Fields, "\n			"), "		}")
// 	}

// 	if len(token1Fields) > 0 {
// 		parts = append(parts, "		token1 {", "			"+strings.Join(token1Fields, "\n			"), "		}")
// 	}

// 	parts = append(parts, "	}", "}")
// 	query := strings.Join(parts, "\n")

// 	req := graphql.NewRequest(query)
// 	req.Var("poolAddress", id)

// 	return req, nil
// }

// func gatherAllFactoryByIdFields(excludeFields []string) []string {
// 	fields := []string{}
// 	for _, field := range FactoryFields {
// 		if !slices.Contains(excludeFields, field) {
// 			fields = append(fields, field)
// 		}
// 	}
// 	return fields
// }

// func gatherAllPoolByIdFields(excludeFields []string) []string {
// 	fields := []string{}
// 	for _, field := range PoolFields {
// 		if !slices.Contains(excludeFields, field) {
// 			fields = append(fields, field)
// 		}
// 	}
// 	for _, field := range TokenFields {
// 		token0Field := fmt.Sprintf("token0.%s", field)
// 		if !slices.Contains(excludeFields, token0Field) {
// 			fields = append(fields, token0Field)
// 		}
// 		token1Field := fmt.Sprintf("token1.%s", field)
// 		if !slices.Contains(excludeFields, token1Field) {
// 			fields = append(fields, token1Field)
// 		}
// 	}
// 	return fields
// }

// func validatePoolField(field string) bool {
// 	return slices.Contains(PoolFields, field)
// }

// func validateTokenField(field string) bool {
// 	return slices.Contains(TokenFields, field)
// }
