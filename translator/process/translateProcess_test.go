package process

import (
	"reflect"
	"testing"
)

func TestNewProcessTranslate(t *testing.T) {
	tests := []struct {
		name string
		want *processTranslate
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProcessTranslate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProcessTranslate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_processTranslate_GetTranslateProcess(t *testing.T) {
	type args struct {
		request interface{}
	}
	tests := []struct {
		name    string
		pt      *processTranslate
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.pt.GetTranslateProcess(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("processTranslate.GetTranslateProcess() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processTranslate.GetTranslateProcess() = %v, want %v", got, tt.want)
			}
		})
	}
}
