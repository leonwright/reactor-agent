package config

import (
	"runtime"
	"time"

	"github.com/namsral/flag"
)

const defaultTick = 60 * time.Second

// Config is the Application configuration Object
type Config struct {
	DNSUpdateInterval   time.Duration
	DNSARecordNames     string
	AzureClientID       string
	AzureClientSecret   string
	AzureSubscriptionID string
	AzureTenantID       string
	AzureDNSZoneName    string
	AzureResourceGroup  string
}

func (c *Config) Init(args []string) error {
	var defaultConfigPath string
	if runtime.GOOS == "darwin" {
		defaultConfigPath = "/usr/local/etc/reactorapp/default.conf"
	} else if runtime.GOOS == "linux" {
		defaultConfigPath = "/etc/reactoragent/default.conf"
	}

	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	flags.String(flag.DefaultConfigFlagname, defaultConfigPath, "Path to config file")


	var (
		dnsUpdateInterval   = flags.Duration("dns_update_interval", defaultTick, "Ticking interval")
		dnsARecordNames     = flags.String("dns_a_records", "dev,*.dev", "DNS A Record Names")
		azureClientId       = flags.String("azure_client_id", "", "Azure Client ID")
		azureClientSecret   = flags.String("azure_client_secret", "", "Azure Client Secret")
		azureSubscriptionId = flags.String("azure_subscription_id", "", "Azure Subscription ID")
		azureTenantId       = flags.String("azure_tenant_id", "", "Azure Tenant ID")
		azureDNSZoneName    = flags.String("azure_dns_zone_name", "", "Azure DNS Zone name")
		azureResourceGroup  = flags.String("azure_resource_group", "", "Azure Resource Group")
	)

	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	c.DNSUpdateInterval = *dnsUpdateInterval
	c.DNSARecordNames = *dnsARecordNames
	c.AzureClientID = *azureClientId
	c.AzureClientSecret = *azureClientSecret
	c.AzureSubscriptionID = *azureSubscriptionId
	c.AzureTenantID = *azureTenantId
	c.AzureDNSZoneName = *azureDNSZoneName
	c.AzureResourceGroup = *azureResourceGroup
	return nil
}
