<?php
declare(strict_types=1);

// BusinessDayCalculator SDK utility: result_body

class BusinessDayCalculatorResultBody
{
    public static function call(BusinessDayCalculatorContext $ctx): ?BusinessDayCalculatorResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}
