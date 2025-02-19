package _default

import (
	"reflect"
	"testing"
)

func TestDefault_Query(t *testing.T) {
	type fields struct {
		Addr string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Normal",
			fields: fields{
				Addr: "http://ip.sb",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Default{
				Addr: tt.fields.Addr,
			}
			got, err := d.Query()
			if (err != nil) != tt.wantErr {
				t.Errorf("Query() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, "") {
				t.Errorf("Query() got = %v", got)
			}
		})
	}
}
