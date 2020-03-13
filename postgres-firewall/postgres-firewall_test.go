package main

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/postgresql/mgmt/postgresql"
	"github.com/Azure/azure-sdk-for-go/sdk/to"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/cucumber/godog"
)

func theSubscriptionIsSet() error {
	cmd := exec.Command("sh", "-c", "az account set -s ")
	fmt.Println(cmd)
	return godog.ErrPending
}

func policyIsApplied() error {
	cmd := exec.Command("sh", "-c", "")
	fmt.Println(cmd)
	return godog.ErrPending
}

func firewallRuleShouldBeRejected() error {
	// create a firewall client
	pgfirewall := postgresql.NewFirewallRulesClient("e32cf796-5dbc-49a6-a569-c7255a117e0b")
	// create an authorizer from env vars or Azure Managed Service Idenity
	authorizer, err := auth.NewAuthorizerFromCLI()
	if err == nil {
		pgfirewall.Authorizer = authorizer
	}
	_, err = pgfirewall.CreateOrUpdate(
		context.Background(),
		"test-rg",
		"test-server-for-bdd",
		"allow-internet",
		postgresql.FirewallRule{
			FirewallRuleProperties: &postgresql.FirewallRuleProperties{
				StartIPAddress: to.StringPtr("0.0.0.0"),
				EndIPAddress:   to.StringPtr("255.255.255.255"),
			},
		},
	)
	if err != nil {
		panic(err)
	}
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^the subscription is set$`, theSubscriptionIsSet)
	s.Step(`^policy is applied$`, policyIsApplied)
	s.Step(`^firewall rule should be rejected$`, firewallRuleShouldBeRejected)
}
