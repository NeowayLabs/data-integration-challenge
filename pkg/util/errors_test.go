package util

import (
	"errors"
	"reflect"
	"testing"
)

func TestAppendError(t *testing.T) {
	type args struct {
		es []error
		e  error
	}
	tests := []struct {
		name string
		args args
		want []error
	}{
		{
			name: "nil + nil = nil",
			args: args{es: nil, e: nil},
			want: nil,
		},
		{
			name: "[] + nil = []",
			args: args{es: []error{}, e: nil},
			want: []error{},
		},
		{
			name: "[] + e = [e]",
			args: args{es: []error{}, e: errors.New("e")},
			want: []error{errors.New("e")},
		},
		{
			name: "nil + e = nil",
			args: args{es: nil, e: errors.New("e")},
			want: nil,
		},
		{
			name: "[x] + y = [x, y]",
			args: args{es: []error{
				errors.New("x")},
				e: errors.New("y"),
			},
			want: []error{
				errors.New("x"),
				errors.New("y"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AppendError(tt.args.es, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeErrors(t *testing.T) {
	type args struct {
		es []error
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name:    "Nil",
			args:    args{es: nil},
			wantErr: nil,
		},
		{
			name:    "Zero errors",
			args:    args{es: []error{}},
			wantErr: nil,
		},
		{
			name:    "Single error",
			args:    args{es: []error{errors.New("e")}},
			wantErr: errors.New("e"),
		},
		{
			name: "Multiple errors",
			args: args{
				es: []error{
					errors.New("x"),
					errors.New("y"),
				},
			},
			wantErr: errors.New("x. y"),
		},
		{
			name: "Multiple errors with nil",
			args: args{
				es: []error{
					nil,
					errors.New("x"),
					nil,
					errors.New("y"),
					nil,
				},
			},
			wantErr: errors.New("x. y"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MergeErrors(tt.args.es)
			if tt.wantErr != nil && err != nil {
				if tt.wantErr.Error() != err.Error() {
					t.Errorf("MergeErrors() error = `%v`, wantErr `%v`", err.Error(), tt.wantErr.Error())
				}
			} else if tt.wantErr != err {
				t.Errorf("MergeErrors() error = `%v`, wantErr `%v`", err, tt.wantErr)
			}
		})
	}
}
