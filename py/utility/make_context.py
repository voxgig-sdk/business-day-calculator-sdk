# BusinessDayCalculator SDK utility: make_context

from core.context import BusinessDayCalculatorContext


def make_context_util(ctxmap, basectx):
    return BusinessDayCalculatorContext(ctxmap, basectx)
