package unigraphclient

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/emersonmacro/go-uniswap-subgraph-client/graphql"
)

func constructByIdQuery(id string, model modelFields, opts *RequestOptions) (*graphql.Request, error) {
	if opts == nil {
		opts = &RequestOptions{
			IncludeFields: []string{"*"},
		}
	}

	err := validateRequestOpts(ById, opts)
	if err != nil {
		return nil, err
	}

	if slices.Contains(opts.IncludeFields, "*") {
		fields, err := gatherModelFields(model, opts.ExcludeFields, true)
		if err != nil {
			return nil, err
		}
		opts.IncludeFields = fields
	}

	query, err := assembleQuery(ById, opts.IncludeFields, model, opts)
	if err != nil {
		return nil, err
	}

	req := graphql.NewRequest(query)
	req.Var("id", id)

	fmt.Println("*** DEBUG req.Query() ***")
	fmt.Println(req.Query())
	fmt.Println("*************")

	return req, nil
}

func constructListQuery(model modelFields, opts *RequestOptions) (*graphql.Request, error) {
	if opts == nil {
		opts = &RequestOptions{
			IncludeFields: []string{"*"},
		}
	}

	err := validateRequestOpts(List, opts)
	if err != nil {
		return nil, err
	}

	if slices.Contains(opts.IncludeFields, "*") {
		fields, err := gatherModelFields(model, opts.ExcludeFields, true)
		if err != nil {
			return nil, err
		}
		opts.IncludeFields = fields
	}

	query, err := assembleQuery(List, opts.IncludeFields, model, opts)
	if err != nil {
		return nil, err
	}

	req := graphql.NewRequest(query)
	req.Var("first", opts.First)
	req.Var("skip", opts.Skip)
	req.Var("orderBy", opts.OrderBy)
	req.Var("orderDir", opts.OrderDir)

	fmt.Println("*** DEBUG req.Query() ***")
	fmt.Println(req.Query())
	fmt.Println("*************")

	return req, nil
}

// assembles a properly formatted graphql query based on the provided includeFields
func assembleQuery(queryType QueryType, includeFields []string, model modelFields, opts *RequestOptions) (string, error) {
	var parts []string

	switch queryType {
	case ById:
		parts = []string{
			fmt.Sprintf("query %s($id: String!) {", model.name),
		}

		if opts.Block != 0 {
			parts = append(parts, fmt.Sprintf("	%s(id: $id, block: {number: %d}) {", model.name, opts.Block))
		} else {
			parts = append(parts, fmt.Sprintf("	%s(id: $id) {", model.name))
		}
	case List:
		parts = []string{
			fmt.Sprintf("query %s($first: Int!, $skip: Int!, $orderBy: String!, $orderDir: String!) {", pluralizeModelName(model.name)),
		}

		if opts.Block != 0 {
			parts = append(parts, fmt.Sprintf("	%s(first: $first, skip: $skip, orderBy: $orderBy, orderDirection: $orderDir, block: {number: %d}) {", pluralizeModelName(model.name), opts.Block))
		} else {
			parts = append(parts, fmt.Sprintf("	%s(first: $first, skip: $skip, orderBy: $orderBy, orderDirection: $orderDir) {", pluralizeModelName(model.name)))
		}
	default:
		return "", fmt.Errorf("unrecognized query type (%v)", queryType)
	}

	var refFieldMap map[string]fieldRefs = make(map[string]fieldRefs)

	for _, field := range opts.IncludeFields {
		isRef := false
		for k, v := range model.reference {
			prefix := fmt.Sprintf("%s.", k)
			if strings.HasPrefix(field, prefix) {
				isRef = true
				refModel, ok := modelMap[v]
				if !ok {
					return "", fmt.Errorf("reference field not found (%s)", k)
				}
				fieldWithoutPrefix := cutPrefix(field, prefix)
				fieldRef := refFieldMap[k]
				if strings.Contains(fieldWithoutPrefix, ".") {
					// TODO: validate reference to reference field
					fieldRef.refs = append(fieldRef.refs, fieldWithoutPrefix)
				} else {
					if !validateField(refModel, fieldWithoutPrefix) {
						return "", fmt.Errorf("unrecognized field given in opts.IncludeFields (%s)", field)
					}
					fieldRef.directs = append(fieldRef.directs, fieldWithoutPrefix)
				}
				refFieldMap[k] = fieldRef
				break
			}
		}
		if !isRef {
			if !validateField(model, field) {
				return "", fmt.Errorf("unrecognized field given in opts.IncludeFields (%s)", field)
			}
			parts = append(parts, fmt.Sprintf("		%s", field))
		}
	}

	for k, v := range refFieldMap {
		// TODO: add validation
		parts = append(parts, fmt.Sprintf("		%s {", k))
		if len(v.directs) > 0 {
			parts = append(parts, "			"+strings.Join(v.directs, "\n			"))
		}
		subRefMap := make(map[string][]string)
		for _, ref := range v.refs {
			fieldSplit := strings.Split(ref, ".")
			if len(fieldSplit) < 2 {
				return "", fmt.Errorf("error parsing reference field (%s)", ref)
			}
			parent := fieldSplit[0]
			child := fieldSplit[1]
			subRefMap[parent] = append(subRefMap[parent], child)
		}
		for subK, subV := range subRefMap {
			parts = append(parts, fmt.Sprintf("			%s {", subK), "				"+strings.Join(subV, "\n				"), "			}")
		}
		parts = append(parts, "		}")
	}

	parts = append(parts, "	}", "}")
	query := strings.Join(parts, "\n")

	return query, nil
}

// recursively gathers all fields for the given model, while honoring fields to be excluded
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
				fullRef := fmt.Sprintf("%s.%s", k, refField)
				if !slices.Contains(excludeFields, fullRef) {
					fields = append(fields, fullRef)
				}
			}
		} else {
			refWithId := fmt.Sprintf("%s.id", k)
			if !slices.Contains(excludeFields, refWithId) {
				fields = append(fields, refWithId)
			}
		}
	}
	return fields, nil
}

func validateField(model modelFields, field string) bool {
	return slices.Contains(model.direct, field)
}

func cutPrefix(s string, prefix string) string {
	cut, _ := strings.CutPrefix(s, prefix)
	return cut
}

func validateRequestOpts(queryType QueryType, opts *RequestOptions) error {
	if len(opts.IncludeFields) == 0 {
		opts.IncludeFields = []string{"*"}
	}

	if !slices.Contains(opts.IncludeFields, "*") && len(opts.ExcludeFields) > 0 {
		return errors.New("request options error: ExcludeFields can only be provided when IncludeFields is set to '*'")
	}

	switch queryType {
	case ById:
		if opts.First != 0 || opts.Skip != 0 || opts.OrderBy != "" || opts.OrderDir != "" {
			return errors.New("request options error: List query options (First, Skip, OrderBy, OrderDir) should not be provided for ById queries")
		}
	case List:
		if opts.First > 1000 {
			return errors.New("request options error: First is too large (must be <= 1000)")
		}
		if opts.First == 0 {
			opts.First = 100
		}
		if opts.OrderBy == "" {
			opts.OrderBy = "id"
		}
		if opts.OrderDir == "" {
			opts.OrderDir = "asc"
		}
		if opts.OrderDir != "asc" && opts.OrderDir != "desc" {
			return errors.New("request options error: 'asc' and 'desc' are the only valid options for OrderDir")
		}
	}

	return nil
}

func pluralizeModelName(name string) string {
	if name == "factory" {
		return "factories"
	}
	return fmt.Sprintf("%ss", name)
}
