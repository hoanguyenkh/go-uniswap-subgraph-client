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
	if len(opts.IncludeFields) == 0 && len(opts.ExcludeFields) == 0 {
		opts.IncludeFields = []string{"*"}
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

	var refFieldMap map[string]fieldRefs = make(map[string]fieldRefs)

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
				fieldWithoutPrefix := cutPrefix(field, prefix)
				fieldRef := refFieldMap[k]
				if strings.Contains(fieldWithoutPrefix, ".") {
					// TODO: validate reference to reference field
					fieldRef.refs = append(fieldRef.refs, fieldWithoutPrefix)
				} else {
					if !validateField(refModel, fieldWithoutPrefix) {
						return nil, fmt.Errorf("unrecognized field given in opts.IncludeFields (%s)", field)
					}
					fieldRef.directs = append(fieldRef.directs, fieldWithoutPrefix)
				}
				refFieldMap[k] = fieldRef
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
		// TODO: add validation
		parts = append(parts, fmt.Sprintf("		%s {", k))
		if len(v.directs) > 0 {
			parts = append(parts, "			"+strings.Join(v.directs, "\n			"))
		}
		subRefMap := make(map[string][]string)
		for _, ref := range v.refs {
			fieldSplit := strings.Split(ref, ".")
			if len(fieldSplit) < 2 {
				return nil, fmt.Errorf("error parsing reference field (%s)", ref)
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

	req := graphql.NewRequest(query)
	req.Var("id", id)

	fmt.Println("*** DEBUG req.Query() ***")
	fmt.Println(req.Query())
	fmt.Println("*************")

	return req, nil
}

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

func validateField(model modelFields, field string) bool {
	return slices.Contains(model.direct, field)
}

func cutPrefix(s string, prefix string) string {
	cut, _ := strings.CutPrefix(s, prefix)
	return cut
}

type fieldRefs struct {
	directs []string
	refs    []string
}
