package contract

import (
	"reflect"
	"testing"
)

func TestCreateUser_Validate(t *testing.T) {
	tests := []struct {
		name string
		c    *SignUpUser
		want map[string]string
	}{
		{
			name: "test valid request",
			c: &SignUpUser{
				Name:        "venkat raman kannan",
				Username:    "venkatramankannanxo",
				Password:    "ThisIsAGoodPassword@99On100",
				PhoneNumber: "9900903821",
			},
			want: map[string]string{},
		},
		{
			name: "test invalid request",
			c: &SignUpUser{
				Name:        "",
				Username:    "1920234",
				Password:    "thisnot",
				PhoneNumber: "23900921",
			},
			want: map[string]string{
				"name":         "err-name-is-required",
				"username":     "err-username-should-not-be-lesser-than-8-characters",
				"phone_number": "err-phone-number-is-invalid",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Validate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
