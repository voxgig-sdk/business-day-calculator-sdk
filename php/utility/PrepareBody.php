<?php
declare(strict_types=1);

// BusinessDayCalculator SDK utility: prepare_body

class BusinessDayCalculatorPrepareBody
{
    public static function call(BusinessDayCalculatorContext $ctx): mixed
    {
        if ($ctx->op->input === 'data') {
            return ($ctx->utility->transform_request)($ctx);
        }
        return null;
    }
}
