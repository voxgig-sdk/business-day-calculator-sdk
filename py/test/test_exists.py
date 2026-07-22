# ProjectName SDK exists test

import pytest
from businessdaycalculator_sdk import BusinessDayCalculatorSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = BusinessDayCalculatorSDK.test(None, None)
        assert testsdk is not None
