# BusinessDayCalculator SDK exists test

require "minitest/autorun"
require_relative "../BusinessDayCalculator_sdk"

class ExistsTest < Minitest::Test
  def test_create_test_sdk
    testsdk = BusinessDayCalculatorSDK.test(nil, nil)
    assert !testsdk.nil?
  end
end
