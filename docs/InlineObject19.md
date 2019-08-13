# InlineObject19

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppDisplayName** | **string** | A string that will allow the user to identify what application owns this pusher. | 
**AppId** | **string** | This is a reverse-DNS style identifier for the application. It is recommended that this end with the platform, such that different platform versions get different app identifiers. Max length, 64 chars.  If the &#x60;&#x60;kind&#x60;&#x60; is &#x60;&#x60;\&quot;email\&quot;&#x60;&#x60;, this is &#x60;&#x60;\&quot;m.email\&quot;&#x60;&#x60;. | 
**Append** | **bool** | If true, the homeserver should add another pusher with the given pushkey and App ID in addition to any others with different user IDs. Otherwise, the homeserver must remove any other pushers with the same App ID and pushkey for different users. The default is &#x60;&#x60;false&#x60;&#x60;. | [optional] 
**Data** | [**PusherData1**](PusherData_1.md) |  | 
**DeviceDisplayName** | **string** | A string that will allow the user to identify what device owns this pusher. | 
**Kind** | **string** | The kind of pusher to configure. &#x60;&#x60;\&quot;http\&quot;&#x60;&#x60; makes a pusher that sends HTTP pokes. &#x60;&#x60;\&quot;email\&quot;&#x60;&#x60; makes a pusher that emails the user with unread notifications. &#x60;&#x60;null&#x60;&#x60; deletes the pusher. | 
**Lang** | **string** | The preferred language for receiving notifications (e.g. &#39;en&#39; or &#39;en-US&#39;). | 
**ProfileTag** | **string** | This string determines which set of device specific rules this pusher executes. | [optional] 
**Pushkey** | **string** | This is a unique identifier for this pusher. The value you should use for this is the routing or destination address information for the notification, for example, the APNS token for APNS or the Registration ID for GCM. If your notification client has no such concept, use any unique identifier. Max length, 512 bytes.  If the &#x60;&#x60;kind&#x60;&#x60; is &#x60;&#x60;\&quot;email\&quot;&#x60;&#x60;, this is the email address to send notifications to. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


