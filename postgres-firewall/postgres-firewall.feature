Feature: reject postgres firewall rule 0.0.0.0
Scenario: reject postgres firewall rule 0.0.0.0
    Given the subscription is set
    When policy is applied
    Then firewall rule should be rejected