package main

import (
	"context"
	"os/exec"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/postgresql/mgmt/postgresql"
	"github.com/Azure/azure-sdk-for-go/sdk/to"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/cucumber/godog"
)

func theSubscriptionIsSet() error {
	cmd := exec.Command("sh", "-c", "az account set -s ")
	return godog.ErrPending
}

func policyIsApplied() error {
	cmd := exec.Command("sh", "-c", "")
	return godog.ErrPending
}

func firewallRuleShouldBeRejected() error {
	// create a firewall client
	pgfirewall := postgresql.NewFirewallRulesClient("")
	// create an authorizer from env vars or Azure Managed Service Idenity
	authorizer, err := auth.NewAuthorizerFromEnvironment()
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
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^the subscription is set$`, theSubscriptionIsSet)
	s.Step(`^policy is applied$`, policyIsApplied)
	s.Step(`^firewall rule should be rejected$`, firewallRuleShouldBeRejected)
}
