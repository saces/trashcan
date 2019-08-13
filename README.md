# Go API client for openapi

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: unstable
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./openapi"
```

## Documentation for API Endpoints

All URIs are relative to *https://matrix.org*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ApplicationServiceRoomDirectoryManagementApi* | [**UpdateAppserviceRoomDirectoryVsibility**](docs/ApplicationServiceRoomDirectoryManagementApi.md#updateappserviceroomdirectoryvsibility) | **Put** /_matrix/client/unstable/directory/list/appservice/{networkId}/{roomId} | Updates a room&#39;s visibility in the application service&#39;s room directory.
*CapabilitiesApi* | [**GetCapabilities**](docs/CapabilitiesApi.md#getcapabilities) | **Get** /_matrix/client/unstable/capabilities | Gets information about the server&#39;s capabilities.
*DeviceManagementApi* | [**DeleteDevice**](docs/DeviceManagementApi.md#deletedevice) | **Delete** /_matrix/client/unstable/devices/{deviceId} | Delete a device
*DeviceManagementApi* | [**DeleteDevices**](docs/DeviceManagementApi.md#deletedevices) | **Post** /_matrix/client/unstable/delete_devices | Bulk deletion of devices
*DeviceManagementApi* | [**GetDevice**](docs/DeviceManagementApi.md#getdevice) | **Get** /_matrix/client/unstable/devices/{deviceId} | Get a single device
*DeviceManagementApi* | [**GetDevices**](docs/DeviceManagementApi.md#getdevices) | **Get** /_matrix/client/unstable/devices | List registered devices for the current user
*DeviceManagementApi* | [**UpdateDevice**](docs/DeviceManagementApi.md#updatedevice) | **Put** /_matrix/client/unstable/devices/{deviceId} | Update a device
*EndToEndEncryptionApi* | [**ClaimKeys**](docs/EndToEndEncryptionApi.md#claimkeys) | **Post** /_matrix/client/unstable/keys/claim | Claim one-time encryption keys.
*EndToEndEncryptionApi* | [**GetKeysChanges**](docs/EndToEndEncryptionApi.md#getkeyschanges) | **Get** /_matrix/client/unstable/keys/changes | Query users with recent device key updates.
*EndToEndEncryptionApi* | [**QueryKeys**](docs/EndToEndEncryptionApi.md#querykeys) | **Post** /_matrix/client/unstable/keys/query | Download device identity keys.
*EndToEndEncryptionApi* | [**UploadKeys**](docs/EndToEndEncryptionApi.md#uploadkeys) | **Post** /_matrix/client/unstable/keys/upload | Upload end-to-end encryption keys.
*MediaApi* | [**GetConfig**](docs/MediaApi.md#getconfig) | **Get** /_matrix/media/unstable/config | Get the configuration for the content repository.
*MediaApi* | [**GetContent**](docs/MediaApi.md#getcontent) | **Get** /_matrix/media/unstable/download/{serverName}/{mediaId} | Download content from the content repository.
*MediaApi* | [**GetContentOverrideName**](docs/MediaApi.md#getcontentoverridename) | **Get** /_matrix/media/unstable/download/{serverName}/{mediaId}/{fileName} | Download content from the content repository. This is the same as the download endpoint above, except permitting a desired file name.
*MediaApi* | [**GetContentThumbnail**](docs/MediaApi.md#getcontentthumbnail) | **Get** /_matrix/media/unstable/thumbnail/{serverName}/{mediaId} | Download a thumbnail of content from the content repository. See the &#x60;thumbnailing &lt;#thumbnails&gt;&#x60;_ section for more information.
*MediaApi* | [**GetUrlPreview**](docs/MediaApi.md#geturlpreview) | **Get** /_matrix/media/unstable/preview_url | Get information about a URL for a client
*MediaApi* | [**UploadContent**](docs/MediaApi.md#uploadcontent) | **Post** /_matrix/media/unstable/upload | Upload some content to the content repository.
*OpenIDApi* | [**RequestOpenIdToken**](docs/OpenIDApi.md#requestopenidtoken) | **Post** /_matrix/client/unstable/user/{userId}/openid/request_token | Get an OpenID token object to verify the requester&#39;s identity.
*PresenceApi* | [**GetPresence**](docs/PresenceApi.md#getpresence) | **Get** /_matrix/client/unstable/presence/{userId}/status | Get this user&#39;s presence state.
*PresenceApi* | [**SetPresence**](docs/PresenceApi.md#setpresence) | **Put** /_matrix/client/unstable/presence/{userId}/status | Update this user&#39;s presence state.
*PushNotificationsApi* | [**DeletePushRule**](docs/PushNotificationsApi.md#deletepushrule) | **Delete** /_matrix/client/unstable/pushrules/{scope}/{kind}/{ruleId} | Delete a push rule.
*PushNotificationsApi* | [**GetNotifications**](docs/PushNotificationsApi.md#getnotifications) | **Get** /_matrix/client/unstable/notifications | Gets a list of events that the user has been notified about
*PushNotificationsApi* | [**GetPushRule**](docs/PushNotificationsApi.md#getpushrule) | **Get** /_matrix/client/unstable/pushrules/{scope}/{kind}/{ruleId} | Retrieve a push rule.
*PushNotificationsApi* | [**GetPushRules**](docs/PushNotificationsApi.md#getpushrules) | **Get** /_matrix/client/unstable/pushrules/ | Retrieve all push rulesets.
*PushNotificationsApi* | [**GetPushers**](docs/PushNotificationsApi.md#getpushers) | **Get** /_matrix/client/unstable/pushers | Gets the current pushers for the authenticated user
*PushNotificationsApi* | [**PostPusher**](docs/PushNotificationsApi.md#postpusher) | **Post** /_matrix/client/unstable/pushers/set | Modify a pusher for this user on the homeserver.
*PushNotificationsApi* | [**SetPushRule**](docs/PushNotificationsApi.md#setpushrule) | **Put** /_matrix/client/unstable/pushrules/{scope}/{kind}/{ruleId} | Add or change a push rule.
*PushNotificationsApi* | [**SetPushRuleActions**](docs/PushNotificationsApi.md#setpushruleactions) | **Put** /_matrix/client/unstable/pushrules/{scope}/{kind}/{ruleId}/actions | Set the actions for a push rule.
*PushNotificationsApi* | [**SetPushRuleEnabled**](docs/PushNotificationsApi.md#setpushruleenabled) | **Put** /_matrix/client/unstable/pushrules/{scope}/{kind}/{ruleId}/enabled | Enable or disable a push rule.
*ReadMarkersApi* | [**SetReadMarker**](docs/ReadMarkersApi.md#setreadmarker) | **Post** /_matrix/client/unstable/rooms/{roomId}/read_markers | Set the position of the read marker for a room.
*ReportingContentApi* | [**ReportContent**](docs/ReportingContentApi.md#reportcontent) | **Post** /_matrix/client/unstable/rooms/{roomId}/report/{eventId} | Reports an event as inappropriate.
*RoomCreationApi* | [**CreateRoom**](docs/RoomCreationApi.md#createroom) | **Post** /_matrix/client/unstable/createRoom | Create a new room
*RoomDirectoryApi* | [**DeleteRoomAlias**](docs/RoomDirectoryApi.md#deleteroomalias) | **Delete** /_matrix/client/unstable/directory/room/{roomAlias} | Remove a mapping of room alias to room ID.
*RoomDirectoryApi* | [**GetRoomIdByAlias**](docs/RoomDirectoryApi.md#getroomidbyalias) | **Get** /_matrix/client/unstable/directory/room/{roomAlias} | Get the room ID corresponding to this room alias.
*RoomDirectoryApi* | [**SetRoomAlias**](docs/RoomDirectoryApi.md#setroomalias) | **Put** /_matrix/client/unstable/directory/room/{roomAlias} | Create a new mapping from room alias to room ID.
*RoomDiscoveryApi* | [**GetPublicRooms**](docs/RoomDiscoveryApi.md#getpublicrooms) | **Get** /_matrix/client/unstable/publicRooms | Lists the public rooms on the server.
*RoomDiscoveryApi* | [**QueryPublicRooms**](docs/RoomDiscoveryApi.md#querypublicrooms) | **Post** /_matrix/client/unstable/publicRooms | Lists the public rooms on the server with optional filter.
*RoomMembershipApi* | [**Ban**](docs/RoomMembershipApi.md#ban) | **Post** /_matrix/client/unstable/rooms/{roomId}/ban | Ban a user in the room.
*RoomMembershipApi* | [**ForgetRoom**](docs/RoomMembershipApi.md#forgetroom) | **Post** /_matrix/client/unstable/rooms/{roomId}/forget | Stop the requesting user remembering about a particular room.
*RoomMembershipApi* | [**GetJoinedRooms**](docs/RoomMembershipApi.md#getjoinedrooms) | **Get** /_matrix/client/unstable/joined_rooms | Lists the user&#39;s current rooms.
*RoomMembershipApi* | [**InviteBy3PID**](docs/RoomMembershipApi.md#inviteby3pid) | **Post** /_matrix/client/unstable/rooms/{roomId}/invite | Invite a user to participate in a particular room.
*RoomMembershipApi* | [**InviteUser**](docs/RoomMembershipApi.md#inviteuser) | **Post** /_matrix/client/unstable/rooms/{roomId}/invite  | Invite a user to participate in a particular room.
*RoomMembershipApi* | [**JoinRoom**](docs/RoomMembershipApi.md#joinroom) | **Post** /_matrix/client/unstable/join/{roomIdOrAlias} | Start the requesting user participating in a particular room.
*RoomMembershipApi* | [**JoinRoomById**](docs/RoomMembershipApi.md#joinroombyid) | **Post** /_matrix/client/unstable/rooms/{roomId}/join | Start the requesting user participating in a particular room.
*RoomMembershipApi* | [**Kick**](docs/RoomMembershipApi.md#kick) | **Post** /_matrix/client/unstable/rooms/{roomId}/kick | Kick a user from the room.
*RoomMembershipApi* | [**LeaveRoom**](docs/RoomMembershipApi.md#leaveroom) | **Post** /_matrix/client/unstable/rooms/{roomId}/leave | Stop the requesting user participating in a particular room.
*RoomMembershipApi* | [**Unban**](docs/RoomMembershipApi.md#unban) | **Post** /_matrix/client/unstable/rooms/{roomId}/unban | Unban a user from the room.
*RoomParticipationApi* | [**DefineFilter**](docs/RoomParticipationApi.md#definefilter) | **Post** /_matrix/client/unstable/user/{userId}/filter | Upload a new filter.
*RoomParticipationApi* | [**GetEventContext**](docs/RoomParticipationApi.md#geteventcontext) | **Get** /_matrix/client/unstable/rooms/{roomId}/context/{eventId} | Get events and state around the specified event.
*RoomParticipationApi* | [**GetEvents**](docs/RoomParticipationApi.md#getevents) | **Get** /_matrix/client/unstable/events | Listen on the event stream.
*RoomParticipationApi* | [**GetFilter**](docs/RoomParticipationApi.md#getfilter) | **Get** /_matrix/client/unstable/user/{userId}/filter/{filterId} | Download a filter
*RoomParticipationApi* | [**GetJoinedMembersByRoom**](docs/RoomParticipationApi.md#getjoinedmembersbyroom) | **Get** /_matrix/client/unstable/rooms/{roomId}/joined_members | Gets the list of currently joined users and their profile data.
*RoomParticipationApi* | [**GetMembersByRoom**](docs/RoomParticipationApi.md#getmembersbyroom) | **Get** /_matrix/client/unstable/rooms/{roomId}/members | Get the m.room.member events for the room.
*RoomParticipationApi* | [**GetOneEvent**](docs/RoomParticipationApi.md#getoneevent) | **Get** /_matrix/client/unstable/events/{eventId} | Get a single event by event ID.
*RoomParticipationApi* | [**GetOneRoomEvent**](docs/RoomParticipationApi.md#getoneroomevent) | **Get** /_matrix/client/unstable/rooms/{roomId}/event/{eventId} | Get a single event by event ID.
*RoomParticipationApi* | [**GetRoomEvents**](docs/RoomParticipationApi.md#getroomevents) | **Get** /_matrix/client/unstable/rooms/{roomId}/messages | Get a list of events for this room
*RoomParticipationApi* | [**GetRoomState**](docs/RoomParticipationApi.md#getroomstate) | **Get** /_matrix/client/unstable/rooms/{roomId}/state | Get all state events in the current state of a room.
*RoomParticipationApi* | [**GetRoomStateWithKey**](docs/RoomParticipationApi.md#getroomstatewithkey) | **Get** /_matrix/client/unstable/rooms/{roomId}/state/{eventType}/{stateKey} | Get the state identified by the type and key.
*RoomParticipationApi* | [**InitialSync**](docs/RoomParticipationApi.md#initialsync) | **Get** /_matrix/client/unstable/initialSync | Get the user&#39;s current state.
*RoomParticipationApi* | [**PostReceipt**](docs/RoomParticipationApi.md#postreceipt) | **Post** /_matrix/client/unstable/rooms/{roomId}/receipt/{receiptType}/{eventId} | Send a receipt for the given event ID.
*RoomParticipationApi* | [**RedactEvent**](docs/RoomParticipationApi.md#redactevent) | **Put** /_matrix/client/unstable/rooms/{roomId}/redact/{eventId}/{txnId} | Strips all non-integrity-critical information out of an event.
*RoomParticipationApi* | [**RoomInitialSync**](docs/RoomParticipationApi.md#roominitialsync) | **Get** /_matrix/client/unstable/rooms/{roomId}/initialSync | Snapshot the current state of a room and its most recent messages.
*RoomParticipationApi* | [**SendMessage**](docs/RoomParticipationApi.md#sendmessage) | **Put** /_matrix/client/unstable/rooms/{roomId}/send/{eventType}/{txnId} | Send a message event to the given room.
*RoomParticipationApi* | [**SetRoomStateWithKey**](docs/RoomParticipationApi.md#setroomstatewithkey) | **Put** /_matrix/client/unstable/rooms/{roomId}/state/{eventType}/{stateKey} | Send a state event to the given room.
*RoomParticipationApi* | [**SetTyping**](docs/RoomParticipationApi.md#settyping) | **Put** /_matrix/client/unstable/rooms/{roomId}/typing/{userId} | Informs the server that the user has started or stopped typing.
*RoomParticipationApi* | [**Sync**](docs/RoomParticipationApi.md#sync) | **Get** /_matrix/client/unstable/sync | Synchronise the client&#39;s state and receive new messages.
*RoomUgpradesApi* | [**UpgradeRoom**](docs/RoomUgpradesApi.md#upgraderoom) | **Post** /_matrix/client/unstable/rooms/{roomId}/upgrade | Upgrades a room to a new room version.
*SearchApi* | [**Search**](docs/SearchApi.md#search) | **Post** /_matrix/client/unstable/search | Perform a server-side search.
*SendToDeviceMessagingApi* | [**SendToDevice**](docs/SendToDeviceMessagingApi.md#sendtodevice) | **Put** /_matrix/client/unstable/sendToDevice/{eventType}/{txnId} | Send an event to a given set of devices.
*ServerAdministrationApi* | [**GetVersions**](docs/ServerAdministrationApi.md#getversions) | **Get** /_matrix/client/versions | Gets the versions of the specification supported by the server.
*ServerAdministrationApi* | [**GetWellknown**](docs/ServerAdministrationApi.md#getwellknown) | **Get** /.well-known/matrix/client | Gets Matrix server discovery information about the domain.
*ServerAdministrationApi* | [**GetWhoIs**](docs/ServerAdministrationApi.md#getwhois) | **Get** /_matrix/client/unstable/admin/whois/{userId} | Gets information about a particular user.
*SessionManagementApi* | [**GetLoginFlows**](docs/SessionManagementApi.md#getloginflows) | **Get** /_matrix/client/unstable/login | Get the supported login types to authenticate users
*SessionManagementApi* | [**Login**](docs/SessionManagementApi.md#login) | **Post** /_matrix/client/unstable/login | Authenticates the user.
*SessionManagementApi* | [**Logout**](docs/SessionManagementApi.md#logout) | **Post** /_matrix/client/unstable/logout | Invalidates a user access token
*SessionManagementApi* | [**LogoutAll**](docs/SessionManagementApi.md#logoutall) | **Post** /_matrix/client/unstable/logout/all | Invalidates all access tokens for a user
*UserDataApi* | [**ChangePassword**](docs/UserDataApi.md#changepassword) | **Post** /_matrix/client/unstable/account/password | Changes a user&#39;s password.
*UserDataApi* | [**CheckUsernameAvailability**](docs/UserDataApi.md#checkusernameavailability) | **Get** /_matrix/client/unstable/register/available | Checks to see if a username is available on the server.
*UserDataApi* | [**DeactivateAccount**](docs/UserDataApi.md#deactivateaccount) | **Post** /_matrix/client/unstable/account/deactivate | Deactivate a user&#39;s account.
*UserDataApi* | [**Delete3pidFromAccount**](docs/UserDataApi.md#delete3pidfromaccount) | **Post** /_matrix/client/unstable/account/3pid/delete | Deletes a third party identifier from the user&#39;s account
*UserDataApi* | [**DeleteRoomTag**](docs/UserDataApi.md#deleteroomtag) | **Delete** /_matrix/client/unstable/user/{userId}/rooms/{roomId}/tags/{tag} | Remove a tag from the room.
*UserDataApi* | [**GetAccount3PIDs**](docs/UserDataApi.md#getaccount3pids) | **Get** /_matrix/client/unstable/account/3pid | Gets a list of a user&#39;s third party identifiers.
*UserDataApi* | [**GetAccountData**](docs/UserDataApi.md#getaccountdata) | **Get** /_matrix/client/unstable/user/{userId}/account_data/{type} | Get some account_data for the user.
*UserDataApi* | [**GetAccountDataPerRoom**](docs/UserDataApi.md#getaccountdataperroom) | **Get** /_matrix/client/unstable/user/{userId}/rooms/{roomId}/account_data/{type} | Get some account_data for the user.
*UserDataApi* | [**GetAvatarUrl**](docs/UserDataApi.md#getavatarurl) | **Get** /_matrix/client/unstable/profile/{userId}/avatar_url | Get the user&#39;s avatar URL.
*UserDataApi* | [**GetDisplayName**](docs/UserDataApi.md#getdisplayname) | **Get** /_matrix/client/unstable/profile/{userId}/displayname | Get the user&#39;s display name.
*UserDataApi* | [**GetRoomTags**](docs/UserDataApi.md#getroomtags) | **Get** /_matrix/client/unstable/user/{userId}/rooms/{roomId}/tags | List the tags for a room.
*UserDataApi* | [**GetTokenOwner**](docs/UserDataApi.md#gettokenowner) | **Get** /_matrix/client/unstable/account/whoami | Gets information about the owner of an access token.
*UserDataApi* | [**GetUserProfile**](docs/UserDataApi.md#getuserprofile) | **Get** /_matrix/client/unstable/profile/{userId} | Get this user&#39;s profile information.
*UserDataApi* | [**Post3PIDs**](docs/UserDataApi.md#post3pids) | **Post** /_matrix/client/unstable/account/3pid | Adds contact information to the user&#39;s account.
*UserDataApi* | [**Register**](docs/UserDataApi.md#register) | **Post** /_matrix/client/unstable/register | Register for an account on this homeserver.
*UserDataApi* | [**SearchUserDirectory**](docs/UserDataApi.md#searchuserdirectory) | **Post** /_matrix/client/unstable/user_directory/search | Searches the user directory.
*UserDataApi* | [**SetAccountData**](docs/UserDataApi.md#setaccountdata) | **Put** /_matrix/client/unstable/user/{userId}/account_data/{type} | Set some account_data for the user.
*UserDataApi* | [**SetAccountDataPerRoom**](docs/UserDataApi.md#setaccountdataperroom) | **Put** /_matrix/client/unstable/user/{userId}/rooms/{roomId}/account_data/{type} | Set some account_data for the user.
*UserDataApi* | [**SetAvatarUrl**](docs/UserDataApi.md#setavatarurl) | **Put** /_matrix/client/unstable/profile/{userId}/avatar_url | Set the user&#39;s avatar URL.
*UserDataApi* | [**SetDisplayName**](docs/UserDataApi.md#setdisplayname) | **Put** /_matrix/client/unstable/profile/{userId}/displayname | Set the user&#39;s display name.
*UserDataApi* | [**SetRoomTag**](docs/UserDataApi.md#setroomtag) | **Put** /_matrix/client/unstable/user/{userId}/rooms/{roomId}/tags/{tag} | Add a tag to a room.
*VOIPApi* | [**GetTurnServer**](docs/VOIPApi.md#getturnserver) | **Get** /_matrix/client/unstable/voip/turnServer | Obtain TURN server credentials.


## Documentation For Models

 - [AuthenticationResponse](docs/AuthenticationResponse.md)
 - [AuthenticationResponseFlows](docs/AuthenticationResponseFlows.md)
 - [Body](docs/Body.md)
 - [Categories](docs/Categories.md)
 - [ConnectionInfo](docs/ConnectionInfo.md)
 - [DeviceInfo](docs/DeviceInfo.md)
 - [Event](docs/Event.md)
 - [EventContext](docs/EventContext.md)
 - [Filter](docs/Filter.md)
 - [Group](docs/Group.md)
 - [Groupings](docs/Groupings.md)
 - [IncludeEventContext](docs/IncludeEventContext.md)
 - [InlineObject](docs/InlineObject.md)
 - [InlineObject1](docs/InlineObject1.md)
 - [InlineObject10](docs/InlineObject10.md)
 - [InlineObject11](docs/InlineObject11.md)
 - [InlineObject12](docs/InlineObject12.md)
 - [InlineObject13](docs/InlineObject13.md)
 - [InlineObject14](docs/InlineObject14.md)
 - [InlineObject15](docs/InlineObject15.md)
 - [InlineObject16](docs/InlineObject16.md)
 - [InlineObject17](docs/InlineObject17.md)
 - [InlineObject18](docs/InlineObject18.md)
 - [InlineObject19](docs/InlineObject19.md)
 - [InlineObject2](docs/InlineObject2.md)
 - [InlineObject20](docs/InlineObject20.md)
 - [InlineObject21](docs/InlineObject21.md)
 - [InlineObject22](docs/InlineObject22.md)
 - [InlineObject23](docs/InlineObject23.md)
 - [InlineObject24](docs/InlineObject24.md)
 - [InlineObject25](docs/InlineObject25.md)
 - [InlineObject26](docs/InlineObject26.md)
 - [InlineObject27](docs/InlineObject27.md)
 - [InlineObject28](docs/InlineObject28.md)
 - [InlineObject29](docs/InlineObject29.md)
 - [InlineObject3](docs/InlineObject3.md)
 - [InlineObject30](docs/InlineObject30.md)
 - [InlineObject31](docs/InlineObject31.md)
 - [InlineObject32](docs/InlineObject32.md)
 - [InlineObject33](docs/InlineObject33.md)
 - [InlineObject34](docs/InlineObject34.md)
 - [InlineObject35](docs/InlineObject35.md)
 - [InlineObject36](docs/InlineObject36.md)
 - [InlineObject37](docs/InlineObject37.md)
 - [InlineObject4](docs/InlineObject4.md)
 - [InlineObject5](docs/InlineObject5.md)
 - [InlineObject6](docs/InlineObject6.md)
 - [InlineObject7](docs/InlineObject7.md)
 - [InlineObject8](docs/InlineObject8.md)
 - [InlineObject9](docs/InlineObject9.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [InlineResponse20010](docs/InlineResponse20010.md)
 - [InlineResponse20011](docs/InlineResponse20011.md)
 - [InlineResponse20012](docs/InlineResponse20012.md)
 - [InlineResponse20013](docs/InlineResponse20013.md)
 - [InlineResponse20014](docs/InlineResponse20014.md)
 - [InlineResponse20015](docs/InlineResponse20015.md)
 - [InlineResponse20016](docs/InlineResponse20016.md)
 - [InlineResponse20017](docs/InlineResponse20017.md)
 - [InlineResponse20018](docs/InlineResponse20018.md)
 - [InlineResponse20019](docs/InlineResponse20019.md)
 - [InlineResponse2002](docs/InlineResponse2002.md)
 - [InlineResponse20020](docs/InlineResponse20020.md)
 - [InlineResponse20021](docs/InlineResponse20021.md)
 - [InlineResponse20022](docs/InlineResponse20022.md)
 - [InlineResponse20023](docs/InlineResponse20023.md)
 - [InlineResponse20024](docs/InlineResponse20024.md)
 - [InlineResponse20025](docs/InlineResponse20025.md)
 - [InlineResponse20026](docs/InlineResponse20026.md)
 - [InlineResponse20027](docs/InlineResponse20027.md)
 - [InlineResponse20028](docs/InlineResponse20028.md)
 - [InlineResponse20029](docs/InlineResponse20029.md)
 - [InlineResponse2003](docs/InlineResponse2003.md)
 - [InlineResponse20030](docs/InlineResponse20030.md)
 - [InlineResponse20031](docs/InlineResponse20031.md)
 - [InlineResponse20032](docs/InlineResponse20032.md)
 - [InlineResponse20033](docs/InlineResponse20033.md)
 - [InlineResponse20034](docs/InlineResponse20034.md)
 - [InlineResponse20035](docs/InlineResponse20035.md)
 - [InlineResponse20036](docs/InlineResponse20036.md)
 - [InlineResponse20037](docs/InlineResponse20037.md)
 - [InlineResponse20038](docs/InlineResponse20038.md)
 - [InlineResponse20039](docs/InlineResponse20039.md)
 - [InlineResponse2004](docs/InlineResponse2004.md)
 - [InlineResponse20040](docs/InlineResponse20040.md)
 - [InlineResponse20041](docs/InlineResponse20041.md)
 - [InlineResponse20042](docs/InlineResponse20042.md)
 - [InlineResponse20043](docs/InlineResponse20043.md)
 - [InlineResponse20044](docs/InlineResponse20044.md)
 - [InlineResponse20045](docs/InlineResponse20045.md)
 - [InlineResponse20046](docs/InlineResponse20046.md)
 - [InlineResponse20047](docs/InlineResponse20047.md)
 - [InlineResponse20048](docs/InlineResponse20048.md)
 - [InlineResponse20049](docs/InlineResponse20049.md)
 - [InlineResponse2005](docs/InlineResponse2005.md)
 - [InlineResponse2006](docs/InlineResponse2006.md)
 - [InlineResponse2007](docs/InlineResponse2007.md)
 - [InlineResponse2008](docs/InlineResponse2008.md)
 - [InlineResponse2009](docs/InlineResponse2009.md)
 - [InlineResponse403](docs/InlineResponse403.md)
 - [InlineResponse429](docs/InlineResponse429.md)
 - [Invite3pid](docs/Invite3pid.md)
 - [InviteState](docs/InviteState.md)
 - [InvitedRoom](docs/InvitedRoom.md)
 - [JoinedRoom](docs/JoinedRoom.md)
 - [LeftRoom](docs/LeftRoom.md)
 - [LoginFlow](docs/LoginFlow.md)
 - [Notification](docs/Notification.md)
 - [PaginationChunk](docs/PaginationChunk.md)
 - [PublicRoomsChunk](docs/PublicRoomsChunk.md)
 - [Pusher](docs/Pusher.md)
 - [PusherData](docs/PusherData.md)
 - [PusherData1](docs/PusherData1.md)
 - [Result](docs/Result.md)
 - [ResultCategories](docs/ResultCategories.md)
 - [ResultRoomEvents](docs/ResultRoomEvents.md)
 - [Results](docs/Results.md)
 - [RoomEvent](docs/RoomEvent.md)
 - [RoomEventsCriteria](docs/RoomEventsCriteria.md)
 - [RoomFilter](docs/RoomFilter.md)
 - [RoomInfo](docs/RoomInfo.md)
 - [RoomInfo1](docs/RoomInfo1.md)
 - [RoomMember](docs/RoomMember.md)
 - [RoomSummary](docs/RoomSummary.md)
 - [Rooms](docs/Rooms.md)
 - [SessionInfo](docs/SessionInfo.md)
 - [StateEvent](docs/StateEvent.md)
 - [StrippedState](docs/StrippedState.md)
 - [Tag](docs/Tag.md)
 - [ThirdPartyIdentifier](docs/ThirdPartyIdentifier.md)
 - [ThirdPartySigned](docs/ThirdPartySigned.md)
 - [ThreePidCredentials](docs/ThreePidCredentials.md)
 - [UnreadNotificationCounts](docs/UnreadNotificationCounts.md)
 - [User](docs/User.md)
 - [UserIdentifier](docs/UserIdentifier.md)
 - [UserProfile](docs/UserProfile.md)


## Documentation For Authorization



## accessToken

- **Type**: API key

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
    Key: "APIKEY",
    Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```


## Author


