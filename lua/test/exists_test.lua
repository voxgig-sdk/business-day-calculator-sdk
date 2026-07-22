-- BusinessDayCalculator SDK exists test

local sdk = require("business-day-calculator_sdk")

describe("BusinessDayCalculatorSDK", function()
  it("should create test SDK", function()
    local testsdk = sdk.test(nil, nil)
    assert.is_not_nil(testsdk)
  end)
end)
