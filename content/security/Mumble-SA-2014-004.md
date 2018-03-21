---
title: Mumble Security Advisory 2014-004
---

ID:              Mumble-SA-2014-004 ([txt](../Mumble-SA-2014-003.txt), [sig](../Mumble-SA-2014-003.txt.sig), [Blog post](/blog/mumble-for-ios-1.2.3/))

Date:            February 2014

CVE Reference:   Not assigned

Product:         MumbleKit, Mumble for iOS

Mumble Website:  http://mumble.info

Permalink:       http://mumble.info/security/mumble-sa-2014-004

Last Updated:    05-02-2014

## Vulnerability

A malformed Opus voice packet sent to a MumbleKit client (such as Mumble for iOS) could trigger a heap-based buffer overflow. This causes a client crash (Denial of Service) and can potentially be used to execute arbitrary code, though this is unconfirmed.

This issue can be triggered remotely by an entity participating in a Mumble voice chat.

## Affected versions and configurations

All versions of MumbleKit with Opus support are affected unless they include the fix for this issue, which is available in the MumbleKit Git repository as commit fd190328a9b24d37382b269a5674b0c0c7a7e36d.

Mumble for iOS version 1.1 through 1.2.2 are vulnerable, as they use vulnerable versions of MumbleKit.

## Mitigation

No known mitigation strategy for this issue exists.

Version 1.2.3 and prior of Mumble's server component ('Murmur' or 'mumble-server') does not allow the transmission of Opus packets, and as such a vulnerable client connected to a stock Murmur server that runs version 1.2.3 or prior should not be affected by this
	issue.

Note however that since Mumble is a centralized VoIP system, a modified server could potentially also trigger malformed Opus packets to be sent to clients of its choosing, thus triggering this issue.

## Details

Mumble failed to properly check the return value of a call to the opus_decode_float() function in Mumble's MKAudioOutputSpeech#needSamples: method.

When opus_decode_float() encounters an error, it returns a negative integer signalling the error condition it met.

Instead of catching these errors, Mumble would assign the negative values to a variable denoting the amount of decoded samples (decodedSamples) by the call to the opus_decode_float() function and continue its processing.

Later on in the MKAudioOutputSpeech#needSamples: method, the decodedSamples variable is converted to a pair of unsigned integers: inlen and outlen. The inlen variable's value becomes close to UINT_MAX, since the error codes returned by opus_decode_float() are small negative integers. The outlen variable's value is bounded by a sample rate calculation, which causes the value to be somewhere around UINT_MAX / 48000, depending on the Opus error code and the current sample rate being used by the Mumble client.

Following this, these two unsigned integers are then used as buffer lengths in calls to speex_resampler_process_float() and in a memory-copying "for"-loop at the top of the MKAudioOutputSpeech#needSamples: method.

The inadvertently large buffer lengths cause the two cases above to perform reads and writes outside the bounds of their heap-allocated buffers.

## Credits

This issue was discovered by the Mumble team after a reproducible crash that happened when transmitting audio was reported by Wesley Wolfe on January 25, 2014.

## Fix

A fix for this issue has been released in Mumble for iOS 1.2.3.

Other users of MumbleKit should make sure they are using a version of MumbleKit that includes the fixed commit, which is fd190328a9b24d37382b269a5674b0c0c7a7e36d.
