---
title: "Guide to Creating a Plugin"
---
Warning: Be aware that the tools/methods used in this tutorial might trigger anti cheat protection. Read the sections below (e.g. this) carefully before attempting to use them on a game.

Game positional audio is a feature of Mumble that many users consider very useful. However, creating the game plugin can sometimes be complicated, and for the average person, daunting. This guide will help you understand how a game plugin works, what it does, and how you can make one.

If you have any questions regarding the process hit us up in [`#mumble-dev:matrix.org`](https://matrix.to/#/#mumble-dev:matrix.org) and we will try to help.

We also accept new plugins into our codebase as long as you are willing to support them. Be aware however that we might reject plugins for games with short update cycles as our current distribution method makes keeping those working for longer times quite painful. When in doubt please verify with us before puttin in the work.

This guide is supposed to be a step-by-step tutorial for the general case. For a more abstract view and more methodology take a look at [the introduction]({{< relref ".." >}}).

## Prerequisites

### Programs needed

[Analysis Tools]({{< relref "../memory-analysis-tools" >}})

#### Tools

C++ IDE and compiler. Read the entire section to know what IDE/Compiler is right for you.

[MumblePAHelper](http://upload.dbclan.de/MumblePAHelper_VS2010.exe). **This binary is built with MSVC 2010, it doesn't work with plugins built with other compilers or MSVC versions.**
This is a useful tool to quickly see if your plugin is working as it should. Simply place it together with `QtCore4.dll` and `QtGui4.dll` (found in the Mumble directory) in your `%APPDATA%\Mumble` folder and run it while your game is running, and it will show if the plugin is linked, if the positional coordinates are working properly, etc. Note that it will only show the plugins which are present inside %APPDATA%\Mumble\plugins.

[Notepad++](https://notepad-plus-plus.org/) (optional). After you install Notepad++, start it, go to Preferences -> New Document/Default Directory, and check "Unix" in the Format box.

### Learn a Little C++

Although you do not need to be an expert programmer in order to write a plugin, you do need to understand fundamental data types. Here are a few of the most important:

float: This is the data type that almost all positional audio game addresses use. They are 32 bit, decimal numbers stored in the memory. A float data type is 4 bytes * 8 = 32 bits. An example of a floating point value would be "1234.0123456".

byte: This is the smallest data type in Intel x86-based computing. This type of memory address holds 1 byte of information (1 byte * 8 = 8 bits). From this type of memory address, you can get 0-255 base^10 values, or -127 to 128, depending on whether or not you use a signed byte (it has a + or - on the front of the value), or an unsigned byte (no + or -). An example of a byte value would be "12".

In C++, you must declare a variable before you can use it. If you want to use a float variable, you declare it with

```cpp
float <variable name>;
```

If you are pointing to an array, you specify how many addresses are in the array. With a float array, it automatically assumes that each address is 4 bytes away from the other. We can declare a 3 address array using

```cpp
float <variable name>[3];
```

declare a byte value using

```cpp
uint8_t <variable name>;
```

remember that there is a difference between amount and location. [3] means three addresses, but locations in the memory start from 0. Therefore, the first address in <variable name>[3] is

```cpp
<variable name>[0]
```

## How a plugin works

Below is a standard template that you can use for your plugin making. The code itself will be explained in the comments that follow.

### Template

```cpp
// Copyright 2019 The Mumble Developers. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file at the root of the
// Mumble source tree or at <https://www.mumble.info/LICENSE>.

#include "../mumble_plugin_main.h" // Include standard plugin header.

static int fetch(float *avatar_pos, float *avatar_front, float *avatar_top, float *camera_pos, float *camera_front, float *camera_top, std::string &context, std::wstring &identity) {
	for (int i=0;i<3;i++) {
		avatar_pos[i] = avatar_front[i] = avatar_top[i] = camera_pos[i] = camera_front[i] = camera_top[i] = 0.0f;
	}
	
	// Boolean values to check if game addresses retrieval is successful and if the player is in-game
	bool ok, state;
	// Create containers to stuff our raw data into, so we can convert it to Mumble's coordinate system
	float avatar_pos_corrector[3], avatar_front_corrector[3], avatar_top_corrector[3];

	// Peekproc and assign game addresses to our containers, so we can retrieve positional data
	ok = peekProc(pModule + 0x<offset>, &state, 1) && // Magical state value: 1 when in-game and 0 when in main menu.
		peekProc(pModule + 0x<offset>, avatar_pos_corrector, 12) && // Avatar Position values (X, Z and Y, respectively).
		peekProc(pModule + 0x<offset>, avatar_front_corrector, 12) && // Avatar Front Vector values (X, Z and Y, respectively).
		peekProc(pModule + 0x<offset>, avatar_top_corrector, 12); // Avatar Top Vector values (X, Z and Y, respectively).

	// This prevents the plugin from linking to the game in case something goes wrong during values retrieval from memory addresses.
	if (! ok) {
		return false;
	}

	if (! state) { // If not in-game
		context.clear(); // Clear context
		identity.clear(); // Clear identity
		// Set vectors values to 0.
		for (int i=0;i<3;i++)
			avatar_pos[i] = avatar_front[i] = avatar_top[i] = camera_pos[i] =  camera_front[i] = camera_top[i] = 0.0f;
		
		return true; // This tells Mumble to ignore all vectors.
	}
	
	/*
	Mumble | Game
	X      | X
	Y      | Z
	Z      | Y
	*/
	avatar_pos[0] = avatar_pos_corrector[0];
	avatar_pos[1] = avatar_pos_corrector[2];
	avatar_pos[2] = avatar_pos_corrector[1];

	avatar_front[0] = avatar_front_corrector[0];
	avatar_front[1] = avatar_front_corrector[2];
	avatar_front[2] = avatar_front_corrector[1];

	avatar_top[0] = avatar_top_corrector[0];
	avatar_top[1] = avatar_top_corrector[2];
	avatar_top[2] = avatar_top_corrector[1];
	
	// Sync camera with avatar
	for (int i=0;i<3;i++) {
		avatar_pos[i]/=32.0f; // Scale to meters
		camera_pos[i] = avatar_pos[i];
		camera_front[i] = avatar_front[i];
		camera_top[i] = avatar_top[i];
	}

	return true;
}

static int trylock(const std::multimap<std::wstring, unsigned long long int> &pids) {
	if (! initialize(pids, L"<executable name>.exe")) { // Retrieve game executable's memory address
		return false;
	}

	// Check if we can get meaningful data from it
	float apos[3], afront[3], atop[3], cpos[3], cfront[3], ctop[3];
	std::wstring sidentity;
	std::string scontext;

	if (fetch(apos, afront, atop, cpos, cfront, ctop, scontext, sidentity)) {
		return true;
	} else {
		generic_unlock();
		return false;
	}
}

static const std::wstring longdesc() {
	return std::wstring(L"Supports <game name> version <game version> without context or identity support."); // Plugin long description
}

static std::wstring description(L"<game name> (v<game version>)"); // Plugin short description
static std::wstring shortname(L"<game name>"); // Plugin short name

static int trylock1() {
	return trylock(std::multimap<std::wstring, unsigned long long int>());
}

static MumblePlugin <game acronym, lowercase>plug = {
	MUMBLE_PLUGIN_MAGIC,
	description,
	shortname,
	NULL,
	NULL,
	trylock1,
	generic_unlock,
	longdesc,
	fetch
};

static MumblePlugin2 <game acronym, lowercase>plug2 = {
	MUMBLE_PLUGIN_MAGIC_2,
	MUMBLE_PLUGIN_VERSION,
	trylock
};

extern "C" MUMBLE_PLUGIN_EXPORT MumblePlugin *getMumblePlugin() {
	return &<game acronym, lowercase>plug;
}

extern "C" MUMBLE_PLUGIN_EXPORT MumblePlugin2 *getMumblePlugin2() {
	return &<game acronym, lowercase>plug2;
}
```

#### Explanation

All that probably looks pretty daunting, right? It isn't really, but even if it does, you actually don't need to understand all of it. You just need to understand the parts that need to be changed in order to make this standard plugin hook to the right game, and fetch the right memory addresses.

First, you need to understand how C++ assignment and functions work. From the code above, let's look at

```cpp
peekProc(pModule + 0x<offset>, avatar_pos_corrector, 12)
peekProc(pModule + 0x<offset>, avatar_front_corrector, 12)
peekProc(pModule + 0x<offset>, avatar_top_corrector, 12)
```

You see that avatar_pos_corrector is a three address array. You would point this to the first address in your positional coordinate array in the memory. If you found a static address in the memory for "something.dll + 24acf2", then you would put this line in the code:

```cpp
procptr_t mod=getModuleAddr(L"something.dll");
```

And set your peekProc to

```cpp
peekProc(mod + 0x024acf2, avatar_pos_corrector, 12)
```

If your game doesn't have static memory addresses from its executable, but only from modules, you could set pModule to the module that has most of the addresses:

```cpp
if (! initialize(pids, L"game.exe", L"client.dll"))
```

But, you're still wondering exactly how the plugin uses one address to get several. Well, right here is where the action happens:

```cpp
ok = peekProc(pModule + 0x<offset>, &state, 1) &&
	peekProc(pModule + 0x<offset>, avatar_pos_corrector, 12) &
	peekProc(pModule + 0x<offset>, avatar_front_corrector, 12) &&
	peekProc(pModule + 0x<offset>, avatar_top_corrector, 12);
```

You see, in this statement, we pass the three pointers that we assigned to the peekProc function, and then assign the results to the avatar_pos_corrector array. Let's look at this one part:

```cpp
peekProc(pModule + 0x<offset>, avatar_pos_corrector, 12)
```

here is what happens: the peekProc function is called and it is instructed to take the first address and add two more. See the "12" at the end? That's the size of the array. Remember that 4 bytes = one float address. 4 * 3 = 12. This means that peekProc sends back three addresses:

```cpp
avatar_pos_corrector[0] = pModule + 0x<offset>
avatar_pos_corrector[1] = pModule + 0x<offset> + 0x4
avatar_pos_corrector[2] = pModule + 0x<offset> + 0x4 + 0x4
```

## Beginning the Hunt

### Introduction

CE = Cheat Engine;

For this game, we will be memory searching the game Alien Arena, an FPS game. If you are trying to create a third-person or camera-based game plugin, you will need to find the avatar positional data and the camera positional data. This is explained further in a few minutes.

Almost all games will have the positional, front, and top coordinates/vectors you need in arrays. This means that the memory addresses will be sequential, one after the other. For instance, 1234ABC0 = X coordinate, 1234ABC4 = Y coordinate, and 1234ABC8 = Z coordinate.

Now it's time for a little Cheat Engine tutorial. Note, again, that Mumble does not support cheating in any way, and this guide does in no way try to teach any cheating methods.

Note that you need to make sure that your server has NO anti-cheat setting enabled, as it might flag Cheat Engine as a hack. NEVER use Cheat Engine on a game that has an anti-cheat method currently engaged, or risk getting banned from that server/game!

If you find a static address, it will either be static from a module, or static from the game executable. If it is static from the game executable, it will be something like

```
something.exe+123ABC
```

when you double click on the address.

### Explanation of Sound and Coordinate Systems

Mumble, like most sound systems, uses a left handed coordinate system. If you imagine yourself looking over a large empty field; X increases towards your right, Y increases above your head, and Z increases in front of you. In other words, if we place origo in your chest and you strech your arms out to your sides, your right hand will be (1,0,0), your left hand will be (-1,0,0) and your head will be (0,0.2,0). If you then stretch your arms out in front of you instead, they'll become (0,0,1).

We need three vectors. First is the position vector. This should be in meters, but if it isn't, you may need to scale it. If it is not in meters, distance attenuation will be different for each game, meaning users will have a bad experience with positional audio.

These three vectors make what is called a **unit vector**. [Here](http://en.wikipedia.org/wiki/Unit_vector) is an excellent explanation of how a unit vector works. It will help you understand how to convert Azimuth and Heading coordinates into a unit vector that Mumble can use, as well.

The next two vectors are the heading. These should be unit vectors, and should be perpendicular. The first vector is the front vector, which is simply the direction you are looking in. The second is the top vector, which is an imaginary vector pointing straight out the top of your head. If you do not supply a top vector, Mumble will assume you have a "Y-is-up" coordinate system and that the user cannot tilt their head, and then compute the top vector based on that.

Once you have the position, you need to find the heading. Since you now know what is the positive direction of the X axis, position yourself so you are looking straight down it. Your 'front' heading should be (1,0,0), so search for a floating point value between 0.7 and 1.05. Turn 180 degrees and search for a value between -0.7 and -1.05. Repeat for Y and Z.

The top vector is done the exact same way, just look down into the ground when finding X; your head now points along the X axis. Note that some games do not have a top vector; a top vector is only "needed" if the game allows you to tilt your head from side to side.

## Hunt

### Part 1 - Find the Position Array

1. First, start your game. If there is a way to make the game windowed, do so. Usually Googling for "<game name> windowed mode" will get the results you need. If you can start your own server for the game, that is preferred. Now load into a map. The game needs to change as little as possible, so make sure you don't have bots or artificial intelligence players enabled.
2. Start Cheat Engine. On the main window, you will see a little computer icon, that is flashing red/green/blue. Click it, find your game executable name on the list, and then double click it. In this case, you would click "crx.exe".
3. You are now hooked to the executable. In the main Cheat Engine window, set "Value type" to "Float" and on "Scan type", select "Unknown initial value". Now click "First Scan". Depending on how fast your computer is, this could take from a few seconds to a few minutes.
4. Move ingame a little. Move forwards, backwards, whatever.
5. Open CE, set "Scan type" to "Changed value", then click "Next Scan".
6. Set "Scan type" to "Unchanged value". Wait a little while (10-20 seconds), then click Click "Next Scan" five or six times.
7. Go back ingame, move a little, and repeat steps 4 to 6.
8. Repeat step 7 a few times.
9. Go back ingame, and look around with your mouse. Do NOT press any WASD keys. Repeat step 6.
10. At this point, you can begin to analyze the addresses that you see on the left. Try to find any addresses that are green. If you can't, it's still ok. This guide found an address with a value of "802.8125", that kept changing when one moved ingame. Now double click the address, and it will be added to the bottom address box.

So, now you should have a position address. Position addresses are almost always an even number in the memory, in hex, offset by four addresses. So, if you had a memory address of 0x5399010, then click "Add address manually", and in the address field, put 0x53990**14**. 0x5399010 + 4 hex = 0x5399014. Now add 0x53990**18**.

The second address should also have a similar looking value, and when you move ingame, it should change accordingly. Depending on how the specific game coordinate system works, the first address might be the X value (east to west), the Y value (up and down), or the Z value (north to south). Jump up and down ingame, and see which value changes the most. That address will be your Y coordinate.

### Part 2 - Find the Front Vector

Now it's time to search for the front and top coordinates. These are a little bit tricky, so you will need some patience. First, see if you can figure out which direction "north" is for the map you are on. Although this may not apply to all games, generally the textures on a map are lined up north-south and east-west, perfectly. This means that if you look straight down a wall, you will be looking perfectly in any one of the four cardinal directions.

1. front straight north, or straight down a texture, whichever works for you.
2. In the main Cheat Engine window, set "Value type" to "Float" and on "Scan type", select "Unknown initial value". Now click "First Scan". Depending on how fast your computer is, this could take from a few seconds to a few minutes.
3. Now look a little bit to the left, just move your mouse enough that you can see a change in the pixels.
4. Open CE, set "Scan type" to "Changed value", then click "Next Scan".
5. Set "Scan type" to "Unchanged value". Wait a little while (10-20 seconds), then click Click "Next Scan" five or six times.
6. Go back ingame, move a little, and repeat steps 4 and 5.
7. Repeat step 6 two or three times.
8. Look straight north or down your wall. Look for a value in the -0.999 or 0.999 range. Move in a circle and see if this decreases or increases, but never gets larger than 0.999, or less than -0.999. By now, the addresses should be narrowed down enough that you should be able to find the right value by just looking through the results list.

### Part 3 - Find the Top Vector

In almost any first person game, the top vector will be within 300 hex of your front vector. So, we will set a scan range for CE. Example: front vector pointer is 1234ABC0, so we will make a scan range of 1234A000 to 1234AFFF. In the memory scan options, set your range to these two addresses.

Go ingame and place yourself looking straight forward.

1. In the main Cheat Engine window, set "Value type" to "Float" and on "Scan type", select "Unknown initial value". Now click "First Scan".
2. Look a little bit down.
3. Open CE, set "Scan type" to "Changed value", then click "Next Scan".
4. Set "Scan type" to "Unchanged value". Wait a little while (10-20 seconds), then click Click "Next Scan" five or six times.
5. Look back up, so your are looking straight at the horizon.

While looking straight toward the horizon, look for a value in the 0.980 to 0.999 range. Now look down at the ground. The value should change to somewhere close to 0 - anywhere between 0.2 and -0.2. Find and add any addresses in that range.

Now add the addresses from the value that you find. Depending on whether or not the coordinate system is left-handed, this could change. In the game that this guide used, the vector component that changed when looking up and down had an address of 05399038, and it was the last address of the vector. So, to get the complete vector, subtract 4 from each address two times, so you would have the following addresses:

```text
05399030
05399034
05399038
```

### Please Note

The Position, Front, and Top vectors are very, very oftentimes right next to each other in memory. If you find the position coordinates, chances are, the rotational ones will be within 100 hex of the former.

Also, note that depending on the game type, all the positional data may be very slightly changing at all times. Therefore, if you cannot find the addresses you are searching for, your game probably falls into this category. If so, you cannot use the "unchanged value" parameter in CE, but must search for relative changes in value.

### Part 4 - Find a State Value

This is probably the easiest part. Simple search for a byte value that remains constant, and changes as soon as you load into a map. Assign it using the sample code in Part 6.
Part 5 - Determine the Coordinate System

So, now you should have the three positional components you need:

1. The position array
2. The front vector array
3. The top vector array

But, you will still need to determine how to arrange the coordinates.

In the left handed coordinate system, the X value will increase, as one vector component remains around 0.999. In Alien Arena, the second address in the position array increases as the second address in the front vector array remains at 0.999. So, we know that this is a centered coordinate system. Therefore, if we spawn at (0,0,0), and we front north, and move one meter, our position coordinate system will be (0,1,0). From that, we determine the following: a left-handed coordinate system uses array of type

```cpp
address[0]
address[1]
address[2]
```

and a center coordinate system uses the same thing, but offset by 2; it uses

```cpp
address[0]
address[2]
address[1]
```

Also, we note that the coordinate system is in Quake units. 1 meter = ~ 32 Quake units, so we will convert it using the following code:

```cpp
for (int i=0;i<3;i++)
	avatar_pos[i]/=32.0f; // Scale to meters
```

now we will take all this information, and piece together our final plugin code.

### Part 6 - Using Pointers

Unfortunately, not all games have static addresses. If yours does not, you will need to perform a pointer scan.

1. After you have found your non-static addresses for each of your respective position, front, and top addresses, right click them and click "Pointer scan for this address"; leave the default settings and start the scan.
2. After the pointer scan is completed, click Pointer scanner -> Rescan memory and type in the actual memory address you're trying to find a pointer to, then click OK.
3. Now it comes down to trial and error. Keep the scan window opened, restart the game and hook Cheat Engine to the process. See if there are pointers that still point to the right address and delete the bad ones by repeating step 2. You can find the updated non-static memory address from a working pointer or by doing a new scan in Cheat Engine. Do this more times to be sure that they are reliable and choose one of them.

After you think you have a reliable pointer, you code it in the fetch() function, above the peekProc part. For example, if I had a pointer of

```text
Level 1 -> Offset CF
Level 2 -> Offset 250
Level 3 -> Offset 230
Level 4 -> something.dll + 242adf
```

for my `avatar_pos`, then I would use the following code:

```cpp
procptr_t avatar_base = peekProcPtr(pModule + 0x0242adf);
if (!avatar_base) return false;
procptr_t avatar_offset_0 = peekProcPtr(avatar_base + 0x230);
if (!avatar_offset_0) return false;
procptr_t avatar_offset_1 = peekProcPtr(avatar_offset_0 + 0x250);
if (!avatar_offset_1) return false;
```

And then I would change peekProc to this:

```cpp
peekProc(avatar_offset_1 + 0xCF, avatar_pos_corrector, 12)
```

## Code the Plugin

The comments are marked in the code. Look it over carefully, and pay close attention to the commenting.

```cpp
// Copyright 2019 The Mumble Developers. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file at the root of the
// Mumble source tree or at <https://www.mumble.info/LICENSE>.

#include "../mumble_plugin_main.h" // Include standard plugin header.

static int fetch(float *avatar_pos, float *avatar_front, float *avatar_top, float *camera_pos, float *camera_front, float *camera_top, std::string &context, std::wstring &identity) {
	for (int i=0;i<3;i++) {
		avatar_pos[i] = avatar_front[i] = avatar_top[i] = camera_pos[i] = camera_front[i] = camera_top[i] = 0.0f;
	}
	
	// Boolean values to check if game addresses retrieval is successful and if the player is in-game
	bool ok, state;
	// Create containers to stuff our raw data into, so we can convert it to Mumble's coordinate system
	float avatar_pos_corrector[3], avatar_front_corrector[3], avatar_top_corrector[3];

	// Peekproc and assign game addresses to our containers, so we can retrieve positional data
	ok = peekProc(pModule + 0x5399000, &state, 1) && // Magical state value: 1 when in-game and 0 when in main menu.
		peekProc(pModule + 0x5399010, avatar_pos_corrector, 12) && // Avatar Position values (X, Z and Y, respectively).
		peekProc(pModule + 0x5399020, avatar_front_corrector, 12) && // Avatar Front Vector values (X, Z and Y, respectively).
		peekProc(pModule + 0x5399030, avatar_top_corrector, 12); // Avatar Top Vector values (X, Z and Y, respectively).

	// This prevents the plugin from linking to the game in case something goes wrong during values retrieval from memory addresses.
	if (! ok) {
		return false;
	}

	if (! state) { // If not in-game
		context.clear(); // Clear context
		identity.clear(); // Clear identity
		// Set vectors values to 0.
		for (int i=0;i<3;i++)
			avatar_pos[i] = avatar_front[i] = avatar_top[i] = camera_pos[i] =  camera_front[i] = camera_top[i] = 0.0f;

		return true; // This tells Mumble to ignore all vectors.
	}
	
	/*
	Mumble | Game
	X      | X
	Y      | Z
	Z      | Y
	*/
	avatar_pos[0] = avatar_pos_corrector[0];
	avatar_pos[1] = avatar_pos_corrector[2];
	avatar_pos[2] = avatar_pos_corrector[1];

	avatar_front[0] = avatar_front_corrector[0];
	avatar_front[1] = avatar_front_corrector[2];
	avatar_front[2] = avatar_front_corrector[1];

	avatar_top[0] = avatar_top_corrector[0];
	avatar_top[1] = avatar_top_corrector[2];
	avatar_top[2] = avatar_top_corrector[1];
	
	// Sync camera with avatar
	for (int i=0;i<3;i++) {
		avatar_pos[i]/=32.0f; // Scale to meters
		camera_pos[i] = avatar_pos[i];
		camera_front[i] = avatar_front[i];
		camera_top[i] = avatar_top[i];
	}

	return true;
}

static int trylock(const std::multimap<std::wstring, unsigned long long int> &pids) {
	if (! initialize(pids, L"Game.exe")) { // Retrieve game executable's memory address
		return false;
	}

	// Check if we can get meaningful data from it
	float apos[3], afront[3], atop[3], cpos[3], cfront[3], ctop[3];
	std::wstring sidentity;
	std::string scontext;

	if (fetch(apos, afront, atop, cpos, cfront, ctop, scontext, sidentity)) {
		return true;
	} else {
		generic_unlock();
		return false;
	}
}

static const std::wstring longdesc() {
	return std::wstring(L"Supports Game version 1.0 without context or identity support."); // Plugin long description
}

static std::wstring description(L"Game (v1.0)"); // Plugin short description
static std::wstring shortname(L"Game"); // Plugin short name

static int trylock1() {
	return trylock(std::multimap<std::wstring, unsigned long long int>());
}

static MumblePlugin gameplug = {
	MUMBLE_PLUGIN_MAGIC,
	description,
	shortname,
	NULL,
	NULL,
	trylock1,
	generic_unlock,
	longdesc,
	fetch
};

static MumblePlugin2 gameplug2 = {
	MUMBLE_PLUGIN_MAGIC_2,
	MUMBLE_PLUGIN_VERSION,
	trylock
};

extern "C" MUMBLE_PLUGIN_EXPORT MumblePlugin *getMumblePlugin() {
	return &gameplug;
}

extern "C" MUMBLE_PLUGIN_EXPORT MumblePlugin2 *getMumblePlugin2() {
	return &gameplug2;
}
```

## Compile the Plugin

You can either build the plugin out of tree, meaning without the rest of the mumble environment around it, or as a part of Mumble. Note that to submit the plugin we prefer a patch against our tree with full integration. However we will do that part for you if you don't want to bother.

### Out of tree build (doesn't need a Mumble build environment)

Make a folder for your plugin and then make a subfolder, with your plugin's name. Now download [`mumble_plugin.h`](https://github.com/mumble-voip/mumble/blob/master/plugins/mumble_plugin.h), [`mumble_plugin_main.h`](https://github.com/mumble-voip/mumble/blob/master/plugins/mumble_plugin_main.h) and the right platform header ([`mumble_plugin_win32.h`](https://github.com/mumble-voip/mumble/blob/master/plugins/mumble_plugin_win32.h) or [`mumble_plugin_linux.h`](https://github.com/mumble-voip/mumble/blob/master/plugins/mumble_plugin_linux.h)). You may also want to download [`mumble_plugin_utils.h`](https://github.com/mumble-voip/mumble/blob/master/plugins/mumble_plugin_utils.h), which contains useful functions you could need in your plugin.

Put the files in the first folder and your .cpp game plugin file into the second folder that is inside of the first.

#### Info on compilers

Mumble 1.3.x is built with MSVC 2015 (Visual Studio 2015 compiler).

Mumble 1.2.x is built with MSVC 2010 (Visual Studio 2010 compiler).

Plugins built with GCC **will not work** with the version of Mumble distributed from the website.

Warning: Since the plugin API exposes internal details of the C++ library (e.g. std::string), plugins must be built with the same compiler/runtime as the Mumble client itself to work.

#### Using Qt Creator

Download the right Qt package version from [Qt.io](https://www.qt.io/download-open-source/#section-2). Qt for MSVC 2010 is discontinued, [this](http://download.qt.io/archive/qt/5.5/5.5.1/qt-opensource-windows-x86-msvc2010-5.5.1.exe) is the latest version that supports it.

You also have to download [Visual C++ 2015 Build Tools](http://landinghub.visualstudio.com/visual-cpp-build-tools) (which contains MSVC 2015, for Mumble 1.3.x) or [MSBuild 4.0](https://www.microsoft.com/en-us/download/details.aspx?id=8279) (MSVC 2010, for Mumble 1.2.x), to install the compiler.

#### Using Visual Studio

[Visual Studio 2015 Community](https://go.microsoft.com/fwlink/?LinkId=615448)  
[Visual Studio 2010 Express edition](http://download.microsoft.com/download/1/E/5/1E5F1C0A-0D5B-426A-A603-1798B951DDAE/VS2010Express1.iso). Default options, except for the SQL server, which you can uncheck.

Start Visual Studio and go to File -> New -> Project. Enter a name, then double click "Win32 Project". Click Next, select "DLL" and check "Empty project". Click Finish. Now open the folder that contains your game's cpp plugin file, and drag that file into the "Source Files" folder on the left. Now click Build -> Batch Build... and check the box in the Build column that corresponds to "Release". Click Build and the plugin will compile. Once it is compiled, go to [My] Documents\Visual Studio 2010\Projects\<project name>\Release. You can take <plugin name>.dll and put it in %AppData%\Mumble\plugins, and when you start Mumble, the plugin will load and you can test it.

### In tree build

If you do not have a working build environment yet follow the [BuildingWindows](https://wiki.mumble.info/wiki/BuildingWindows) or [BuildingLinux](https://wiki.mumble.info/wiki/BuildingLinux) guide to create one. Once you successfully built Mumble (client suffices) perform the following steps:

* Create a new sub-directory for your plugin in `plugins/` named after your plugin (e.g. bf2 for Battlefield 2).
* Name your your primary source file the same as the newly created directory and put it in there.
* Create a `<pluginname>.pro` file in the new directory. You can use an existing plugin's pro file (e.g. bf2/bf2.pro) as a template.
* Add your plugin sub-directory to the list in `plugins/plugins.pro`
* Add your plugin to the installer in `installer/Plugins.wxs` (top and bottom)
* Reconfigure mumble with `-recursive` and rebuild.

The plugin should now be built as part of mumble.
