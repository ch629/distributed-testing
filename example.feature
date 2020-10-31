Feature: adding items controller
    Scenario: valid request inserts into the database
    References: BJT-101
        Given: the database is empty
        When: sending a request to endpoint "endpoint_name"
        With: a "valid_request" body
        Then: the request should be successful
        And: the database should contain "1" record