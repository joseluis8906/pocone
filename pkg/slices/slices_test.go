package slices_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/joseluis8906/pocone/pkg/slices"
)

func TestMap(t *testing.T) {
	testCases := map[string]struct {
		fn    func(any) any
		input []any
		want  []any
	}{
		"strings": {
			fn: func(a any) any {
				x := a.(string)
				return fmt.Sprintf("modified %s", x)
			},
			input: []any{
				"house",
				"sky",
			},
			want: []any{
				"modified house",
				"modified sky",
			},
		},
		"ints": {
			fn: func(a any) any {
				x := a.(int)
				return x * 10
			},
			input: []any{
				10,
				20,
			},
			want: []any{
				100,
				200,
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := slices.Map(tc.input, tc.fn)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				tv := reflect.TypeOf(tc.fn)
				t.Errorf("slices.Map(%v, %v) = %v; want = %v", tc.input, tv.String(), got, tc.want)
			}
		})
	}
}

func TestSplice(t *testing.T) {
	testCases := map[string]struct {
		index int
		input []any
		want  []any
	}{
		"first": {
			index: 0,
			input: []any{
				"a",
				"b",
				"c",
			},
			want: []any{
				"b",
				"c",
			},
		},
		"middle": {
			index: 1,
			input: []any{
				10,
				20,
				30,
			},
			want: []any{
				10,
				30,
			},
		},
		"last": {
			index: 2,
			input: []any{
				true,
				false,
				true,
			},
			want: []any{
				true,
				false,
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := slices.Splice(tc.input, tc.index)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("slices.Splice(%v, %v) = %v; want %v", tc.input, tc.index, got, tc.want)
			}
		})
	}
}
