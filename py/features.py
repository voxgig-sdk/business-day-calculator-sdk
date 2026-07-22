# BusinessDayCalculator SDK feature factory

from feature.base_feature import BusinessDayCalculatorBaseFeature
from feature.test_feature import BusinessDayCalculatorTestFeature


def _make_feature(name):
    features = {
        "base": lambda: BusinessDayCalculatorBaseFeature(),
        "test": lambda: BusinessDayCalculatorTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
