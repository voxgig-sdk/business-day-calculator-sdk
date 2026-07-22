
import { Context } from './Context'


class BusinessDayCalculatorError extends Error {

  isBusinessDayCalculatorError = true

  sdk = 'BusinessDayCalculator'

  code: string
  ctx: Context

  constructor(code: string, msg: string, ctx: Context) {
    super(msg)
    this.code = code
    this.ctx = ctx
  }

}

export {
  BusinessDayCalculatorError
}

