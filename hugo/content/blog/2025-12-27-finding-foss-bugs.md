---
title: Bug Hunt - Image Size Boundaries
author: abstractappsec (Sam)
date: 2025-12-27
categories:
  - News
---

*This is a guest post by [abstractappsec](https://github.com/abstractappsec) who found a bug that was fixed in the Mumble 1.5.857 release ([PR #6875](https://github.com/mumble-voip/mumble/pull/6875)).*

This patch was due to an experimental bug hunt that I began a few months ago, and for which I coordinated findings with @Krzmbrzl. 
In this blog post the developers invited me to present, I'll be taking a little detour from your standard Mumble blogs and talking about a way to contribute to open source even if you're not a particularly strong coder. 
This is not an "information disclosure" or remote code execution type of bug, but I'll make the case for its potential security implications as well.

One of the fun things about cybersecurity is that a large part of it is thinking outside of the box.
I work in application security and while I'm not super confident in my coding skills, I have noticed I do have a knack for finding bugs. Not necessarily those crazy bugs you hear about in the news, but I can find some strange ways to make software do things it's not normally expected to do.

This idea and knack is how I ended up digging through the Mumble source code and discovering an interesting bug.

## Picking a Project 

When I decided to start doing some bug hunting, my first goal was to select a target “software under test”. I wanted something that fit into the following categories:

- Open Source software primarily developed with minimal funding/nonprofit/volunteers.
- Software that other bug hunters are probably not very interested in.
- Software where finding and fixing a bug would have an impact.

To the first end, the idea was that I wanted to give back to the community a bit, and FOSS tools are well known to be understaffed and underfunded. Many of them are passion projects, and even when large organizations decide to use them, it’s common to not give back to the communities very much. I’m just a little cog in all of these big wheels, but it felt like a good way to do my part as an application security person and be able to start contributing to FOSS.

What about “software other bug hunters aren’t very interested in” ? Many bug hunters primarily target projects that have bug bounty programs. These programs mean that discovery of a meaningful security bug comes with a payout. These programs, in my opinion, are great and can incentivize independent researchers, but similar to my first point, FOSS projects generally don’t have great funding to do this kind of thing. So many bug hunters might just ignore smaller FOSS projects because of the lack of economic incentive. 

And to my third point, while somebody might find a bug in a local offline calculator application and it could be cool, it’s unlikely that the result would have a significant effect on the users of the application. In contrast, when software is used in a distributed, networked manner, it can be a pretty big deal for someone to discover a bug that one user can use to affect others. 

Mumble seemed to fit all of these categories well. It’s a great program I remember using fondly for playing Minecraft with my brother back in the day, but has a relatively small community and footprint, so it's unlikely there will be much in the way of professional application security people combing through the code. 

What makes it actually “important” to find and patch bugs still? 
When I initially started looking at Mumble, I was aware of a few online gaming communities that still use Mumble pretty heavily. Someone wanting to nefariously cheat in a game without exploiting the game itself could theoretically attack a chat application and cause chaos on an opposing team in a competitive context, gaining an unfair advantage. Additionally, for all companies using Mumble for communications, it could be an avenue of attack against them if someone found a remote code execution or Denial of Service vulnerability.  

So now that we’ve identified our target, what comes next?

## Triage

Initial Triage is the first step once you have decided on a Software under test. This is where you identify user input that could potentially cause issues.
There are vulnerabilities outside of “user input” but as a general rule these are the most straightforward and impactful. More generally, any input from a foreign device can be a cause for a remote code execution or DOS vulnerability. With that in mind, it was easy to start by poking at the Mumble Chat messages. These can be sent by any user on the server to other users on the server. There are high profile vulnerabilities in chat applications all the time, because parsing is very difficult when you want to support rich features such as images, formatted text, and link previews.

One of the things I found that I quite liked is that while Mumble supports HTML embedded in messages, it can be turned off server wide. This is a great feature in my opinion, because of a principle that has been expounded upon by the langsec community (https://langsec.org), which is that as you increase the complexity of “accepted inputs”, you decrease your ability to identify if there are any particular inputs that create error conditions. Being able to send HTML messages rather than simple plaintext greatly increases this complexity. 
If this seems like an interesting topic, I recommend visiting the link above and learning about Language-theoretic security, but to keep it simple here, my general rule of thumb is that if a computer program accepts complex input, it’s a good place to look for bugs!

Often when you find a bug, it’s higher impact when the bug is present with default settings, rather than only being a problem if administrators actively put a server into a dangerous configuration. In the case of the bug I found for instance, even if Mumble hadn’t been able to push out a patch, a security advisory suggesting that users disable HTML chats would have completely mitigated the problem, which is good software design.  Being able to make a decision as the system administrator to either accept the higher risk of supporting a more complex protocol to improve user experience, or restrict options to mitigate potential bugs is a great way to mitigate these kinds of risks.


As a general rule, when you are auditing code, the scary stuff, like complex nested functions, is the important stuff to pay more attention to. Additionally, it is important to pay attention to anywhere that the code is checking for an error condition. This may seem counter intuitive, but the reality is that when we program, we often solve for some parts of errors, and then forget to go back and solve them more generally. For example, a developer might write a function without error checking at all just to make sure it works in an expected scenario, then go back through the code and check for a simple error case like a null pointer, but forget to go back a third time and check for scenarios where a user is maliciously trying to exploit the function.


## Discovery

Mumble chat messages are handled primarily by functions in a file named Log.cpp. I haven’t dug into the history, but I’d guess originally this functionality was for logging primarily, but worked well enough that it was repurposed for user messages as well.

Among other things, this code handles functions like displaying inline images to the user, which can be displayed in html like this(as an example):

```
<br>
<img src=<source of image> height=200 width=200>
</br>
```


Notice the height and width values. Normally these are used when a website or html parser wants to shrink or grow an image. Maybe for a thumbnail or for a window being resized.
Nothing stops an HTML parser from having huge numbers in these values by default- typically most things that want to display an image will not want it to be huge, but since there are exceptions, it’s usually on the application writers to identify a number that is sane for their applications.

In Mumble, there are a couple of checks in this order, in the validHTML() function of Log.cpp:


```cpp
// we start with a LogDocument object named qtd, and set the Html of that object to the html being passed in. 
qtd.setHtml(html)
// skipping lines here and there.

qtd.adjustSize()
QSizeF s = qtd.size()
s.isValid() 

    int messageSize = static_cast< int >(s.width() * s.height());
    int allowedSize = 2048 * 2048;

    if (messageSize > allowedSize) { // error }
```

Notice this last check - we’re validating that the size is not greater than our allowedSize of 2048 * 2048 = 4,194,304. 

However, the problem is that `s` is a `QSizeF` object, which means that width and height are doubles in Mumble's code. Doubles are twice the bitness of an int. If you’re not super comfortable with integer overflow, what this means is that in CPP, “normally” an int is a datatype that can store 4 bytes of data. 1 bit is used to signify whether the number is positive or negative, so that leaves us with 2^31 as our maximum value that can be stored in an int: 2,147,483,647. 

A double is, well, a double integer. That means you have twice as many bits to use, which allows you to store a way bigger number!
When you cast from a double to an int, in CPP, it just truncates (chops off) the leftmost 4 bytes. That means you can multiply two numbers together and get one that uses a few of the bytes in the right 4, and a few in the left 4, and suddenly when you cast it to an int, it will leave you only with the bytes in the right four.

Here’s an example using python:

```python
>>> int('0xff_ff_ff_ff_00_00_00_01', 16)
18446744069414584321
>>> int('00_00_00_01', 16)
1

```
This number on the top is using 8 bytes, which in cpp normally could be stored in a double (technically this is a generalization if you aren’t using fixed width integers, but it is how it normally compiles in most scenarios).
The second number, only using 4 bytes, is something you could store in an int. If you took the first number and casted it to an int in cpp, it would just make the left 4 bytes magically disappear. 
(We can emulate this behavior in python for quick testing using ctypes.)

```python
>>> import ctypes    
>>> doubleval = int('0xff_ff_ff_ff_00_00_00_01', 16)
>>> print(doubleval)
18446744069414584321
>>> print(ctypes.c_int(doubleval).value)
1
```
That means in this scenario, the double value is cast to an int, then the int is checked against the 2048 * 2048 value. A little bit of math later proves that multiplying width and height to a number that wraps to a bit past INT32MAX is going to allow the message to pass this check, by making the large part of the number all be stored in the leftmost 4 bytes.

```
16777218 * 128 = 2,147,483,904 
```
How does this particular number look when stored in a double?


```python
>>> import ctypes                                                                                                                                     
>>> double_val = ctypes.c_double(2_147_483_904)  
>>> double_val.value
2147483904.0 
```

Looks fine!

But what if we cast it to an int?

```python
>>> print(ctypes.c_int(int(double_val.value)).value)
-2147483392
```
Woa! That’s a way different number!. This is because when we truncate the number down to a 4 byte type, we lose all sorts of things. There are good and bad reasons why C and CPP behave this way - sometimes you actually do want to truncate numbers - but in general, this is a vulnerability pattern known as an “integer overflow” or a “type confusion.” (Though type confusions are often much more complex than issues like this one.)

What was that check again?

```cpp
    if (messageSize > allowedSize) { // error }
```
So we’ve now generated a message that has a very very large value in its original form as a QT object, but is only checked after being type confused - and the key part is that the original object is still the original size, the check does not actually mutate the object. 
As a result we can do something like this:

```html
<br>
<img src=data:,Hello%2C%@0World%21 height=16777218 width = 128>
</br>
```
(src doesn’t really matter here, it just gives the parser something to try to use as an image).

QT will now happily attempt to display this entire image in its huge dimensions to the entire server! 

## Impact
As a security researcher, impact is one of the most important things to identify. The recent React vulnerability is a great example of a bug that is low sophistication, high impact. There are tools that allow you to measure various elements of what makes an attack possible, and how difficult it is to exploit, and gives you a general idea of how big of a deal it is on a scale of 1-10 usually.
https://nvd.nist.gov/vuln-metrics/cvss/v3-calculator
This is the NIST tool, and I like it a good bit, but I will warn that it’s a bit more complex than the number that gets spit out. Not every 10 is as big of a deal as people make it out to be, and not every 3 is a nothingburger. It’s very important to understand the environments that the code is likely to run in, and how big of a problem it would be for it to break. For instance, if a hospital is using an application, something that might just be a nuisance somewhere else, like a very large image in the messaging platform, could be the difference between life and death in an emergency.

I was fairly confident that mumble wasn’t being used in the medical field, at least not extensively, so I figured no matter how big of a deal the bug was, it wasn’t going to be the top of importance for most organizations. I stand by that with the information that I have now, though I have since learned that mumble is used or has been used in some surprisingly large organizations, so the impact of a bug like this could be somewhat higher than I originally expected. 

Ok but what actually happens when you send a malformed message?

Well, the entire chat gets “nuked”. Sending a message like my proof of concept completely removes the ability for anyone to see any older messages by my testing. Additionally, it makes it so new messages can’t be seen either. Disconnecting and reconnecting does not fix the issue. The only way to get new messages to show up is to completely close your mumble client and reconnect to the server. However, you won’t be able to see any of the old messages that way. As a result, your entire message history is just lost. 

This can actually get worse on Windows. Users can have a “comment” which makes a little yellow chat message appear next to their name.

If you set the comment to the proof of concept, anyone on a Windows client who clicks on the message to read it will have their client crash!

Crashing is often what makes bugs especially problematic. While not always the case, this can be a sign of memory corruption, which is one of the most common historical avenues to go from a crash to an attacker getting code execution on a remote machine.


## Impact

Here is the calculation I performed, when identifying the bug based on the impact that I estimated.
This attack can be performed over a network, and is easy to execute. Privileges required are just to be a user on the server, and for the normal message sending, there is no user interaction required. (Though the crashing bug does require user interaction.).

It doesn’t do any information leaking, so it’s not really affecting confidentiality. But integrity is compromised because the server chat history is wiped. It’s not the worst integrity violation, which would be something like a bug that allows you to manipulate other user’s chat messages and things like that. 

The main impact is on availability: as long as a malicious user can keep reconnecting to the server, they can keep spamming this message and rendering chat completely useless. Worse, because the entire chat history gets messed up like this, it can be very difficult for an administrator to figure out who the actual problematic user is.

Plugging this into the NIST tool gives us a 7.1 CVSS by my estimation. Fairly high in the abstract, but the actual impact always depends on how people actually use the software. For instance, I have a friend who plays EVE online, and heavily uses Mumble, but only the voice chat. In his usage, this bug would be very unlikely to affect him. Someone would need to trick him over voice chat into clicking on a user comment to get his client to crash - not something likely to be doable at scale or more than once!

## Proof of Concept

The last step was to ask the question “is this enough?”. Proofs of Concept are very important in security research because they demonstrate “just how bad” a bug can be. If I had been able to generate a PoC that executed remote code, that would have certainly been a big deal, so why didn’t I search further along the crash to see if I could do that?

Well, there’s a few good reasons. One is that I did do some brief inspection of the crash. I’m not very good at debugging, so it wasn’t immediately obvious if the crash was caused by something that could be exploited. I also noticed that Mumble uses a lot of good modern binary protections, such as stack canaries, which make exploiting memory corruption a lot harder.

Finally, it was a question. The only real point of making a more complex PoC is to convince the developers that a problem is significant enough that it warrants a patch, and associate a relevant timeliness to the patch. Pretty much as soon as I reported what I had, the devs began working with me to get a patch pushed out. Additionally, since I am not tracking that Mumble is used in any of those contexts where it would be a huge deal for someone to break it like this (like life or death) I figured what I had was enough to deliver. 

It is easy for those of us who work in application security or bug hunt as a hobby to feel like we need to chain every bug into its worst case scenario situation. That’s what gets the headlines and what often makes people very interested in your work. If you’re chasing notoriety or a big payout for a bounty, developing more crazy exploits can make sense, but if you just want to help projects be a bit more secure, it doesn’t necessarily matter. At the end of the day, the basic goal of a PoC for responsible disclosure is to convince the developers to patch something. If they’re convinced, you’ve succeeded. 


## Interfacing

One of the other reasons I picked Mumble as a good software for a bug hunt was the clear way you can submit bug reports. https://www.mumble.info/security/#reporting-a-vulnerability. 
This is a heavily underrated feature in software, though I’ve been pleasantly surprised to see many projects adopting various forms of clear security disclosures. These make it easier to incentivize independent security researchers and developers to help fix things.

I had the pleasure of working with @Krzmbrzl who responded to my initial email only 4 days after I sent, which is a great turnaround time for a volunteer run project. (And from what I’ve heard from some friends who do similar things, it’s a great turnaround for enterprise organizations as well!).

Throughout the disclosure process, we worked together (ok, mostly krzmbrzl did the work from here on out ;D ) on ensuring that the bug was real (not just on my machine), and validating that his patch did indeed fix the bug. Even though he wasn’t able to fully replicate all of my findings, with what I sent him he confirmed it was significant enough to fix. Additionally, while we had some differences on the semantics of whether it was a “security” issue per se, ultimately the bug was patched and any potential concerns, real or imaginary, are mitigated. I think sometimes it’s easy when you live in the security world to see everything as a dumpster fire, but ultimately the goal is to improve user experience. If you can do that by increasing the validity of confidentiality, integrity, or availability, I think the goal is accomplished and I like to think that we did that in a small way with this patch.  

## Conclusion

I hope that this little detour into the world of bug hunting might convince some folks to help FOSS projects by doing bug hunting as well. Some bugs can be surprisingly easy to find, others surprisingly important. However, even if all you find are ‘3’s and below, you can still be helping the community and improving software for users around the world.

