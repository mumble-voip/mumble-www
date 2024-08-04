---
title: "Channel Viewer Protocol"
---

The Mumble server offers a ZeroC Ice API endpoint that allows programmatic access to the server.

The Channel Viewer Protocol (“CVP”) is a simpler JSON or XML protocol offered by other applications that bind to the
Mumble server Ice API.

Multiple provider and consumer / viewer implementations exist.

The document is encoded in UTF-8.

## CVP Provider Guidelines

- Information from the Mumble Server shall be provided as complete and unaltered as reasonably possible
- Custom fields that are not sourced directly from Mumble server shall use a `x_` name prefix to evade name clashes
- Any content-encoding headers shall announce the UTF-8 encoding being used
- Providers, Provider developers, publishers, and hosters MUST take sufficient care regarding user privacy, encryption,
  and authentication
- A provider may offer JSON, XML, or both
- Providers that offer JSON shall offer simple JSON responses as well as [JSONP](https://en.wikipedia.org/wiki/JSONP)
  with `callback` as the callback parameter

## CVP Consumer Guidelines

- A field not recognized by the client shall be silently ignored (forward-compatibility)
- Empty fields shall be interpreted as unset. A consumer shall not rely on a field that exists to have a usable value.

## Address field, authentication, and user privacy

User data MUST be protected through adequate means and authentication.

If the provider does not support delivering users' addresses at all, it shall ignore any authentication data from client
the client.

Integrations into web and website frameworks and platforms should integrate into their session, authentication, and
permission scheme to ensure deliberate data sharing over unintended leaking of information.

To ease client implementation, the server shall use a `x_addrstr` text field that contains one of:

- a fully qualified domain name like `somewhere.example.com`
- an IPv4 address in dotted quad notation like `192.168.0.1`
- an IPv6 address in colon notation like `2001:6f8:108f:1:21f:3cff:wx:yz`

## Examples

```text
Root
├── ChannelA
└── ChannelB
    ├── htmlsubtesting
    │   └── subsubtesting
    └── plainsubtesting
        └── subsubtesting
```

### JSON

#### Server

Mandatory server root object fields:

```json
{
  "id": 1,
  "name": "The Server Name",
  "root": {}
}
```

Optional server fields:

- `x_uptime`: The number of seconds this server has been running

```json
    "x_connecturl": "mumble://somewhere/",
    "x_uptime": 7156
```

#### Channel

Mandatory channel fields:

```json
{
  "id": 0,
  "parent": -1,
  "position": 0,
  "name": "Root",
  "description": "",
  "links": [],
  "users": [],
  "channels": [],
  "temporary": false
}
```

The fields `links`, `users`, and `channels` contain IDs of linked channels, users in the channel, and subchannels
respectively.

Optional channel fields:

```json
"x_connecturl": "mumble://somewhere/somechannel"
```

#### User

_Take note of the address and privacy notes in the corresponding general section._

Mandatory User fields:

```json
{
  "session": 143,
  "userid": 4,
  "name": "Yps0",
  "channel": 40,
  "deaf": false,
  "mute": false,
  "selfDeaf": false,
  "selfMute": false,
  "suppress": false
}
```

Optional User fields:

```json
{
    "address": (0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 192, 168, 0, 10),
    "x_addrstring": "192.168.0.10",
    "comment": ,
    "onlinesecs": 161,
    "idlesecs": 0,
    "bytespersec": 6580,
    "context": ,
    "identity": ,
    "recording": false,
    "prioritySpeaker": false,
    "tcponly": false,
    "x_texture": "https://example.com/avatars/mumble/texture.png",
    "os": "Win",
    "osversion": "5.1.2600.1",
    "release": "02071e",
    "version": 66051
}
```

- `x_addrstr` or `x_addrstring` shall be a human-readable string format of the address; hostname, dotted IPv4, or colon
  IPv6

#### Full JSON Example

```json
{
   "x_connecturl": "mumble://mumble.example.com?version=1.2.2",
   "x_uptime": 53289,
   "root": {
       "channels": [
           {
               "channels": [],
               "x_connecturl": "mumble://mumble.example.com/newtesting?version=1.2.2",
               "temporary": false,
               "description": "",
               "parent": 0,
               "position": 0,
               "users": [],
               "name": "ChannelA",
               "links": [],
               "id": 6
           },
           {
               "channels": [
                   {
                       "channels": [
                           {
                               "channels": [],
                               "x_connecturl": "mumble://mumble.example.com/testing/htmlsubtesting%20%C3%B6%C3%A4%C3%BC/subsubtesting?version=1.2.2",
                               "temporary": false,
                               "description": "",
                               "parent": 4,
                               "position": 0,
                               "users": [],
                               "name": "subsubtesting",
                               "links": [],
                               "id": 5
                           }
                       ],
                       "x_connecturl": "mumble://mumble.example.com/testing/htmlsubtesting%20%C3%B6%C3%A4%C3%BC?version=1.2.2",
                       "temporary": false,
                       "description": "This channel has HTML elements in its description that are evaluated as red bold Fonts!",
                       "parent": 1,
                       "position": 0,
                       "users": [
                           {
                               "comment": "",
                               "mute": false,
                               "suppress": false,
                               "selfDeaf": false,
                               "x_addrstring": "2001:6f8:108f:1:21f:3cff:wx:yz",
                               "deaf": false,
                               "selfMute": true,
                               "bytespersec": 0,
                               "session": 1,
                               "tcponly": false,
                               "address": [ 32, 1, 6, 248, 16, 143, 0, 1, 2, 31, 60, 255, w, x, y, z ],
                               "idlesecs": 33,
                               "identity": "",
                               "name": "svedrin",
                               "userid": 1,
                               "version": 66050,
                               "context": "",
                               "release": "1.2.2-2",
                               "osversion": "Debian GNU/Linux unstable (sid)",
                               "os": "X11",
                               "onlinesecs": 35,
                               "channel": 4
                           }
                       ],
                       "name": "htmlsubtesting \u00f6\u00e4\u00fc",
                       "links": [
                           2
                       ],
                       "id": 4
                   },
                   {
                       "channels": [
                           {
                               "channels": [],
                               "x_connecturl": "mumble://mumble.example.com/testing/plainsubtesting/subsubtesting?version=1.2.2",
                               "temporary": false,
                               "description": "",
                               "parent": 2,
                               "position": 0,
                               "users": [],
                               "name": "subsubtesting",
                               "links": [],
                               "id": 3
                           }
                       ],
                       "x_connecturl": "mumble://mumble.example.com/testing/plainsubtesting?version=1.2.2",
                       "temporary": false,
                       "description": "This channel has <i>HTML</i> elements in its description that are not evaluated as <span style="color: red">red <b>bold</b></span> Fonts!",
                       "parent": 1,
                       "position": 0,
                       "users": [],
                       "name": "plainsubtesting",
                       "links": [
                           4
                       ],
                       "id": 2
                   }
               ],
               "x_connecturl": "mumble://mumble.example.com/testing?version=1.2.2",
               "temporary": false,
               "description": "",
               "parent": 0,
               "position": 0,
               "users": [],
               "name": "ChannelB",
               "links": [],
               "id": 1
           }
       ],
       "x_connecturl": "mumble://mumble.example.com?version=1.2.2",
       "temporary": false,
       "description": "",
       "parent": -1,
       "position": 0,
       "users": [],
       "name": "Root",
       "links": [],
       "id": 0
   },
   "name": "The Server Name",
   "id": 1
}
```

### XML

For XML, the same principles apply as for JSON, along with the following encoding principles:

- Boolean values are represented as "true" or "false"
- Lists (e.g. for the players' IP addresses or the channel links) are space-separated numbers

The root element (`<server>`) shall declare the `xmlns` and `xmlns:xsi` namespace attributes as follows:

```xml
<?xml version="1.0" encoding="UTF-8" ?>
<server
  xmlns="http://mumble.sourceforge.net/Channel_Viewer_Protocol"
  xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
  >
```

The following element names shall be used:

- Server: `<server>`
- Channel: `<channel>`
- User: `<user>`

The User context field value MUST be base64 encoded because in the Mumble server it holds `\0` bytes.

### Full XML Example

```xml
<?xml version="1.0" encoding="UTF-8" ?>
<server
  xmlns="http://mumble.sourceforge.net/Channel_Viewer_Protocol"
  xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
  id="1"
  name="The Server Name"
  x_connecturl="mumble://mumble.example.com?version=1.2.2">
    <channel
      description=""
      id="0"
      links=""
      name="Root"
      parent="-1"
      position="0"
      temporary="false"
      x_connecturl="mumble://mumble.example.com?version=1.2.2">
        <channel
          description=""
          id="6"
          links=""
          name="ChannelA"
          parent="0"
          position="0"
          temporary="false"
          x_connecturl="mumble://mumble.example.com/newtesting?version=1.2.2"  />
        <channel
          description=""
          id="1"
          links=""
          name="ChannelB"
          parent="0"
          position="0"
          temporary="false"
          x_connecturl="mumble://mumble.example.com/testing?version=1.2.2">
            <channel
              description="This channel has <i>HTML</i> elements in its description that are evaluated as <span style=&quot;color:#ff0000&quot;>red </span><b><span style=&quot;color:#ff0000&quot;>bold</span></b> Fonts!"
              id="4"
              links="2"
              name="htmlsubtesting"
              parent="1"
              position="0"
              temporary="false"
              x_connecturl="mumble://mumble.example.com/testing/htmlsubtesting?version=1.2.2">
                <channel
                  description=""
                  id="5"
                  links=""
                  name="subsubtesting"
                  parent="4"
                  position="0"
                  temporary="false"
                  x_connecturl="mumble://mumble.example.com/testing/htmlsubtesting/subsubtesting?version=1.2.2"  />
                <user
                  address="32 1 6 248 16 143 0 1 2 31 60 255 w x y z"
                  bytespersec="0"
                  channel="4"
                  comment=""
                  context="QmF0dGxlZmllbGQgQmFkIENvbXBhbnkgMgA= "
                  deaf="false"
                  identity=""
                  idlesecs="1693"
                  mute="false"
                  name="svedrin"
                  onlinesecs="2943"
                  os="X11"
                  osversion="Debian  GNU/Linux  unstable  (sid)"
                  release="1.2.2-2"
                  selfDeaf="false"
                  selfMute="false"
                  session="7"
                  suppress="false"
                  tcponly="false"
                  userid="1"
                  version="66050"
                  x_addrstring="2001:6f8:108f:1:21f:3cff:wx:yz"  />
            </channel>
            <channel
              description="This channel has &lt;i&gt;HTML&lt;/i&gt; elements in its description that are not evaluated as &lt;span style=&quot;color: red&quot;&gt;red &lt;b&gt;bold&lt;/b&gt;&lt;/span&gt; Fonts!"
              id="2"
              links="4"
              name="plainsubtesting"
              parent="1"
              position="0"
              temporary="false"
              x_connecturl="mumble://mumble.example.com/testing/plainsubtesting?version=1.2.2">
                <channel
                  description=""
                  id="3"
                  links=""
                  name="subsubtesting"
                  parent="2"
                  position="0"
                  temporary="false"
                  x_connecturl="mumble://mumble.example.com/testing/plainsubtesting/subsubtesting?version=1.2.2"  />
            </channel>
        </channel>
    </channel>
</server>
```
