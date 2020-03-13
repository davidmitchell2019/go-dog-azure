package main

import (
	"context"
	"log"
	"os"
	"os/exec"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/postgresql/mgmt/postgresql"
	"github.com/Azure/azure-sdk-for-go/sdk/to"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/cucumber/godog"
)

func subscriptionIsSet() error {
	cmd := exec.Command("az", "account", "set", "--subscription", "e32cf796-5dbc-49a6-a569-c7255a117e0b")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	return nil
}
func policyIsApplied() error {
	cmd := exec.Command("az", "policy", "assignment", "show", "--name", "7ee092beb6074d05bfafbe31")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	return nil
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
				EndIPAddress:   to.StringPtr("0.0.0.0"),
			},
		},
	)
	if err != nil {
		panic(err)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^subscription is set$`, subscriptionIsSet)
	s.Step(`^policy is applied$`, policyIsApplied)
	s.Step(`^firewall rule should be rejected$`, firewallRuleShouldBeRejected)
}
