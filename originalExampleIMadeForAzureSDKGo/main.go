package main

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/postgresql/mgmt/postgresqlgit "
	"github.com/Azure/azure-sdk-for-go/sdk/to"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func main() {
	// create a firewall client
	pgfirewall := postgresql.NewFirewallRulesClient("e32cf796-5dbc-49a6-a569-c7255a117e0b")
	// create an authorizer from env vars or Azure Managed Service Idenity
	authorizer, err := auth.NewAuthorizerFromCLI()
	if err == nil {
		pgfirewall.Authorizer = authorizer
	}
	_, err2 := pgfirewall.CreateOrUpdate(
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

	if err != nil {
		panic(err2)
	}
}
