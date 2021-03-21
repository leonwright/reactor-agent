package network

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/dns/mgmt/dns"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/leonwright/devhelper/pkg/config"
)

func GetCurrentIP() string {
	URL := "https://api.ipify.org?format=json"
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Fatalln(err)
	}
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	var reqResp IPResponse
	if err := json.NewDecoder(resp.Body).Decode(&reqResp); err != nil {
		log.Fatalln(err)
	}

	return reqResp.IP
}

func NewServicePrincipalTokenFromCredentials(c *config.Config, scope string) (*adal.ServicePrincipalToken, error) {
	oauthConfig, err := adal.NewOAuthConfig(azure.PublicCloud.ActiveDirectoryEndpoint, c.AzureTenantID)
	if err != nil {
		panic(err)
	}
	return adal.NewServicePrincipalToken(*oauthConfig, c.AzureClientID, c.AzureClientSecret, scope)
}

func printNS(rrset dns.RecordSet) {
	fmt.Printf("*** NS Record ***\n")
	if rrset.NsRecords != nil {
		for _, ns := range *rrset.NsRecords {
			fmt.Printf("Nameserver: %s\n", *ns.Nsdname)
		}
	} else {
		fmt.Printf("*** None ***\n")
	}
}

func printSOA(rrset dns.RecordSet) {
	fmt.Printf("*** SOA Record ***\n")

	if rrset.SoaRecord != nil {
		fmt.Printf("Email: %s\n", *rrset.SoaRecord.Email)
		fmt.Printf("Host: %s\n", *rrset.SoaRecord.Host)
	} else {
		fmt.Printf("*** None ***\n")
	}

}

func printCNames(rrset dns.RecordSet) {
	fmt.Printf("*** CNAME Record ***\n")
	if rrset.CnameRecord != nil {
		fmt.Printf("Cname %s\n", *rrset.CnameRecord.Cname)
	} else {
		fmt.Printf("*** None ***\n")
	}

}
func printARecords(rrset dns.RecordSet) {
	fmt.Printf("*** A Records ***\n")
	if rrset.ARecords != nil {
		for _, arec := range *rrset.ARecords {
			fmt.Printf("record %s\n", *arec.Ipv4Address)
		}
	} else {
		fmt.Printf("*** None ***\n")
	}

}

func GetDNSRecord(c *config.Config) {
	resourceGroup := "networking"
	spt, err := NewServicePrincipalTokenFromCredentials(c, azure.PublicCloud.ResourceManagerEndpoint)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
	dc := dns.NewZonesClient(c.AzureSubscriptionID)
	dc.Authorizer = autorest.NewBearerAuthorizer(spt)
	rc := dns.NewRecordSetsClient(c.AzureSubscriptionID)
	rc.Authorizer = autorest.NewBearerAuthorizer(spt)

	var top int32
	top = 10
	rrsets, err := rc.ListByDNSZone(context.Background(), resourceGroup, "nerderbur.tech", &top, "")
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	for _, rrset := range rrsets.Values() {
		fmt.Printf("Recordset: %s Type: %s\n", *rrset.Name, *rrset.Type)

		switch *rrset.Type {
		case "Microsoft.Network/dnszones/A":
			printARecords(rrset)
		case "Microsoft.Network/dnszones/CNAME":
			printCNames(rrset)
		case "Microsoft.Network/dnszones/NS":
			printNS(rrset)
		case "Microsoft.Network/dnszones/SOA":
			printSOA(rrset)
		}
	}

}

func UpdateDevDNS(c *config.Config, ip string) {
	log.Printf("Running Method UpdatedDevDNS with IP %s", ip)
	defer log.Printf("Exit Method UpdatedDevDNS.")

	resourceGroup := "networking"
	zoneName := "nerderbur.tech"

	spt, err := NewServicePrincipalTokenFromCredentials(c, azure.PublicCloud.ResourceManagerEndpoint)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
	dc := dns.NewZonesClient(c.AzureSubscriptionID)
	dc.Authorizer = autorest.NewBearerAuthorizer(spt)
	rc := dns.NewRecordSetsClient(c.AzureSubscriptionID)
	rc.Authorizer = autorest.NewBearerAuthorizer(spt)

	arecordparams := &dns.RecordSet{
		RecordSetProperties: &dns.RecordSetProperties{
			TTL: to.Int64Ptr(60),
			ARecords: &[]dns.ARecord{
				{
					Ipv4Address: to.StringPtr(ip),
				},
			},
		},
	}

	s := strings.Split(c.DNSARecordNames, ",")

	for _, record := range s {
		_, err = rc.CreateOrUpdate(context.Background(), resourceGroup, zoneName, record, dns.A, *arecordparams, "", "")
		if err != nil {
			log.Fatalf("Error creating dev ARecord: %s, %v", zoneName, err)
			return
		}
	}
}
