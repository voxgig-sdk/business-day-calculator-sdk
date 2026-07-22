# BusinessDayCalculator SDK utility: make_context
require_relative '../core/context'
module BusinessDayCalculatorUtilities
  MakeContext = ->(ctxmap, basectx) {
    BusinessDayCalculatorContext.new(ctxmap, basectx)
  }
end
