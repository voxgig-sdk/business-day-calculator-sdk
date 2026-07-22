<?php
declare(strict_types=1);

// BusinessDayCalculator SDK feature factory

require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/feature/TestFeature.php';


class BusinessDayCalculatorFeatures
{
    public static function make_feature(string $name)
    {
        switch ($name) {
            case "base":
                return new BusinessDayCalculatorBaseFeature();
            case "test":
                return new BusinessDayCalculatorTestFeature();
            default:
                return new BusinessDayCalculatorBaseFeature();
        }
    }
}
