<?php
declare(strict_types=1);

// BusinessDayCalculator SDK exists test

require_once __DIR__ . '/../businessdaycalculator_sdk.php';

use PHPUnit\Framework\TestCase;

class ExistsTest extends TestCase
{
    public function test_create_test_sdk(): void
    {
        $testsdk = BusinessDayCalculatorSDK::test(null, null);
        $this->assertNotNull($testsdk);
    }
}
