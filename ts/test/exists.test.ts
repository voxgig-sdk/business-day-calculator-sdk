
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { BusinessDayCalculatorSDK } from '..'


describe('exists', async () => {

  test('test-mode', () => {
    const testsdk = BusinessDayCalculatorSDK.test()
    equal(testsdk instanceof BusinessDayCalculatorSDK, true,
      'BusinessDayCalculatorSDK.test() must return a client synchronously')
  })

})
