# BusinessDayCalculator SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/ratelimit_feature'
require_relative 'feature/retry_feature'
require_relative 'feature/test_feature'
require_relative 'feature/timeout_feature'


module BusinessDayCalculatorFeatures
  def self.make_feature(name)
    case name
    when "base"
      BusinessDayCalculatorBaseFeature.new
    when "ratelimit"
      BusinessDayCalculatorRatelimitFeature.new
    when "retry"
      BusinessDayCalculatorRetryFeature.new
    when "test"
      BusinessDayCalculatorTestFeature.new
    when "timeout"
      BusinessDayCalculatorTimeoutFeature.new
    else
      BusinessDayCalculatorBaseFeature.new
    end
  end
end
