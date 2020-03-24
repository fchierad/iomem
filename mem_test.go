package iomem

import (
	"reflect"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want *Mem
	}{
		{
			name: "Create in memory size-capped io.Writer",
			args: args{
				n: 10,
			},
			want: &Mem{
				data: make([]byte, 0, 20),
				size: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMem_Write(t *testing.T) {
	type fields struct {
		data []byte
		size int
	}
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantS   string
		wantN   int
		wantErr bool
	}{
		{
			name: "Empty Mem.data, p < size",
			fields: fields{
				data: nil,
				size: 10,
			},
			args: args{
				p: []byte("test1"),
			},
			wantS:   "test1",
			wantN:   5,
			wantErr: false,
		},
		{
			name: "Empty Mem.data, p = size",
			fields: fields{
				data: nil,
				size: 10,
			},
			args: args{
				p: []byte("first test"),
			},
			wantS:   "first test",
			wantN:   10,
			wantErr: false,
		},
		{
			name: "Empty Mem.data, p > size",
			fields: fields{
				data: nil,
				size: 10,
			},
			args: args{
				p: []byte("write data too long"),
			},
			wantS:   "a too long",
			wantN:   10,
			wantErr: false,
		},
		{
			name: "Populated Mem.data, m + p < size",
			fields: fields{
				data: []byte("start"),
				size: 10,
			},
			args: args{
				p: []byte("Test"),
			},
			wantS:   "startTest",
			wantN:   4,
			wantErr: false,
		},
		{
			name: "Populated Mem.data, m + p = size",
			fields: fields{
				data: []byte("start"),
				size: 10,
			},
			args: args{
				p: []byte("Test1"),
			},
			wantS:   "startTest1",
			wantN:   5,
			wantErr: false,
		},
		{
			name: "Populated Mem.data, m + p > size",
			fields: fields{
				data: []byte("AnIdeal"),
				size: 10,
			},
			args: args{
				p: []byte("Test1"),
			},
			wantS:   "IdealTest1",
			wantN:   5,
			wantErr: false,
		},
		{
			name: "Populated Mem.data, p = size",
			fields: fields{
				data: []byte("initial"),
				size: 10,
			},
			args: args{
				p: []byte("first test"),
			},
			wantS:   "first test",
			wantN:   10,
			wantErr: false,
		},
		{
			name: "Populated Mem.data, p > size",
			fields: fields{
				data: []byte("initial"),
				size: 10,
			},
			args: args{
				p: []byte("write data too long"),
			},
			wantS:   "a too long",
			wantN:   10,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := New(tt.fields.size)
			if tt.fields.data != nil {
				m.data = append(m.data, tt.fields.data...)
			}
			gotN, err := m.Write(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("Write() gotN = %v, want %v", gotN, tt.wantN)
			}
			if !strings.EqualFold(m.String(), tt.wantS) {
				t.Errorf("Write() m.String() = '%s', tt.wantS '%s'", m.String(), tt.wantS)
			}
		})
	}
}
