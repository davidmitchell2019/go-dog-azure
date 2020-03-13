package main

import (
	"flag"
	"fmt"
	"os"
	"testing"
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
	"github.com/Azure/azure-sdk-for-go/services/preview/postgresql/mgmt/2017-12-01-preview/postgresql"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)
//Given the subscription is set
func setSubscription()
{
	cmd := exec.Command("sh", "-c", "az account set -s ")
}
//When policy is applied
func getPolicyAssignment()
{
	cmd := exec.Command("sh", "-c", "")
}
//Then firewall rule should be rejected
func createFirewallRule()
{
	// create a firewall client
	pgfirewall := postgresql.NewFirewallRulesClient("e32cf796-5dbc-49a6-a569-c7255a117e0b")
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
}
func FeatureContext(s *godog.Suite) {
	s.Step(`^the subscription is set`, setSubscription)
	s.Step(`^policy is applied$`, getPolicyAssignment)
	s.Step(`^firewall rule should be rejected$`, createFirewallRule)

	s.BeforeScenario(func(*messages.Pickle) {
		Godogs = 0 // clean the state before every scenario
	})
}
