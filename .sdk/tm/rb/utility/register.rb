# BusinessDayCalculator SDK utility registration
require_relative '../core/utility_type'
require_relative 'clean'
require_relative 'done'
require_relative 'make_error'
require_relative 'feature_add'
require_relative 'feature_hook'
require_relative 'feature_init'
require_relative 'fetcher'
require_relative 'make_fetch_def'
require_relative 'make_context'
require_relative 'make_options'
require_relative 'make_request'
require_relative 'make_response'
require_relative 'make_result'
require_relative 'make_point'
require_relative 'make_spec'
require_relative 'make_url'
require_relative 'param'
require_relative 'prepare_auth'
require_relative 'prepare_body'
require_relative 'prepare_headers'
require_relative 'prepare_method'
require_relative 'prepare_params'
require_relative 'prepare_path'
require_relative 'prepare_query'
require_relative 'graphql'
require_relative 'result_basic'
require_relative 'result_body'
require_relative 'result_headers'
require_relative 'transform_request'
require_relative 'transform_response'

BusinessDayCalculatorUtility.registrar = ->(u) {
  u.clean = BusinessDayCalculatorUtilities::Clean
  u.done = BusinessDayCalculatorUtilities::Done
  u.make_error = BusinessDayCalculatorUtilities::MakeError
  u.feature_add = BusinessDayCalculatorUtilities::FeatureAdd
  u.feature_hook = BusinessDayCalculatorUtilities::FeatureHook
  u.feature_init = BusinessDayCalculatorUtilities::FeatureInit
  u.fetcher = BusinessDayCalculatorUtilities::Fetcher
  u.make_fetch_def = BusinessDayCalculatorUtilities::MakeFetchDef
  u.make_context = BusinessDayCalculatorUtilities::MakeContext
  u.make_options = BusinessDayCalculatorUtilities::MakeOptions
  u.make_request = BusinessDayCalculatorUtilities::MakeRequest
  u.make_response = BusinessDayCalculatorUtilities::MakeResponse
  u.make_result = BusinessDayCalculatorUtilities::MakeResult
  u.make_point = BusinessDayCalculatorUtilities::MakePoint
  u.make_spec = BusinessDayCalculatorUtilities::MakeSpec
  u.make_url = BusinessDayCalculatorUtilities::MakeUrl
  u.param = BusinessDayCalculatorUtilities::Param
  u.prepare_auth = BusinessDayCalculatorUtilities::PrepareAuth
  u.prepare_body = BusinessDayCalculatorUtilities::PrepareBody
  u.prepare_headers = BusinessDayCalculatorUtilities::PrepareHeaders
  u.prepare_method = BusinessDayCalculatorUtilities::PrepareMethod
  u.prepare_params = BusinessDayCalculatorUtilities::PrepareParams
  u.prepare_path = BusinessDayCalculatorUtilities::PreparePath
  u.prepare_query = BusinessDayCalculatorUtilities::PrepareQuery
  u.graphql_body = BusinessDayCalculatorUtilities::GraphqlBody
  u.graphql_errors = BusinessDayCalculatorUtilities::GraphqlErrors
  u.result_basic = BusinessDayCalculatorUtilities::ResultBasic
  u.result_body = BusinessDayCalculatorUtilities::ResultBody
  u.result_headers = BusinessDayCalculatorUtilities::ResultHeaders
  u.transform_request = BusinessDayCalculatorUtilities::TransformRequest
  u.transform_response = BusinessDayCalculatorUtilities::TransformResponse
}
