@endpoint(fastly-integration) @endpoint(fastly-integration-v1)
Feature: Fastly Integration
  Configure your [Datadog-Fastly
  integration](https://docs.datadoghq.com/api/?lang=bash#integration-
  fastly) directly through the Datadog API.

  Background:
    Given a valid "apiKeyAuth" key in the system
    And a valid "appKeyAuth" key in the system
    And an instance of "FastlyIntegration" API

  @generated @skip
  Scenario: Create a new service object returns "Bad Request" response
    Given new "CreateFastlyIntegrationService" request
    And body with value {"service_id": "", "service_tags": "", "service_account": ""}
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip
  Scenario: Create a new service object returns "OK" response
    Given new "CreateFastlyIntegrationService" request
    And body with value {"service_id": "", "service_tags": "", "service_account": ""}
    When the request is sent
    Then the response status is 201 OK

  @generated @skip
  Scenario: Delete a single service object returns "Item Not Found" response
    Given new "DeleteFastlyIntegrationService" request
    And request contains "service_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 404 Item Not Found

  @generated @skip
  Scenario: Delete a single service object returns "OK" response
    Given new "DeleteFastlyIntegrationService" request
    And request contains "service_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Get a single service object returns "Item Not Found" response
    Given new "GetFastlyIntegrationService" request
    And request contains "service_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 404 Item Not Found

  @generated @skip
  Scenario: Get a single service object returns "OK" response
    Given new "GetFastlyIntegrationService" request
    And request contains "service_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Update a single service object returns "Bad Request" response
    Given new "UpdateFastlyIntegrationService" request
    And request contains "service_id" parameter from "<PATH>"
    And body with value {"service_tags": "", "service_account": ""}
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip
  Scenario: Update a single service object returns "Item Not Found" response
    Given new "UpdateFastlyIntegrationService" request
    And request contains "service_id" parameter from "<PATH>"
    And body with value {"service_tags": "", "service_account": ""}
    When the request is sent
    Then the response status is 404 Item Not Found

  @generated @skip
  Scenario: Update a single service object returns "OK" response
    Given new "UpdateFastlyIntegrationService" request
    And request contains "service_id" parameter from "<PATH>"
    And body with value {"service_tags": "", "service_account": ""}
    When the request is sent
    Then the response status is 200 OK
