package network

import (
	"reflect"
	"testing"

	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/leonwright/devhelper/pkg/config"
)

func TestGetCurrentIP(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCurrentIP(); got != tt.want {
				t.Errorf("GetCurrentIP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewServicePrincipalTokenFromCredentials(t *testing.T) {
	type args struct {
		c     *config.Config
		scope string
	}
	tests := []struct {
		name    string
		args    args
		want    *adal.ServicePrincipalToken
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewServicePrincipalTokenFromCredentials(tt.args.c, tt.args.scope)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewServicePrincipalTokenFromCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServicePrincipalTokenFromCredentials() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateDevDNS(t *testing.T) {
	type args struct {
		c  *config.Config
		ip string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateDevDNS(tt.args.c, tt.args.ip)
		})
	}
}
