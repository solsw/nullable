package nullable

import (
	"reflect"
	"testing"
)

func TestNewNull_int(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "NewNull"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NewNull[int]()
			if v, ok := n.Get(); ok {
				t.Errorf("Nullable has value %v", v)
			}
		})
	}
}

func TestNew_bool(t *testing.T) {
	type args struct {
		v bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "false",
			args: args{
				v: false,
			},
			want: false,
		},
		{name: "true",
			args: args{
				v: true,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := New(tt.args.v)
			if !n.Has() {
				t.Errorf("n has no value")
				return
			}
			got, _ := n.Get()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nullable_Get_string(t *testing.T) {
	type args struct {
		n *Nullable[string]
	}
	tests := []struct {
		name   string
		args   args
		wantV  string
		wantOk bool
	}{
		{name: "NewNull",
			args: args{
				n: NewNull[string](),
			},
			wantV:  "",
			wantOk: false,
		},
		{name: "New",
			args: args{
				n: New("qwerty"),
			},
			wantV:  "qwerty",
			wantOk: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotV, gotOk := nullable_Get(tt.args.n)
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("nullable_Get() gotV = %v, want %v", gotV, tt.wantV)
			}
			if gotOk != tt.wantOk {
				t.Errorf("nullable_Get() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func Test_nullable_Set_int(t *testing.T) {
	type args struct {
		n *Nullable[int]
		v int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "NewNull",
			args: args{
				n: NewNull[int](),
				v: 1,
			},
			want: 1,
		},
		{name: "New",
			args: args{
				n: New(1),
				v: 2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nullable_Set(tt.args.n, tt.args.v)
			if got, _ := tt.args.n.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nullable_Set() got = %v, want %v", got, tt.want)
			}
		})
	}
}
