<?php
declare(strict_types=1);

// BusinessDayCalculator SDK utility: feature_hook

class BusinessDayCalculatorFeatureHook
{
    public static function call(BusinessDayCalculatorContext $ctx, string $name): void
    {
        if (!$ctx->client) {
            return;
        }
        $features = $ctx->client->features ?? null;
        if (!$features) {
            return;
        }
        foreach ($features as $f) {
            if (method_exists($f, $name)) {
                $f->$name($ctx);
            }
        }
    }
}
