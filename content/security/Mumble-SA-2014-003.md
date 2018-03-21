---
title: Mumble Security Advisory 2014-003
---

[txt](Mumble-SA-2014-003.txt), [sig](Mumble-SA-2014-003.txt.sig), [Blog post](/blog/mumble-for-ios-1.2.3/)

```
Mumble Security Advisory 2014-003

ID:              Mumble-SA-2014-003
Date:            February 2014
CVE Reference:   Not assigned
Product:         MumbleKit, Mumble for iOS
Mumble Website:  http://mumble.info
Permalink:       http://mumble.info/security/Mumble-SA-2014-003.txt
Last Updated:    05-02-2014

1. Vulnerability

	A malformed Opus voice packet sent to a MumbleKit
	client (such as Mumble for iOS) could trigger a
	NULL pointer dereference. This causes a client
	crash (Denial of Service).

	This can be triggered remotely by an entity
	participating in a Mumble voice chat.

2. Affected versions and configurations

	All versions of MumbleKit with Opus support are
	affected unless they include the fix for this
	issue, which is available in the MumbleKit Git
	repository as commit
	fd190328a9b24d37382b269a5674b0c0c7a7e36d.

	Mumble for iOS version 1.1 through 1.2.2 are
	vulnerable, as they use vulnerable versions
	of MumbleKit.

3. Mitigation

	No known mitigation strategy for this issue exists.

	Version 1.2.3 and prior of Mumble's server component
	('Murmur' or 'mumble-server') does not allow the
	transmission of Opus packets, and as such a vulnerable
	client connected to a stock Murmur server that runs
	version 1.2.3 or prior should not be affected by this
	issue.

	Note however that since Mumble is a centralized VoIP
	system, a modified server could potentially also
	trigger malformed Opus packets to be sent to clients
	of its choosing, thus triggering this issue.

4. Details

	Mumble's Opus voice packets are serialized as a buffer
	with a length-prefix using Mumble's internal
	PacketDataStream serialization format.

	MumbleKit failed to properly validate the length prefix
	of received Opus voice packets.

	If an Opus packet with an invalid length prefix was
	received, MumbleKit would attempt to extract it using
	an MKPacketDataStream object's copyDataBlock: method.

	When the copyDataBlock: method is successful, it returns
	a valid NSData object. When an error occurs, it instead
	returns nil, and sets an error flag in the
	MKPacketDataStream object (which can be queired using the
	'valid' method.)

	Instead of performing proper error checking, MumbleKit
	used the returned NSData object as-is, without first
	checking whether it was nil, or whether the
	MKPacketDataStream object had its error flag set.

	When the returned NSData object is nil, MumbleKit calls
	the Opus functions opus_packet_get_nb_frames() and
	opus_packet_get_samples_per_frame() with a NULL pointer
	as the packet buffer. This causes the functions to
	dereference the NULL pointer.

5. Credits

	This issue was discovered by the Mumble team after a
	reproducible crash that happened when transmitting
	audio was reported by Wesley Wolfe on January 25, 2014.

6. Fix

	A fix for this issue has beeen released in Mumble for iOS 1.2.3.

	Other users of MumbleKit should make sure they are using a
	version of MumbleKit that includes the fixed commit, which is
	fd190328a9b24d37382b269a5674b0c0c7a7e36d.
```
