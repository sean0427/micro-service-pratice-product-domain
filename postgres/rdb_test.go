package rdbreposity

import (
	"testing"
)

// TODO
func Test_repository_Get(t *testing.T) {
	tests := []struct {
		name    string
		r       *repository
		wantLen int
		wantErr bool
	}{
		{
			name:    "get all products",
			r:       &repository{},
			wantLen: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{}
			got, err := r.Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.wantLen {
				t.Errorf("repository.Get() len = %v, want %v", len(got), tt.wantLen)
			}
		})
	}
}
