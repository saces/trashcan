# PusherData1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | **string** | The format to send notifications in to Push Gateways if the &#x60;&#x60;kind&#x60;&#x60; is &#x60;&#x60;http&#x60;&#x60;. The details about what fields the homeserver should send to the push gateway are defined in the &#x60;Push Gateway Specification&#x60;_. Currently the only format available is &#39;event_id_only&#39;. | [optional] 
**Url** | **string** | Required if &#x60;&#x60;kind&#x60;&#x60; is &#x60;&#x60;http&#x60;&#x60;. The URL to use to send notifications to. MUST be an HTTPS URL with a path of  &#x60;&#x60;/_matrix/push/v1/notify&#x60;&#x60;. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


