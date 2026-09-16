package voxgigbusinessdaycalculatorsdk

import (
	"github.com/voxgig-sdk/business-day-calculator-sdk/go/core"
	"github.com/voxgig-sdk/business-day-calculator-sdk/go/entity"
	"github.com/voxgig-sdk/business-day-calculator-sdk/go/feature"
	_ "github.com/voxgig-sdk/business-day-calculator-sdk/go/utility"
)

// Type aliases preserve external API.
type BusinessDayCalculatorSDK = core.BusinessDayCalculatorSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type BusinessDayCalculatorEntity = core.BusinessDayCalculatorEntity
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type BusinessDayCalculatorError = core.BusinessDayCalculatorError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewRatelimitFeatureFunc = func() core.Feature {
		return feature.NewRatelimitFeature()
	}
	core.NewRetryFeatureFunc = func() core.Feature {
		return feature.NewRetryFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewTimeoutFeatureFunc = func() core.Feature {
		return feature.NewTimeoutFeature()
	}
	core.NewDnsResultEntityFunc = func(client *core.BusinessDayCalculatorSDK, entopts map[string]any) core.BusinessDayCalculatorEntity {
		return entity.NewDnsResultEntity(client, entopts)
	}
	core.NewDomainEntityFunc = func(client *core.BusinessDayCalculatorSDK, entopts map[string]any) core.BusinessDayCalculatorEntity {
		return entity.NewDomainEntity(client, entopts)
	}
	core.NewEmailValidateEntityFunc = func(client *core.BusinessDayCalculatorSDK, entopts map[string]any) core.BusinessDayCalculatorEntity {
		return entity.NewEmailValidateEntity(client, entopts)
	}
	core.NewGenerateEntityFunc = func(client *core.BusinessDayCalculatorSDK, entopts map[string]any) core.BusinessDayCalculatorEntity {
		return entity.NewGenerateEntity(client, entopts)
	}
	core.NewGrammarEntityFunc = func(client *core.BusinessDayCalculatorSDK, entopts map[string]any) core.BusinessDayCalculatorEntity {
		return entity.NewGrammarEntity(client, entopts)
	}
	core.NewIpnEntityFunc = func(client *core.BusinessDayCalculatorSDK, entopts map[string]any) core.BusinessDayCalculatorEntity {
		return entity.NewIpnEntity(client, entopts)
	}
	core.NewRedactEntityFunc = func(client *core.BusinessDayCalculatorSDK, entopts map[string]any) core.BusinessDayCalculatorEntity {
		return entity.NewRedactEntity(client, entopts)
	}
	core.NewSslEntityFunc = func(client *core.BusinessDayCalculatorSDK, entopts map[string]any) core.BusinessDayCalculatorEntity {
		return entity.NewSslEntity(client, entopts)
	}
	core.NewUtilityEntityFunc = func(client *core.BusinessDayCalculatorSDK, entopts map[string]any) core.BusinessDayCalculatorEntity {
		return entity.NewUtilityEntity(client, entopts)
	}
	core.NewWhoiEntityFunc = func(client *core.BusinessDayCalculatorSDK, entopts map[string]any) core.BusinessDayCalculatorEntity {
		return entity.NewWhoiEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewBusinessDayCalculatorSDK = core.NewBusinessDayCalculatorSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig
var SharedConfig = core.SharedConfig

// No-arg convenience constructors. Go has no default-argument syntax,
// so these aliases let callers write `sdk.New()` / `sdk.Test()`
// instead of `sdk.NewBusinessDayCalculatorSDK(nil)` / `sdk.TestSDK(nil, nil)`
// for the common no-options case.
func New() *BusinessDayCalculatorSDK  { return NewBusinessDayCalculatorSDK(nil) }
func Test() *BusinessDayCalculatorSDK { return TestSDK(nil, nil) }
var NewBaseFeature = feature.NewBaseFeature
var NewRatelimitFeature = feature.NewRatelimitFeature
var NewRetryFeature = feature.NewRetryFeature
var NewTestFeature = feature.NewTestFeature
var NewTimeoutFeature = feature.NewTimeoutFeature
