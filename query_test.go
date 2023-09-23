package unigraphclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGatherModelFields(t *testing.T) {
	tests := map[string]struct {
		model         modelFields
		excludeFields []string
		populateRefs  bool
		wantLen       int
		wantErr       bool
		wantErrMsg    string
	}{
		"when reference model is invalid": {
			model: modelFields{
				reference: map[string]string{
					"not found": "",
				},
			},
			populateRefs: true,
			wantErr:      true,
			wantErrMsg:   "reference field not found",
		},
		"when model has no references and excludeFields is empty": {
			model:   FactoryFields,
			wantLen: 13,
		},
		"when model has no references and excludeFields is not empty": {
			model:         FactoryFields,
			excludeFields: []string{"owner", "txCount"},
			wantLen:       11,
		},
		"when model has references, excludeFields is empty, and populateRefs is false": {
			model:   PoolFields,
			wantLen: 29,
		},
		"when model has references, excludeFields is not empty, and populateRefs is false": {
			model:         PoolFields,
			excludeFields: []string{"feeTier", "token0.id"},
			wantLen:       27,
		},
		"when model has references, excludeFields is empty, and populateRefs is true": {
			model:        PoolFields,
			populateRefs: true,
			wantLen:      59,
		},
		"when model has references, excludeFields is not empty, and populateRefs is true": {
			model:         PoolFields,
			excludeFields: []string{"feeTier", "token0.id"},
			populateRefs:  true,
			wantLen:       57,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := gatherModelFields(test.model, test.excludeFields, test.populateRefs)

			if test.wantErr {
				assert.NotNil(t, err)
				assert.Contains(t, err.Error(), test.wantErrMsg)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, test.wantLen, len(got))
			}
		})
	}
}

func TestValidateField(t *testing.T) {
	tests := map[string]struct {
		model modelFields
		field string
		want  bool
	}{
		"when field is in model fields": {
			model: PoolFields,
			field: "sqrtPrice",
			want:  true,
		},
		"when field is not in model fields": {
			model: PoolFields,
			field: "poolCount",
			want:  false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := validateField(test.model, test.field)

			assert.Equal(t, test.want, got)
		})
	}
}

func TestCutPrefix(t *testing.T) {
	tests := map[string]struct {
		s      string
		prefix string
		want   string
	}{
		"when prefix is in s": {
			s:      "token0.id",
			prefix: "token0.",
			want:   "id",
		},
		"when prefix is not in s": {
			s:      "hello world",
			prefix: "token0.",
			want:   "hello world",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := cutPrefix(test.s, test.prefix)

			assert.Equal(t, test.want, got)
		})
	}
}

func TestValidateRequestOpts(t *testing.T) {
	t.Run("when query type is ById", func(t *testing.T) {
		t.Run("when IncludeFields and ExcludeFields are both empty", func(t *testing.T) {
			opts := &RequestOptions{}
			err := validateRequestOpts(ById, opts)

			assert.Nil(t, err)
		})

		t.Run("when IncludeFields is not empty and ExcludeFields is empty", func(t *testing.T) {
			opts := &RequestOptions{
				IncludeFields: []string{"id", "liquidity", "txCount"},
				ExcludeFields: []string{},
			}
			err := validateRequestOpts(ById, opts)

			assert.Nil(t, err)
		})

		t.Run("when IncludeFields is empty and ExcludeFields is not empty", func(st *testing.T) {
			opts := &RequestOptions{
				IncludeFields: []string{},
				ExcludeFields: []string{"liquidity", "txCount"},
			}
			err := validateRequestOpts(ById, opts)

			assert.Nil(st, err)
		})

		t.Run("when IncludeFields doesn't have '*' and ExcludeFields is not empty", func(t *testing.T) {
			opts := &RequestOptions{
				IncludeFields: []string{"id"},
				ExcludeFields: []string{"liquidity"},
			}
			err := validateRequestOpts(ById, opts)

			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), "ExcludeFields can only be provided when IncludeFields is set to '*'")
		})

		t.Run("when List options are provided", func(t *testing.T) {
			opts := &RequestOptions{
				IncludeFields: []string{"*"},
				First:         100,
				Skip:          1000,
				OrderBy:       "id",
				OrderDir:      "desc",
			}
			err := validateRequestOpts(ById, opts)

			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), "List query options (First, Skip, OrderBy, OrderDir) should not be provided for ById queries")
		})
	})

	t.Run("when query type is List", func(t *testing.T) {
		t.Run("when IncludeFields and ExcludeFields are both empty", func(t *testing.T) {
			err := validateRequestOpts(List, &RequestOptions{})

			assert.Nil(t, err)
		})

		t.Run("when IncludeFields is not empty and ExcludeFields is empty", func(t *testing.T) {
			opts := &RequestOptions{
				IncludeFields: []string{"id", "liquidity", "txCount"},
				ExcludeFields: []string{},
			}
			err := validateRequestOpts(List, opts)

			assert.Nil(t, err)
		})

		t.Run("when IncludeFields is empty and ExcludeFields is not empty", func(t *testing.T) {
			opts := &RequestOptions{
				IncludeFields: []string{},
				ExcludeFields: []string{"liquidity", "txCount"},
			}
			err := validateRequestOpts(List, opts)

			assert.Nil(t, err)
		})

		t.Run("when List options are provided and valid", func(t *testing.T) {
			opts := &RequestOptions{
				IncludeFields: []string{"*"},
				First:         100,
				Skip:          1000,
				OrderBy:       "id",
				OrderDir:      "desc",
			}
			err := validateRequestOpts(List, opts)

			assert.Nil(t, err)
		})

		t.Run("when default List options are returned", func(t *testing.T) {
			opts := &RequestOptions{
				IncludeFields: []string{"*"},
				First:         0,
				OrderBy:       "",
				OrderDir:      "",
			}
			err := validateRequestOpts(List, opts)

			assert.Nil(t, err)
			assert.Equal(t, 100, opts.First)
			assert.Equal(t, "id", opts.OrderBy)
			assert.Equal(t, "asc", opts.OrderDir)
		})

		t.Run("when IncludeFields doesn't have '*' and ExcludeFields is not empty", func(t *testing.T) {
			opts := &RequestOptions{
				IncludeFields: []string{"id"},
				ExcludeFields: []string{"liquidity"},
			}
			err := validateRequestOpts(List, opts)

			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), "ExcludeFields can only be provided when IncludeFields is set to '*'")
		})

		t.Run("when First is too large", func(t *testing.T) {
			opts := &RequestOptions{
				IncludeFields: []string{"*"},
				First:         10000,
			}
			err := validateRequestOpts(List, opts)

			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), "First is too large")
		})

		t.Run("when OrderDir is invalid", func(t *testing.T) {
			opts := &RequestOptions{
				IncludeFields: []string{"*"},
				OrderDir:      "hello world",
			}
			err := validateRequestOpts(List, opts)

			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), "'asc' and 'desc' are the only valid options for OrderDir")
		})
	})
}

func TestPluralizeModelName(t *testing.T) {
	tests := map[string]struct {
		name string
		want string
	}{
		"normal case": {
			name: "pool",
			want: "pools",
		},
		"special case": {
			name: "factory",
			want: "factories",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := pluralizeModelName(test.name)

			assert.Equal(t, test.want, got)
			// if got != test.want {
			// 	t.Errorf("plural name doesn't match. want: `%v` got: `%v`", test.want, got)
			// }
		})
	}
}
