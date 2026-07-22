<?php
declare(strict_types=1);

// BusinessDayCalculator SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class BusinessDayCalculatorMakeContext
{
    public static function call(array $ctxmap, ?BusinessDayCalculatorContext $basectx): BusinessDayCalculatorContext
    {
        return new BusinessDayCalculatorContext($ctxmap, $basectx);
    }
}
