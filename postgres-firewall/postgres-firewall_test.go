package main

import (
	"github.com/cucumber/godog"
)

func theSubscriptionIsSet() error {
	return godog.ErrPending
}

func policyIsApplied() error {
	return godog.ErrPending
}

func firewallRuleShouldBeRejected() error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^the subscription is set$`, theSubscriptionIsSet)
	s.Step(`^policy is applied$`, policyIsApplied)
	s.Step(`^firewall rule should be rejected$`, firewallRuleShouldBeRejected)
}
