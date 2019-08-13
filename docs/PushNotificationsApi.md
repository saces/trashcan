# \PushNotificationsApi

All URIs are relative to *https://matrix.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePushRule**](PushNotificationsApi.md#DeletePushRule) | **Delete** /_matrix/client/unstable/pushrules/{scope}/{kind}/{ruleId} | Delete a push rule.
[**GetNotifications**](PushNotificationsApi.md#GetNotifications) | **Get** /_matrix/client/unstable/notifications | Gets a list of events that the user has been notified about
[**GetPushRule**](PushNotificationsApi.md#GetPushRule) | **Get** /_matrix/client/unstable/pushrules/{scope}/{kind}/{ruleId} | Retrieve a push rule.
[**GetPushRules**](PushNotificationsApi.md#GetPushRules) | **Get** /_matrix/client/unstable/pushrules/ | Retrieve all push rulesets.
[**GetPushers**](PushNotificationsApi.md#GetPushers) | **Get** /_matrix/client/unstable/pushers | Gets the current pushers for the authenticated user
[**PostPusher**](PushNotificationsApi.md#PostPusher) | **Post** /_matrix/client/unstable/pushers/set | Modify a pusher for this user on the homeserver.
[**SetPushRule**](PushNotificationsApi.md#SetPushRule) | **Put** /_matrix/client/unstable/pushrules/{scope}/{kind}/{ruleId} | Add or change a push rule.
[**SetPushRuleActions**](PushNotificationsApi.md#SetPushRuleActions) | **Put** /_matrix/client/unstable/pushrules/{scope}/{kind}/{ruleId}/actions | Set the actions for a push rule.
[**SetPushRuleEnabled**](PushNotificationsApi.md#SetPushRuleEnabled) | **Put** /_matrix/client/unstable/pushrules/{scope}/{kind}/{ruleId}/enabled | Enable or disable a push rule.



## DeletePushRule

> map[string]interface{} DeletePushRule(ctx, scope, kind, ruleId)
Delete a push rule.

This endpoint removes the push rule defined in the path.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scope** | **string**| &#x60;&#x60;global&#x60;&#x60; to specify global rules. | 
**kind** | **string**| The kind of rule  | 
**ruleId** | **string**| The identifier for the rule.  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotifications

> InlineResponse20021 GetNotifications(ctx, optional)
Gets a list of events that the user has been notified about

This API is used to paginate through the list of events that the user has been, or would have been notified about.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetNotificationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetNotificationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **optional.String**| Pagination token given to retrieve the next set of events. | 
 **limit** | **optional.Int32**| Limit on the number of events to return in this request. | 
 **only** | **optional.String**| Allows basic filtering of events returned. Supply &#x60;&#x60;highlight&#x60;&#x60; to return only events where the notification had the highlight tweak set. | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPushRule

> InlineResponse20030 GetPushRule(ctx, scope, kind, ruleId)
Retrieve a push rule.

Retrieve a single specified push rule.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scope** | **string**| &#x60;&#x60;global&#x60;&#x60; to specify global rules. | 
**kind** | **string**| The kind of rule  | 
**ruleId** | **string**| The identifier for the rule.  | 

### Return type

[**InlineResponse20030**](inline_response_200_30.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPushRules

> InlineResponse20029 GetPushRules(ctx, )
Retrieve all push rulesets.

Retrieve all push rulesets for this user. Clients can \"drill-down\" on the rulesets by suffixing a ``scope`` to this path e.g. ``/pushrules/global/``. This will return a subset of this data under the specified key e.g. the ``global`` key.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse20029**](inline_response_200_29.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPushers

> InlineResponse20028 GetPushers(ctx, )
Gets the current pushers for the authenticated user

Gets all currently active pushers for the authenticated user.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse20028**](inline_response_200_28.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPusher

> map[string]interface{} PostPusher(ctx, pusher)
Modify a pusher for this user on the homeserver.

This endpoint allows the creation, modification and deletion of `pushers`_ for this user ID. The behaviour of this endpoint varies depending on the values in the JSON body.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pusher** | [**InlineObject19**](InlineObject19.md)|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPushRule

> map[string]interface{} SetPushRule(ctx, scope, kind, ruleId, pushrule, optional)
Add or change a push rule.

This endpoint allows the creation, modification and deletion of pushers for this user ID. The behaviour of this endpoint varies depending on the values in the JSON body.  When creating push rules, they MUST be enabled by default.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scope** | **string**| &#x60;&#x60;global&#x60;&#x60; to specify global rules. | 
**kind** | **string**| The kind of rule  | 
**ruleId** | **string**| The identifier for the rule.  | 
**pushrule** | [**InlineObject20**](InlineObject20.md)|  | 
 **optional** | ***SetPushRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SetPushRuleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **before** | **optional.String**| Use &#39;before&#39; with a &#x60;&#x60;rule_id&#x60;&#x60; as its value to make the new rule the next-most important rule with respect to the given user defined rule. It is not possible to add a rule relative to a predefined server rule. | 
 **after** | **optional.String**| This makes the new rule the next-less important rule relative to the given user defined rule. It is not possible to add a rule relative to a predefined server rule. | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPushRuleActions

> map[string]interface{} SetPushRuleActions(ctx, scope, kind, ruleId, body)
Set the actions for a push rule.

This endpoint allows clients to change the actions of a push rule. This can be used to change the actions of builtin rules.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scope** | **string**| &#x60;&#x60;global&#x60;&#x60; to specify global rules. | 
**kind** | **string**| The kind of rule  | 
**ruleId** | **string**| The identifier for the rule.  | 
**body** | [**InlineObject21**](InlineObject21.md)|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPushRuleEnabled

> map[string]interface{} SetPushRuleEnabled(ctx, scope, kind, ruleId, body)
Enable or disable a push rule.

This endpoint allows clients to enable or disable the specified push rule.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scope** | **string**| &#x60;&#x60;global&#x60;&#x60; to specify global rules. | 
**kind** | **string**| The kind of rule  | 
**ruleId** | **string**| The identifier for the rule.  | 
**body** | [**InlineObject22**](InlineObject22.md)|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

