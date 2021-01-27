---
title: Mumble Security Advisory 2014-001
date: 2016-01-10
---

ID:              Mumble-SA-2014-001 ([txt](../Mumble-SA-2014-001.txt), [sig](../Mumble-SA-2014-001.txt.sig), [Blog post](/news/mumble-1.2.5/))

Date:            February 2014

CVE Reference:   CVE-2014-0044

Product:         Mumble

Mumble Website:  http://mumble.info

Permalink:       http://mumble.info/security/Mumble-SA-2014-001.txt

Last Updated:    05-02-2014

## Vulnerability

A malformed Opus voice packet sent to a Mumble client could trigger a NULL pointer dereference or an out-of-bounds array access, leading to a crash (Denial of Service).

This can be triggered remotely by an entity participating in a Mumble voice chat.

## Affected versions and configurations

Mumble 1.2.4 in its default configuration is vulnerable. This is the only stable release that is vulnerable.

Pre-release snapshots released prior to 1.2.4 (these are named 1.2.3-<number>-g<commit>) that include Opus support are potentially vulnerable.

Pre-release snapshots released after 1.2.4 (these are named 1.2.4-<number>-g<commit>) are vulnerable.

Some distributions (such as Debian and Ubuntu) ship a pre-release snapshot in their stable distributions. This snapshot is version 1.2.3-349-g315b5f5, and it is also vulnerable.

## Mitigation

A Mumble client built without Opus support is not vulnerable to this issue.

Opus is enabled in the default build configuration for Mumble, but can be disabled by passing CONFIG+=no-opus to the qmake program when building the Mumble client. (Note that Mumble might still enable Opus support if an installed version of Opus is found via pkg-config. To avoid this, you will need to pass CONFIG+="bundled-opus no-opus" to qmake to also disable the pkg-config querying.)

Version 1.2.3 and prior of Mumble's server component ('Murmur' or 'mumble-server') does not allow the transmission of Opus packets, and as such a vulnerable client connected to a stock Murmur server that runs version 1.2.3 or prior should not be affected by this issue.

Note however that since Mumble is a centralized VoIP system, a modified server could potentially also trigger malformed Opus packets to be sent to clients of its choosing, thus triggering this issue.

## Details

Mumble's Opus voice packets are serialized as a buffer with a length-prefix using Mumble's internal PacketDataStream serialization format.

Mumble failed to properly validate the length prefix of received Opus voice packets.

If an Opus packet with a length prefix of zero was received, the Mumble client would attempt to extract an Opus buffer of size 0 by calling the dataBlock() method on a PacketDataStream object. In this case the dataBlock() method would return a QByteArray that uses the result of malloc(0) as its internal buffer.

Depending on the system's implementation of malloc this call may return either NULL, or a non-NULL pointer that points to a zero-length buffer.

The QByteArray's internal buffer is later used in a call to the opus_packet_get_samples_per_frame() function, which attempts to read the first byte of the passed-in buffer to calculate its return value. This can either cause a NULL pointer dereference, or a read outside the bounds of the zero-sized heap-allocated buffer.

Similarly, if the Mumble client received an Opus packet with a length prefix that is negative, or larger than the encapsulating packet, the dataBlock() method of the PacketDataStream object would return a QByteArray constructed by the QByteArray class's default constructor. That is, a 'null' QByteArray where the internal buffer is a NULL pointer.

This NULL buffer is then passed to the opus_packet_get_samples_per_frame() function which will dereference it when attempting to read the first byte of the buffer.

## Credits

This issue was discovered by the Mumble team after a reproducible crash that happened when transmitting audio was reported by Wesley Wolfe on January 25, 2014.

## Fix

A fix for this issue has been released in Mumble 1.2.5.

A fix is also available in the master branch of Mumble's Git repository.

A patch which can be applied to previous vulnerable versions can be found in-line below.

```
--- ./src/mumble/AudioOutputSpeech.cpp
+++ ./src/mumble/AudioOutputSpeech.cpp
@@ -148,8 +148,15 @@ void AudioOutputSpeech::addFrameToBuffer(const QByteArray &qbaPacket, unsigned i
 		int size;
 		pds >> size;
 		size &= 0x1fff;
+		if (size == 0) {
+			return;
+		}

 		const QByteArray &qba = pds.dataBlock(size);
+		if (size != qba.size() || !pds.isValid()) {
+			return;
+		}
+
 		const unsigned char *packet = reinterpret_cast<const unsigned char*>(qba.constData());

 #ifdef USE_OPUS
```
