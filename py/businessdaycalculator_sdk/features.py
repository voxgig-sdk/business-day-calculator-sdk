# BusinessDayCalculator SDK feature factory

from businessdaycalculator_sdk.feature.base_feature import BusinessDayCalculatorBaseFeature
from businessdaycalculator_sdk.feature.ratelimit_feature import BusinessDayCalculatorRatelimitFeature
from businessdaycalculator_sdk.feature.retry_feature import BusinessDayCalculatorRetryFeature
from businessdaycalculator_sdk.feature.test_feature import BusinessDayCalculatorTestFeature
from businessdaycalculator_sdk.feature.timeout_feature import BusinessDayCalculatorTimeoutFeature


_FEATURES = {
    "base": lambda: BusinessDayCalculatorBaseFeature(),
    "ratelimit": lambda: BusinessDayCalculatorRatelimitFeature(),
    "retry": lambda: BusinessDayCalculatorRetryFeature(),
    "test": lambda: BusinessDayCalculatorTestFeature(),
    "timeout": lambda: BusinessDayCalculatorTimeoutFeature(),
}


def _make_feature(name):
    factory = _FEATURES.get(name)
    if factory is not None:
        return factory()
    return _FEATURES["base"]()


# True when this SDK was generated with the named feature class - the
# constructor's tolerance for extend-carried features reads this (an
# active name with no generated class must not become a BaseFeature
# stray when an extend instance carries it).
def _has_feature(name):
    return name in _FEATURES
