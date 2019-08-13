# Pusher

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppDisplayName** | **string** | A string that will allow the user to identify what application owns this pusher. | 
**AppId** | **string** | This is a reverse-DNS style identifier for the application. Max length, 64 chars. | 
**Data** | [**PusherData**](PusherData.md) |  | 
**DeviceDisplayName** | **string** | A string that will allow the user to identify what device owns this pusher. | 
**Kind** | **string** | The kind of pusher. &#x60;&#x60;\&quot;http\&quot;&#x60;&#x60; is a pusher that sends HTTP pokes. | 
**Lang** | **string** | The preferred language for receiving notifications (e.g. &#39;en&#39; or &#39;en-US&#39;) | 
**ProfileTag** | **string** | This string determines which set of device specific rules this pusher executes. | [optional] 
**Pushkey** | **string** | This is a unique identifier for this pusher. See &#x60;&#x60;/set&#x60;&#x60; for more detail. Max length, 512 bytes. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


