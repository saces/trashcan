# \RoomCreationApi

All URIs are relative to *https://matrix.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoom**](RoomCreationApi.md#CreateRoom) | **Post** /_matrix/client/unstable/createRoom | Create a new room



## CreateRoom

> InlineResponse2006 CreateRoom(ctx, optional)
Create a new room

Create a new room with various configuration options.  The server MUST apply the normal state resolution rules when creating the new room, including checking power levels for each event. It MUST apply the events implied by the request in the following order:  0. A default ``m.room.power_levels`` event, giving the room creator    (and not other members) permission to send state events. Overridden    by the ``power_level_content_override`` parameter.  1. Events set by the ``preset``. Currently these are the ``m.room.join_rules``,    ``m.room.history_visibility``, and ``m.room.guest_access`` state events.  2. Events listed in ``initial_state``, in the order that they are    listed.  3. Events implied by ``name`` and ``topic`` (``m.room.name`` and ``m.room.topic``    state events).  4. Invite events implied by ``invite`` and ``invite_3pid`` (``m.room.member`` with    ``membership: invite`` and ``m.room.third_party_invite``).  The available presets do the following with respect to room state:  ========================  ==============  ======================  ================  =========          Preset           ``join_rules``  ``history_visibility``  ``guest_access``  Other ========================  ==============  ======================  ================  ========= ``private_chat``          ``invite``      ``shared``              ``can_join`` ``trusted_private_chat``  ``invite``      ``shared``              ``can_join``      All invitees are given the same power level as the room creator. ``public_chat``           ``public``      ``shared``              ``forbidden`` ========================  ==============  ======================  ================  =========  The server will create a ``m.room.create`` event in the room with the requesting user as the creator, alongside other keys provided in the ``creation_content``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateRoomOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateRoomOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of InlineObject4**](InlineObject4.md)|  | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

