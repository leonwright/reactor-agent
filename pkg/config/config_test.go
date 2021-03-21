package config

import (
	"testing"
	"time"
)

func TestConfig_Init(t *testing.T) {
	type fields struct {
		DNSUpdateInterval   time.Duration
		DNSARecordNames     string
		AzureClientID       string
		AzureClientSecret   string
		AzureSubscriptionID string
		AzureTenantID       string
		AzureDNSZoneName    string
		AzureResourceGroup  string
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				DNSUpdateInterval:   tt.fields.DNSUpdateInterval,
				DNSARecordNames:     tt.fields.DNSARecordNames,
				AzureClientID:       tt.fields.AzureClientID,
				AzureClientSecret:   tt.fields.AzureClientSecret,
				AzureSubscriptionID: tt.fields.AzureSubscriptionID,
				AzureTenantID:       tt.fields.AzureTenantID,
				AzureDNSZoneName:    tt.fields.AzureDNSZoneName,
				AzureResourceGroup:  tt.fields.AzureResourceGroup,
			}
			if err := c.Init(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("Config.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
