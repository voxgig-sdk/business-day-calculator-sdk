-- BusinessDayCalculator SDK error

local BusinessDayCalculatorError = {}
BusinessDayCalculatorError.__index = BusinessDayCalculatorError


function BusinessDayCalculatorError.new(code, msg, ctx)
  local self = setmetatable({}, BusinessDayCalculatorError)
  self.is_sdk_error = true
  self.sdk = "BusinessDayCalculator"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function BusinessDayCalculatorError:error()
  return self.msg
end


function BusinessDayCalculatorError:__tostring()
  return self.msg
end


return BusinessDayCalculatorError
