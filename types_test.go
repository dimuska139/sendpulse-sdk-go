package sendpulse_sdk_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloat32Type_UnmarshalJSON(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		f       Float32
		args    args
		wantErr bool
	}{
		{
			name: "valid float string",
			f:    1.3,
			args: args{
				b: []byte("\"1.3\""),
			},
		}, {
			name: "valid float",
			f:    1.3,
			args: args{
				b: []byte("1.3"),
			},
		}, {
			name: "null",
			f:    0,
			args: args{
				b: []byte("null"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var f Float32
			err := f.UnmarshalJSON(tt.args.b)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.f, f)
		})
	}
}
