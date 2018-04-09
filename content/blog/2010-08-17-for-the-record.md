---
title: For the record
author: .D0T
date: 2010-08-17T21:20:17+00:00
categories:
  - News
  - Tech talk
tags:
  - feature
  - recording

---
As you already know if you read the article on our [Tracker Squash Meeting][1] or listened to the buzz in `#mumble` on {{< wiki IRC />}} the upcoming {{< wiki "1.2.3" "1.2.3 release" />}} will have a recording feature. The work on this has progressed nicely and since we are now shipping the first snapshots with this new feature in them we think it is time to give you a proper introduction.

<!--more-->

Let&#8217;s start with the basic rules we established for the recorder:

  * Recording is only possible while you are connected to a server and you can only record what you hear or say.
  * For privacy reasons we will not allow recording on pre 1.2.3 servers. (They are not able to properly warn their users)
  * There will be no ACLs to disallow it on newer servers. (We feel that if we restrict the feature that way the Open Source nature of our project will lead to forks with the warnings completely removed which is an even worse situation for the users)

Once you are on a server that fulfils these prerequisites you can open the Recorder via our Icon toolbar or the menu (Self->Recorder) and &#8220;Start&#8221; the recording. You can record for as long as your disk space lasts or until you grow tired of it. The data is encoded and written away on the fly. Supported audio formats at this time are:

  * wav (uncompressed)
  * ogg/vorbis (compressed)
  * au (uncompressed)
  * flac (lossless compressed)

You might ask: &#8220;Where is mp3 support?&#8221;. Well, mp3 has [licensing issues][4] which have prevented the library we use for encoding ([libsndfile][5]) from supporting it. As we do not want any legal troubles either we decided it would be best to keep it that way. We think as we offer multiple lossless formats this should not be much of an issue as you can always encode to mp3 afterwards.

Besides the format you can configure the recording mode. As a normal user you will most likely want to stick with the default: Mixdown. This means all audio streams Mumble receives are mixed together and stored in one audio file. This has the advantage of taking up little space and being easy to handle. But sometimes this is not enough: Say you want to record a Podcast interview or be able to cut out specific comments/noises from a single user in your recording later on. In this case Multichannel recording will be your mode of choice. In this mode Mumble creates a separate audio file for every user you hear. Once you are done recording you can then use audio editors like [Audacity][6], Cubase or Ableton Live to put things back together and to do your editing. To make this as easy as possible Mumble adds silence to the beginning of the files of users who joined the conversation later and in general fills the gaps between talk with the right amount of silence to keep the tracks lined up.

Last but not least let me get to these funny %user, %date etc. things in the target directory and filename. These are so called template variables, if you ever did some (web-)programming you already know what those are good for but let me explain it for the rest. Template variables are special &#8220;words&#8221; that will be replaced with actual content by the program later on. Mumble currently offers the following variables:

  * %user &#8211; Inserts the users name
  * %date &#8211; Inserts the current date
  * %time &#8211; Inserts the current time
  * %host &#8211; Inserts the hostname

This means that if as your filename you choose something like &#8220;%time-%host&#8221; Mumble might expand this to the actual filename &#8220;11-23-21-mycoolserver.net.ogg&#8221; (Note: Mumble appends the right format extension on its own, it is not included in %host or anything). This helps you to customize the way your files are named the way you can handle them best later on. But as I said those variables can also be used in directory names. If you do a lot of recording all those files will clutter up your recording directory pretty fast and even with a great naming scheme you might have trouble to find what you want, this is especially true for multichannel recording with a lot of users. So if this is the way you will use recording you definitely want to add some template variables to your directory names. Mumble will automatically create these directories as needed so &#8220;C:\mygreatpath\to\somewhere\%date\%time\&#8221; will create a directory with the current date, in it a directory for the current time and in that will put your recordings. Do a second recording an hour later? Time has changed so it&#8217;ll end up in another directory. Same goes for the date of course.

This concludes this rundown of our new recording feature. For the audio geeks there is a small appendix with more information on what we actually save in terms of audio data. We hope you will test the hell out of this new feature before we release it with the 1.2.3 stable. Whether you find any bugs or not please give us feedback. The easiest way to do so is to leave comment below.

The Mumble Team

Appendix: So this is how we save it. But the audio geeks among you might ask themselves what we do actually save? Currently the answer to that is: Pretty much the same as you can hear minus positional audio (mono only at this time) plus what you say. Meaning it receives the same post-processing as the audio you hear does and if you have 44.1Khz audio output the file will contain 44.1Khz afterwards (this will change later, internally Mumble always operates at 48Khz and we want to retain that level of quality). The files contain either 24bit PCM (wav & flac) or the full 32bit float range (au & vorbis). Note that due to an internal API restriction our preprocessors currently only operate on 16bit PCM data. This is going to change soon which will make our whole processing pipeline 48Khz/32bit float. As far as codecs go the only lossy codec we offer is ogg/vorbis. For now we stick with the default quality level libsndfile chose for it which is [0.4][7].

 [1]: http://blog.mumble.info/first-mumble-tracker-squash-meeting
 [4]: http://www.mega-nerd.com/libsndfile/FAQ.html#Q020
 [5]: http://www.mega-nerd.com/libsndfile/
 [6]: https://sourceforge.net/projects/audacity/
 [7]: https://en.wikipedia.org/wiki/Vorbis#Technical_details
