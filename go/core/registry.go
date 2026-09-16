package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewRatelimitFeatureFunc func() Feature

var NewRetryFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewTimeoutFeatureFunc func() Feature

var NewDnsResultEntityFunc func(client *BusinessDayCalculatorSDK, entopts map[string]any) BusinessDayCalculatorEntity

var NewDomainEntityFunc func(client *BusinessDayCalculatorSDK, entopts map[string]any) BusinessDayCalculatorEntity

var NewEmailValidateEntityFunc func(client *BusinessDayCalculatorSDK, entopts map[string]any) BusinessDayCalculatorEntity

var NewGenerateEntityFunc func(client *BusinessDayCalculatorSDK, entopts map[string]any) BusinessDayCalculatorEntity

var NewGrammarEntityFunc func(client *BusinessDayCalculatorSDK, entopts map[string]any) BusinessDayCalculatorEntity

var NewIpnEntityFunc func(client *BusinessDayCalculatorSDK, entopts map[string]any) BusinessDayCalculatorEntity

var NewRedactEntityFunc func(client *BusinessDayCalculatorSDK, entopts map[string]any) BusinessDayCalculatorEntity

var NewSslEntityFunc func(client *BusinessDayCalculatorSDK, entopts map[string]any) BusinessDayCalculatorEntity

var NewUtilityEntityFunc func(client *BusinessDayCalculatorSDK, entopts map[string]any) BusinessDayCalculatorEntity

var NewWhoiEntityFunc func(client *BusinessDayCalculatorSDK, entopts map[string]any) BusinessDayCalculatorEntity

