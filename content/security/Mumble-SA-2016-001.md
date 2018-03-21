---
title: Mumble Security Advisory 2016-001
---

ID:              Mumble-SA-2016-001 ([txt](../Mumble-SA-2016-001.txt), [sig](../Mumble-SA-2016-001.txt.sig), [Blog post](/blog/mumble-1.2.13/))

Date:            January 2016

CVE Reference:   Not assigned

Product:         Murmur

Mumble Website:  https://www.mumble.info

Permalink:       https://www.mumble.info/security/mumble-sa-2016-001

Last Updated:    2016-01-10

## Vulnerability

It is possible to cause Murmur on Windows to stop receiving UDP packets by sending a crafted UDP packet.

## Affected versions and configurations
All 1.2.x versions of Murmur are affected when running on a Windows system.

## Mitigation
It is possible to mitigate this issue by disallowing UDP packets of length 0 to reach Murmur, for example via a firewall.

## Details

Murmur on Windows uses a slightly different code path than Murmur on Unix-like systems when receiving UDP packets.

On Unix-like systems, the UDP receive path is inside a loop.

On Windows, it is not directly inside a loop.

However, the code path that handles error cases from the "recvfrom()" function is written to expect a loop, by making use of "break" and "continue". This code is shared between both Windows and Unix-like systems.

Because of this mismatch, it was possible to make Murmur on Windows break out of the outermost loop in the function that handles receiving UDP packets. The outermost loop is the main loop of the UDP receive code path in Murmur. Breaking out of this loop would cause all UDP processing to terminate.

A Murmur server that cannot receive UDP packets will lose two pieces of functionality:

- The ability to route voice packets between users via UDP. This means that clients cannot use UDP for voice communication. This can lead to temporary loss of service until the connected clients determine that no UDP transport is available. When no UDP connectivity is available, Mumble clients will fall back and use the TCP/TLS as a voice transport instead. Once clients have done this, service is restored.

- The ability to respond to ping packets.
  This means that users cannot see the ping time or user count when viewing the server in Mumble's connection dialog.

It is possible to regain the ability to receive UDP packets when a new user connects to the server. This is because a new client connection causes Murmur to ensure the UDP receive loop is running if it is not already.

## Credits

This issue was reported to the Mumble security contact by LuaMilkshake on 2016-01-04.

## Fix

A fix for this issue has been released in Mumble 1.2.13.

It converts the UDP receive path in Murmur for Windows to use a loop construct, like Murmur on Unix-like systems does.

A patch that applies on top of Mumble 1.2.12 is attached inline below:

```
diff --git a/src/murmur/Server.cpp b/src/murmur/Server.cpp
index 5c7ef94..e2d129f 100644
--- a/src/murmur/Server.cpp
+++ b/src/murmur/Server.cpp
@@ -664,7 +664,7 @@ void Server::run() {

 				int sock = fds[i].fd;
 #else
-		{
+		for (int i=0;i<1;++i) {
 			{
 				DWORD ret = WaitForMultipleObjects(nfds, events, FALSE, INFINITE);
 				if (ret == (WAIT_OBJECT_0 + nfds - 1)) {
```
