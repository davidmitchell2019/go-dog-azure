package main

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
	"github.com/Azure/azure-sdk-for-go/services/preview/postgresql/mgmt/2017-12-01-preview/postgresql"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func createFirewallRule() {
	pgfirewall := postgresql.NewFirewallRulesClient("e32cf796-5dbc-49a6-a569-c7255a117e0b")
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err == nil {
		pgfirewall.Authorizer = authorizer
	}
	pgfirewall.CreateOrUpdate(
		context.Background(),
		"databricks",
		"test-for-logs",
		"allow 0.0.0.0",
		postgresql.FirewallRule{
			FirewallRuleProperties: &postgresql.FirewallRuleProperties{
				StartIPAddress: to.StringPtr("0.0.0.0"),
				EndIPAddress:   to.StringPtr("0.0.0.0"),
			},
		},
	)
}

func main() {
	createFirewallRule()
}
