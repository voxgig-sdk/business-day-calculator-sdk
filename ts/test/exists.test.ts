
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { BusinessDayCalculatorSDK } from '..'


describe('exists', async () => {

  test('test-mode', async () => {
    const testsdk = await BusinessDayCalculatorSDK.test()
    equal(null !== testsdk, true)
  })

})
