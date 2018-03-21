---
title: Mumble Security Advisory 2014-002
---

ID:              Mumble-SA-2014-002 ([txt](../Mumble-SA-2014-002.txt), [sig](../Mumble-SA-2014-002.txt.sig), [Blog post](/blog/mumble-1.2.5/))

Date:            February 2014

CVE Reference:   CVE-2014-0045

Product:         Mumble

Mumble Website:  http://mumble.info

Permalink:       http://mumble.info/security/Mumble-SA-2014-002.txt

Last Updated:    05-02-2014

## Vulnerability

A malformed Opus voice packet sent to a Mumble client could trigger a heap-based buffer overflow. This causes a client crash (Denial of Service) and can potentially be used to execute arbitrary code, though this is unconfirmed.

This issue can be triggered remotely by an entity participating in a Mumble voice chat.

## Affected versions and configurations

Mumble 1.2.4 in its default configuration is vulnerable. This is the only stable release that is vulnerable.

Pre-release snapshots released prior to 1.2.4 (these are named 1.2.3-<number>-g<commit>) that include Opus support are potentially vulnerable.

Pre-release snapshots released after 1.2.4 (these are named 1.2.4-<number>-g<commit>) are vulnerable.

Some distributions (such as Debian and Ubuntu) ship a pre-release snapshot in their stable distributions. This snapshot is version 1.2.3-349-g315b5f5, and it is also vulnerable.

## Mitigation

A Mumble client built without Opus support is not vulnerable to this issue.

Opus is enabled in the default build configuration for Mumble, but can be disabled by passing CONFIG+=no-opus to the qmake program when building the Mumble client. (Note that Mumble might still enable Opus support if an installed version of Opus is found via pkg-config. To avoid this, you will need to pass CONFIG+="bundled-opus no-opus" to qmake to also disable the pkg-config querying.)

Version 1.2.3 and prior of Mumble's server component ('Murmur' or 'mumble-server') does not allow the transmission of Opus packets, and as such a vulnerable client connected to a stock Murmur server that runs version 1.2.3 or prior should not be affected by this
	issue.

Note however that since Mumble is a centralized VoIP system, a modified server could potentially also trigger malformed Opus packets to be sent to clients of its choosing, thus triggering this issue.

## Details

Mumble failed to properly check the return value of a call to the opus_decode_float() function in Mumble's AudioOutputSpeech::needSamples() method.

When opus_decode_float() encounters an error, it returns a negative integer signalling the error condition it met.

Instead of catching these errors, Mumble would assign the negative values to a variable denoting the amount of decoded samples (decodedSamples) by the call to the opus_decode_float() function and continue its processing.

Later on in the AudioOutputSpeech::needSamples() method, the decodedSamples variable is converted to a pair of unsigned integers: inlen and outlen. The inlen variable's value becomes close to UINT_MAX, since the error codes returned by opus_decode_float() are small negative integers. The outlen variable's value is bounded by a sample rate calculation, which causes the value to be somewhere around UINT_MAX / 48000, depending on the Opus error code and the current sample rate being used by the Mumble client.

Following this, these two unsigned integers are then used as buffer lengths in calls to speex_resampler_process_float() and in a memory-copying "for"-loop at the top of the AudioOutputSpeech::needSamples() method.

The inadvertently large buffer lengths cause the two cases above to perform reads and writes outside the bounds of their heap-allocated buffers.

## Credits

This issue was discovered by the Mumble team after a reproducible crash that happened when transmitting audio was reported by Wesley Wolfe on January 25, 2014.

## Fix

A fix for this issue has been released in Mumble 1.2.5.

A fix is also available in the master branch of Mumble's Git repository.

A patch which can be applied to previous vulnerable versions can be found in-line below.

(Note: this patch does not apply to Debian and Ubuntu's 1.2.3-349-g315b5f5 version due to whitespace changes of the code above the inserted "if"-statement in the first hunk of the patch.  It is, however, trivially applied or fixed by hand.)

```
--- ./src/mumble/AudioOutputSpeech.cpp
+++ ./src/mumble/AudioOutputSpeech.cpp
@@ -335,6 +335,10 @@ bool AudioOutputSpeech::needSamples(unsigned int snum) {
 					                                   pOut,
 					                                   iAudioBufferSize,
 					                                   0);
+					if (decodedSamples < 0) {
+						decodedSamples = iFrameSize;
+						memset(pOut, 0, iFrameSize * sizeof(float));
+					}
 #endif
 				} else {
 					if (qba.isEmpty()) {
@@ -384,6 +388,10 @@ bool AudioOutputSpeech::needSamples(unsigned int snum) {
 				} else if (umtType == MessageHandler::UDPVoiceOpus) {
 #ifdef USE_OPUS
 					decodedSamples = opus_decode_float(opusState, NULL, 0, pOut, iFrameSize, 0);
+					if (decodedSamples < 0) {
+						decodedSamples = iFrameSize;
+						memset(pOut, 0, iFrameSize * sizeof(float));
+					}
 #endif
 				} else {
 					speex_decode(dsSpeex, NULL, pOut);
```
