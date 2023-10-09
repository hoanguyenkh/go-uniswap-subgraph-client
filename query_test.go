package unigraphclient

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructByIdQuery(t *testing.T) {
	t.Run("when request options are valid", func(t *testing.T) {
		opts := &RequestOptions{
			IncludeFields: []string{"id"},
		}
		testId := "test"
		req, err := constructByIdQuery(testId, PoolFields, opts)

		assert.Nil(t, err)
		assert.NotNil(t, req)

		vars := req.Vars()
		id, ok := vars["id"]
		assert.True(t, ok)
		assert.Equal(t, testId, id)
	})

	t.Run("when query assembly fails", func(t *testing.T) {
		opts := &RequestOptions{
			IncludeFields: []string{"not found"},
		}
		_, err := constructByIdQuery("test", PoolFields, opts)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("when gathering all model fields fails", func(t *testing.T) {
		model := modelFields{
			reference: map[string]string{
				"not found": "",
			},
		}
		_, err := constructByIdQuery("test", model, nil)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "reference field not found")
	})

	t.Run("when request options validation fails", func(t *testing.T) {
		opts := &RequestOptions{
			IncludeFields: []string{"id", "liquidity"},
			ExcludeFields: []string{"id", "liquidity"},
		}
		_, err := constructByIdQuery("test", PoolFields, opts)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "request options error")
	})
}

func TestConstructListQuery(t *testing.T) {
	t.Run("when request options are valid", func(t *testing.T) {
		opts := &RequestOptions{
			IncludeFields: []string{"id"},
			First:         50,
		}
		req, err := constructListQuery(PoolFields, opts)

		assert.Nil(t, err)
		assert.NotNil(t, req)

		vars := req.Vars()
		first, ok := vars["first"]
		assert.True(t, ok)
		assert.Equal(t, 50, first)
	})

	t.Run("when query assembly fails", func(t *testing.T) {
		opts := &RequestOptions{
			IncludeFields: []string{"not found"},
		}
		_, err := constructListQuery(PoolFields, opts)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("when gathering all model fields fails", func(t *testing.T) {
		model := modelFields{
			reference: map[string]string{
				"not found": "",
			},
		}
		_, err := constructListQuery(model, nil)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "reference field not found")
	})

	t.Run("when request options validation fails", func(t *testing.T) {
		opts := &RequestOptions{
			IncludeFields: []string{"id", "liquidity"},
			ExcludeFields: []string{"id", "liquidity"},
		}
		_, err := constructListQuery(PoolFields, opts)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "request options error")
	})
}

func TestAssembleQuery(t *testing.T) {
	// TODO: add more test cases based on IncludeFields & ExcludeFields combinations
	t.Run("when query type is ById", func(t *testing.T) {
		opts := &RequestOptions{
			IncludeFields: []string{"id"},
		}
		query, err := assembleQuery(ById, PoolFields, opts)

		assert.Nil(t, err)
		assert.Greater(t, len(query), 0)
	})

	t.Run("when query type is List", func(t *testing.T) {
		opts := &RequestOptions{
			IncludeFields: []string{"id"},
		}
		query, err := assembleQuery(List, PoolFields, opts)

		assert.Nil(t, err)
		assert.Greater(t, len(query), 0)
	})

	t.Run("when opts.Block is set", func(t *testing.T) {
		byIdOpts := &RequestOptions{
			IncludeFields: []string{"id"},
			Block:         1000000,
		}
		byIdQuery, err := assembleQuery(ById, PoolFields, byIdOpts)

		assert.Nil(t, err)
		assert.Greater(t, len(byIdQuery), 0)

		listOpts := &RequestOptions{
			IncludeFields: []string{"id"},
			Block:         1000000,
		}
		listQuery, err := assembleQuery(List, PoolFields, listOpts)

		assert.Nil(t, err)
		assert.Greater(t, len(listQuery), 0)
	})

	t.Run("when query type is unrecognized", func(t *testing.T) {
		opts := &RequestOptions{
			IncludeFields: []string{"id"},
		}
		_, err := assembleQuery(1000, PoolFields, opts)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "unrecognized query type")
	})

	t.Run("when reference field doesn't exist", func(t *testing.T) {
		model := modelFields{
			direct: []string{"id"},
			reference: map[string]string{
				"notFound": "",
			},
		}
		opts := &RequestOptions{
			IncludeFields: []string{"id", "notFound.id"},
		}
		_, err := assembleQuery(ById, model, opts)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "reference field not found")
	})

	t.Run("when reference sub-field doesn't exist", func(t *testing.T) {
		opts := &RequestOptions{
			IncludeFields: []string{"id", "token0.notFound"},
		}
		_, err := assembleQuery(ById, PoolFields, opts)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "unrecognized field given in opts.IncludeFields")
	})

	t.Run("when direct field doesn't exist", func(t *testing.T) {
		opts := &RequestOptions{
			IncludeFields: []string{"id", "notFound"},
		}
		_, err := assembleQuery(ById, PoolFields, opts)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "unrecognized field given in opts.IncludeFields")
	})
}

func TestAssembleQuery_FieldCombinations(t *testing.T) {
	tests := map[string]struct {
		model         modelFields
		includeFields []string
		excludeFields []string
		wantErr       bool
		wantErrMsg    string
	}{
		"factory include (valid)": {
			model:         FactoryFields,
			includeFields: []string{"id", "owner"},
			excludeFields: []string{},
		},
		"factory exclude (valid)": {
			model:         FactoryFields,
			includeFields: []string{"*"},
			excludeFields: []string{"totalValueLockedUSDUntracked", "totalValueLockedETHUntracked"},
		},
		"pool include (valid)": {
			model:         PoolFields,
			includeFields: []string{"id", "txCount", "token0.id", "token1.derivedETH", "token1.whitelistPools.txCount"},
			excludeFields: []string{},
		},
		"pool include (invalid)": {
			model:         PoolFields,
			includeFields: []string{"id", "txCount", "token1.whitelistPools.notFound"},
			excludeFields: []string{},
			wantErr:       true,
			wantErrMsg:    "unrecognized field",
		},
		"pool exclude (valid)": {
			model:         PoolFields,
			includeFields: []string{"*"},
			excludeFields: []string{"feeTier", "token0.symbol", "token1.whitelistPools.txCount"},
		},
		"token include (valid)": {
			model:         TokenFields,
			includeFields: []string{"*", "id", "decimals", "whitelistPools.txCount"},
			excludeFields: []string{},
		},
		"token exclude (valid)": {
			model:         TokenFields,
			includeFields: []string{"*"},
			excludeFields: []string{"id", "derivedETH", "notFound"}, // opts.ExcludeFields are not validated currently
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if slices.Contains(test.includeFields, "*") {
				fields, err := gatherModelFields(test.model, test.excludeFields, true)
				if err != nil {
					t.Fatal(err)
				}
				test.includeFields = fields
			}

			opts := &RequestOptions{
				IncludeFields: test.includeFields,
				ExcludeFields: test.excludeFields,
			}
			_, byIdErr := assembleQuery(ById, test.model, opts)
			_, listErr := assembleQuery(List, test.model, opts)

			if test.wantErr {
				assert.NotNil(t, byIdErr)
				assert.Contains(t, byIdErr.Error(), test.wantErrMsg)
				assert.NotNil(t, listErr)
				assert.Contains(t, listErr.Error(), test.wantErrMsg)
			} else {
				assert.Nil(t, byIdErr)
				assert.Nil(t, listErr)
			}
		})
	}
}

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
		"special case 1": {
			name: "factory",
			want: "factories",
		},
		"special case 2": {
			name: "flash",
			want: "flashes",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := pluralizeModelName(test.name)

			assert.Equal(t, test.want, got)
		})
	}
}
