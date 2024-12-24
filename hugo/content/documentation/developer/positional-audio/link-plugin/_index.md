---
title: "Integrate Mumble Positional Audio Through the Link Plugin"
---
If you want to add Mumble positional audio support to a game or game mod you have source code access to you can provide positional data natively to Mumble with some *initialization*, *game session update* and *frame position update* logic.

Supporting Mumble positional audio natively will ensure it remain functional across game updates. Native game support makes use of Mumble’s *Link plugin*.

**Note:** If you add native support to your game or game mod please inform us so we can update our documentation accordingly. You can also send in a change request updating it yourself.

If you are looking for just source code you can find that further down below.

## Coordinate system

Mumble, like most sound systems, uses a **left handed coordinate system** <sup>([see this for a visual representation](https://docs.microsoft.com/en-us/windows/win32/direct3d9/coordinate-systems))</sup>.

As a visualization of the left handed coordinate system: If you imagine yourself looking over a large empty field; X increases towards your right, Y increases above your head, and Z increases in front of you. In other words, if we place origin in your chest and you stretch your arms out to your sides, your right hand will be `(1,0,0)`, your left hand will be `(-1,0,0)` and your head will be `(0,0.2,0)`. If you then stretch your arms out in front of you instead, they'll become `(0,0,1)`.

## Positional Data

The positional data is made up of positiona and orientation. The orientation is defined by a front- and upwards-facing unit vector. The position is defined by a vector in meters.

### Position

Mumble interprets the values as meters for consistency between positional audio applications. If you are using a different unit for your actor positions you may have to convert them to meters.

First is the **position vector**. This should be in meters, so you may need to scale it. The reason: If it is not in meters, distance attenuation will be different to the other games/plugins, meaning users will have a bad experience with positional audio (audio configuration is dependant on a common distance measurement).

### Orientation

The next two vectors define the orientation. These should be [unit vectors (length one)](https://en.wikipedia.org/wiki/Unit_vector), and should be perpendicular to one another. The first vector is the front vector, which is simply the direction you are looking at. The second is the top vector, which is an imaginary vector pointing straight out the top of your head. If you do not supply a top vector, Mumble will assume you have a “Y-is-up” coordinate system and that the user can not tilt their head, and then compute the top vector based on that.

**Tip:** If your game uses a 3D positional audio API such as OpenAL you may want to take a look at the code where the listener gets updated since these APIs usually require the same vectors.

### Context

**Note:** The context does **not** need to be updated every single frame. It **shouldn’t** change more than a few times per second if at all during a game.

The context string is used to determine which users on a Mumble server should hear each other positionally. If context between two Mumble users does not match then the positional data is stripped server-side and voice will be received as non-positional.

Accordingly the context should only match for players on the same server in the same game on the same map. Whether to include things like team in this string depends on the game itself. When in doubt err on the side of including less. This gives more flexibility later on.

### Identity

**Note:** The identity does **not** need to be updated every single frame. It **shouldn’t** change more than a few times per second if at all during a game.

Identity should contain a string which uniquely identifies the player in the given context. This is usually satisfied by a player ID (player- or connection-ID on the server, or a global ID).

Additionally the identity can contain any additional information about the player that might be interesting for the Mumble server (for routing, moving or approving users).

For example by including team information in the identity a script on the Mumble server can move players into team channels automatically. Additional information like squad number, squad leader status and so on can be used to trigger even more behavior like automatically maintaining a leadership structure inside Mumble which is kept in-sync with in-game state. E.g. someone is elected squad leader and now can whisper to all other squad leaders and the team leader. For an example of such a server-side script see the [Battlefield 2 mumo plugin](https://github.com/mumble-voip/mumo/blob/master/modules/bf2.py).

We recommend using an easily parseable format like JSON or CSV for encoding the information but this is up to the game. Remember that the link structures only allow for 256 characters of identity data.

## Source Code

All code listed below is put into the **public domain** and can be used, shared or modified freely. You can consider it an example, reference, or code to be integrated directly into your application as-is.

### Initialization

Within your game initialization, add the following initialization code:

```c
#ifdef _WIN32
	#include <windows.h>
#else
	#include <sys/mman.h>
	#include <fcntl.h> /* For O_* constants */
#endif // _WIN32

struct LinkedMem {
#ifdef _WIN32
	UINT32	uiVersion;
	DWORD	uiTick;
#else
	uint32_t uiVersion;
	uint32_t uiTick;
#endif
	float	fAvatarPosition[3];
	float	fAvatarFront[3];
	float	fAvatarTop[3];
	wchar_t	name[256];
	float	fCameraPosition[3];
	float	fCameraFront[3];
	float	fCameraTop[3];
	wchar_t	identity[256];
#ifdef _WIN32
	UINT32	context_len;
#else
	uint32_t context_len;
#endif
	unsigned char context[256];
	wchar_t description[2048];
};

LinkedMem *lm = NULL;

void initMumble() {

#ifdef _WIN32
	HANDLE hMapObject = OpenFileMappingW(FILE_MAP_ALL_ACCESS, FALSE, L"MumbleLink");
	if (hMapObject == NULL)
		return;

	lm = (LinkedMem *) MapViewOfFile(hMapObject, FILE_MAP_ALL_ACCESS, 0, 0, sizeof(LinkedMem));
	if (lm == NULL) {
		CloseHandle(hMapObject);
		hMapObject = NULL;
		return;
	}
#else
	char memname[256];
	snprintf(memname, 256, "/MumbleLink.%d", getuid());

	int shmfd = shm_open(memname, O_RDWR, S_IRUSR | S_IWUSR);

	if (shmfd < 0) {
		return;
	}

	lm = (LinkedMem *)(mmap(NULL, sizeof(struct LinkedMem), PROT_READ | PROT_WRITE, MAP_SHARED, shmfd,0));

	if (lm == (void *)(-1)) {
		lm = NULL;
		return;
	}
#endif
}
```

### Data Update

For each frame, use the following to update position and camera information:

```c
void updateMumble() {
	if (! lm)
		return;

	if(lm->uiVersion != 2) {
		wcsncpy(lm->name, L"TestLink", 256);
		wcsncpy(lm->description, L"TestLink is a test of the Link plugin.", 2048);
		lm->uiVersion = 2;
	}
	lm->uiTick++;

	// Left handed coordinate system.
	// X positive towards "right".
	// Y positive towards "up".
	// Z positive towards "front".
	//
	// 1 unit = 1 meter

	// Unit vector pointing out of the avatar's eyes aka "At"-vector.
	lm->fAvatarFront[0] = 0.0f;
	lm->fAvatarFront[1] = 0.0f;
	lm->fAvatarFront[2] = 1.0f;

	// Unit vector pointing out of the top of the avatar's head aka "Up"-vector (here Top points straight up).
	lm->fAvatarTop[0] = 0.0f;
	lm->fAvatarTop[1] = 1.0f;
	lm->fAvatarTop[2] = 0.0f;

	// Position of the avatar (here standing slightly off the origin)
	lm->fAvatarPosition[0] = 0.001f;
	lm->fAvatarPosition[1] = 0.0f;
	lm->fAvatarPosition[2] = 0.5f;

	// Same as avatar but for the camera.
	lm->fCameraPosition[0] = 0.0f;
	lm->fCameraPosition[1] = 0.0f;
	lm->fCameraPosition[2] = 0.0f;

	lm->fCameraFront[0] = 0.0f;
	lm->fCameraFront[1] = 0.0f;
	lm->fCameraFront[2] = 1.0f;

	lm->fCameraTop[0] = 0.0f;
	lm->fCameraTop[1] = 1.0f;
	lm->fCameraTop[2] = 0.0f;

	// Identifier which uniquely identifies a certain player in a context (e.g. the ingame name).
	wcsncpy(lm->identity, L"Unique ID", 256);
	// Context should be equal for players which should be able to hear each other positional and
	// differ for those who shouldn't (e.g. it could contain the server+port and team)
	memcpy(lm->context, "ContextBlob\x00\x01\x02\x03\x04", 16);
	lm->context_len = 16;
}
```

#### `f`* variables

`fAvatarPosition` should be the player position in 3D space, to disable the transmission of positional information alongside audio without unlinking the plugin set x,y,z of the tuple to zero. `fAvatarTop` and `fAvatarFront` should be the orientation. The coordinate system is a left-handed one, and the units are in meters.

`fCameraPosition`, `fCameraTop`, and `fCameraFront` contain the player camera position and orientation. If the camera of your game is not independent of the avatar be sure to copy the avatar values over to the  variables. If the camera moves independent of the avatar these values should contain the position, top- and front-vector of the *camera*.

**Note:** Mumble fetches these values 50 times a second, so please update them every frame.
