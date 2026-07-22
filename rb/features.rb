# BusinessDayCalculator SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/test_feature'


module BusinessDayCalculatorFeatures
  def self.make_feature(name)
    case name
    when "base"
      BusinessDayCalculatorBaseFeature.new
    when "test"
      BusinessDayCalculatorTestFeature.new
    else
      BusinessDayCalculatorBaseFeature.new
    end
  end
end
