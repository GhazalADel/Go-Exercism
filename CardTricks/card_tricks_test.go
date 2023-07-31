package cards

import (
	"reflect"
	"testing"
)

func TestGetItem(t *testing.T) {
	type args struct {
		slice []int
		index int
	}
	tests := []struct {
		name   string
		args   args
		want   int
		wantOk bool
	}{
		{
			name: "Retrieve item from slice by index",
			args: args{
				slice: []int{5, 2, 10, 6, 8, 7, 0, 9},
				index: 4,
			},
			want:   8,
			wantOk: true,
		},
		{
			name: "Get first item from slice",
			args: args{
				slice: []int{5, 2, 10, 6, 8, 7, 0, 9},
				index: 0,
			},
			want:   5,
			wantOk: true,
		},
		{
			name: "Get last item from slice",
			args: args{
				slice: []int{5, 2, 10, 6, 8, 7, 0, 9},
				index: 7,
			},
			want:   9,
			wantOk: true,
		},
		{
			name: "Index out of bounds",
			args: args{
				slice: []int{5, 2, 10, 6, 8, 7, 0, 9},
				index: 8,
			},
			want:   -1,
			wantOk: false,
		},
		{
			name: "Negative index",
			args: args{
				slice: []int{5, 2, 10, 6, 8, 7, 0, 9},
				index: -1,
			},
			want:   -1,
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetItem(tt.args.slice, tt.args.index)
			if got != tt.want {
				t.Errorf("GetItem(slice:%v, index:%v) got = %v, want %v", tt.args.slice, tt.args.index, got, tt.want)
			}
		})
	}
}

func TestSetItem(t *testing.T) {
	type args struct {
		slice []int
		index int
		value int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Overwrite an existing item",
			args: args{
				slice: []int{5, 2, 10, 6, 8, 7, 0, 9},
				index: 4,
				value: 1,
			},
			want: []int{5, 2, 10, 6, 1, 7, 0, 9},
		},
		{
			name: "Overwrite first item",
			args: args{
				slice: []int{5, 2, 10, 6, 8, 7, 0, 9},
				index: 0,
				value: 8,
			},
			want: []int{8, 2, 10, 6, 8, 7, 0, 9},
		},
		{
			name: "Overwrite last item",
			args: args{
				slice: []int{5, 2, 10, 6, 8, 7, 0, 9},
				index: 7,
				value: 8,
			},
			want: []int{5, 2, 10, 6, 8, 7, 0, 8},
		},
		{
			name: "Index out of bounds",
			args: args{
				slice: []int{5, 2, 10, 6, 8, 7, 0, 9},
				index: 8,
				value: 8,
			},
			want: []int{5, 2, 10, 6, 8, 7, 0, 9, 8},
		},
		{
			name: "Negative index",
			args: args{
				slice: []int{5, 2, 10, 6, 8, 7, 0, 9},
				index: -1,
				value: 8,
			},
			want: []int{5, 2, 10, 6, 8, 7, 0, 9, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetItem(tt.args.slice, tt.args.index, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetItem(slice:%v, index:%v, value:%v) = %v, want %v",
					tt.args.slice, tt.args.index, tt.args.value, got, tt.want)
			}
		})
	}
}

func TestRemoveItem(t *testing.T) {
	type args struct {
		slice []int
		index int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Remove an item",
			args: args{
				slice: []int{3, 4, 5, 6},
				index: 1,
			},
			want: []int{3, 5, 6},
		},
		{
			name: "Remove the first item",
			args: args{
				slice: []int{3, 4, 5, 6},
				index: 0,
			},
			want: []int{4, 5, 6},
		},
		{
			name: "Remove the last item",
			args: args{
				slice: []int{3, 4, 5, 6},
				index: 3,
			},
			want: []int{3, 4, 5},
		},
		{
			name: "Remove out of bounds index",
			args: args{
				slice: []int{3, 4, 5, 6},
				index: 7,
			},
			want: []int{3, 4, 5, 6},
		},
		{
			name: "Remove negative index",
			args: args{
				slice: []int{3, 4, 5, 6},
				index: -7,
			},
			want: []int{3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveItem(copySlice(tt.args.slice), tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveItem(slice:%v, index:%v) = %v, want %v", tt.args.slice, tt.args.index, got, tt.want)
			}
		})
	}
}

func copySlice(s []int) []int {
	var slice = make([]int, len(s))
	copy(slice, s)
	return slice
}
