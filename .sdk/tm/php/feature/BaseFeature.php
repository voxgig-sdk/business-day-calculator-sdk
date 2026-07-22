<?php
declare(strict_types=1);

// BusinessDayCalculator SDK base feature

class BusinessDayCalculatorBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    // Positions this feature when added via the client `extend` option:
    // "__before__" / "__after__" / "__replace__" name an already-added
    // feature (mirrors the ts feature `_options`). Declared so setting it
    // on an extension instance avoids the dynamic-property deprecation.
    public ?array $_options = null;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(BusinessDayCalculatorContext $ctx, array $options): void {}
    public function PostConstruct(BusinessDayCalculatorContext $ctx): void {}
    public function PostConstructEntity(BusinessDayCalculatorContext $ctx): void {}
    public function SetData(BusinessDayCalculatorContext $ctx): void {}
    public function GetData(BusinessDayCalculatorContext $ctx): void {}
    public function GetMatch(BusinessDayCalculatorContext $ctx): void {}
    public function SetMatch(BusinessDayCalculatorContext $ctx): void {}
    public function PrePoint(BusinessDayCalculatorContext $ctx): void {}
    public function PreSpec(BusinessDayCalculatorContext $ctx): void {}
    public function PreRequest(BusinessDayCalculatorContext $ctx): void {}
    public function PreResponse(BusinessDayCalculatorContext $ctx): void {}
    public function PreResult(BusinessDayCalculatorContext $ctx): void {}
    public function PreDone(BusinessDayCalculatorContext $ctx): void {}
    public function PreUnexpected(BusinessDayCalculatorContext $ctx): void {}
}
