# InlineObject18

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | [**Filter**](Filter.md) |  | [optional] 
**IncludeAllNetworks** | **bool** | Whether or not to include all known networks/protocols from application services on the homeserver. Defaults to false. | [optional] 
**Limit** | **int32** | Limit the number of results returned. | [optional] 
**Since** | **string** | A pagination token from a previous request, allowing clients to get the next (or previous) batch of rooms.  The direction of pagination is specified solely by which token is supplied, rather than via an explicit flag. | [optional] 
**ThirdPartyInstanceId** | **string** | The specific third party network/protocol to request from the homeserver. Can only be used if &#x60;&#x60;include_all_networks&#x60;&#x60; is false. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


