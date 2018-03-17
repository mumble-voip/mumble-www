---
title: Oh no, We’ve made a buzzword
author: slicer
type: post
date: 2010-06-07T23:54:39+00:00
categories:
  - Tech talk
tags:
  - history
  - latency

---
[<img class="alignleft size-full wp-image-44" title="Stopwatch" src="http://blog.mumble.info/wp-uploads/2010/06/Stopwatch-150x150.png" alt="" width="150" height="150" />][1]When we first started with Mumble, we did it because none of the commercial alternatives had quality we were happy with, and there were no open source projects that would have been a good starting base. So, we made our own. This means Mumble was made to satisfy a need; the need we have for high quality voicecom when we play. As such, our two focus areas are voice quality and voice latency. Oh, and &#8220;cool technical stuff&#8221;. But I digress. Quality is subjective, and is hard to quantify, but latency is much easier to test.

<!--more-->

At the time, we were playing a lot of Battlefield 2. And one of our primary problem with the voicecomms we had was that if you saw a grenade a yelled &#8220;GRENADE!&#8221;, what your clanmate actually heard was a loud boom, followed by your desperate attempt to warn him (or her). It really should be the other way around, and we&#8217;ve made sure to emphasize this every chance we got. At first, people were skeptical. Did latency really matter? But as more people moved to use Mumble, more people discovered that quality and latency does matter, as it allows you to speak naturally instead of feeling like a shared walkie talkie.

A few years pass. Other competing products arrive. Old competing products get released in new versions&#8230; And they all mention the magical words &#8220;low latency&#8221;. There are now YouTube videos comparing latency, and numerous technical quasi-explanations around the web on what latency you can achieve with solution X over solution Y. People discuss individual milliseconds. I&#8217;m not sure if I should be proud or ashamed, but we&#8217;ve made &#8220;low-latency&#8221; a new buzzword for gaming VoIP. Proud, because it means we&#8217;ve managed to make the entire field of VoIP apps better for the end user. Ashamed, because to most people it is still just a word without any actual definition to it.

Our best repeatable results, achieved using either ALSA hw: or WASAPI exclusive mode, is around 40-50 ms. That&#8217;s mouth to ear, including network travel time. But what does that actually mean? According to Google, <a href="http://www.google.com/search?q=40ms+*+speed+of+sound" target="_self">40milliseconds equals 13.6 meters</a>. So, place two people 15 meters apart, each with a headset, and have person A say &#8220;Latency&#8221;. Person B will hear it in the headset before he can hear it through the air. That is _low_ latency. Unfortunately, such configurations are useless for actual gaming; exclusive hardware devices means the games can&#8217;t use it.

It&#8217;s still a cool result.

 [1]: http://blog.mumble.info/wp-uploads/2010/06/Stopwatch-150x150.png