package implementations

import (
	"reflect"
	"testing"
)

func TestBinary_Translate(t *testing.T) {
	type args struct {
		req interface{}
	}
	tests := []struct {
		name    string
		b       *Binary
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.b.Translate(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Binary.Translate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Binary.Translate() = %v, want %v", got, tt.want)
			}
		})
	}
}
