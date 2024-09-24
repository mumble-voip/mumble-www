---
title: Mumble Certificates
date: 2024-09-23
---

Mumble uses [certificates](https://en.wikipedia.org/wiki/Public_key_infrastructure) to identify users.
A certificate can be self-generated or issued by a thrid party organization. The
certificate types Mumble supports are:

* Self generated
* Class 1 for individuals, intended for EMail
* Class 2 for individuals or organizations, for which proof of identity is required

The basic way of getting a certificate is to let Mumble auto-generate it upon first start.
If the user (and the server the user is trying to connect to) do not care about identity verification,
you can leave everything as is and just connect.

If the user chooses to do that, and enters a valid EMail address, it is possible to
upgrade to a Class 1 or Class 2 certificate for the same EMail address later on, without
the mumble-server forgetting the user. In other words it highly recommended to use a
valid EMail address, even if you start out with a self generated certificate.

A Class 1 certificate simply guarantees that the user has access to the EMail address
in the certificate, since obtaining it involves responding to an EMail sent by the
certificate provider.

A Class 2 certificate validates that the user is who they claim to be by validating
real-life identification.

Class 1 and 2 certificates have to be obtained by 3rd party authorities, some of
which may offer this service for free.

If the user wants to switch certificates they can do so by opening the Mumble
Certificate Wizard found in the configuration menu.
