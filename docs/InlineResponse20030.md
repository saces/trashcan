# InlineResponse20030

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | **[]string** | The actions to perform when this rule is matched. | 
**Conditions** | [**[]map[string]interface{}**](map[string]interface{}.md) | The conditions that must hold true for an event in order for a rule to be applied to an event. A rule with no conditions always matches. Only applicable to &#x60;&#x60;underride&#x60;&#x60; and &#x60;&#x60;override&#x60;&#x60; rules. | [optional] 
**Default** | **bool** | Whether this is a default rule, or has been set explicitly. | 
**Enabled** | **bool** | Whether the push rule is enabled or not. | 
**Pattern** | **string** | The glob-style pattern to match against.  Only applicable to &#x60;&#x60;content&#x60;&#x60; rules. | [optional] 
**RuleId** | **string** | The ID of this rule. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


