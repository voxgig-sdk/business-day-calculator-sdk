<?php
declare(strict_types=1);

// BusinessDayCalculator SDK utility: result_headers

class BusinessDayCalculatorResultHeaders
{
    public static function call(BusinessDayCalculatorContext $ctx): ?BusinessDayCalculatorResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
