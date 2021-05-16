---
title: "Games, Engines, and Game Mods supporting Mumble Positional Audio"
---
**NOTE:** Plugins for 64-bit games only work with Mumble x64!

*Note for Linux users:* Windows plugins work on Wine!

The following tables display which games positional audio is available for, from which game version on (/for which version), for what platform and if extended support is available.

Mumble version 1.2 introduced additional, extended positional audio features. Tooltips on this column’s content may indicate more information on what data is provided.

## Native Support

Games and game engines that provide native support remain functional across game updates.

*Developer reference: [See developer integration documentation]({{< relref "/documentation/developer/positional-audio/link-plugin/" >}})*

### 3D/Game Engines With Native Support

<table>
<thead>
<tr><th>Engine</th><th>Version</th><th>Action required</th><th>Platform</th><th>Extended support (Context, identity)</th></tr>
</thead>
<tbody>
<tr><td><a href="https://ioquake3.org">ioquake3</a></td><td>&gt;= <a href="https://github.com/ioquake/ioq3/commit/0ee3960225">commit 0ee3960225 (2008)</a></td><td>Set <i>cl_useMumble</i> to <i>1</i></td><td>Any</td><td>No</td></tr>
<tr id="Source_Engine"><td><a href="http://source.valvesoftware.com">Source Engine</a></td><td>&gt;= <a href="http://store.steampowered.com/news/9221">?</a></td><td>No</td><td>Any</td><td style="background-color:#abf8a1"><span title="Context based on Server ID. Identity: Universe, Account type, Account steam3ID, Instance, Team">Yes*</span></td></tr>
<tr><td><a href="https://github.com/yEngine/YEng">YEngine Next Generation</a><br>(if using Y.mumble)</td><td>&gt;= <a href="https://github.com/yEngine/YEng/commit/b4011706e7">commit b4011706e7</a></td><td>No</td><td>Linux</td><td style="background-color:#abf8a1">Yes</td></tr>
</tbody>
</table>

### Games with native support

<table>
<thead>
<tr><th>Game</th><th>Version</th><th>Action required</th><th>Platform</th><th>Extended support<br>(Context, identity)</th></tr>
</thead>
<tbody>
<tr><td><a href="http://www.arma2.com">ArmA 2: Operation Arrowhead</a></td><td>&gt;= 1.60</td><td>Install the <a href="https://dev.withsix.com/projects/mumblelink">dedicated mod</a></td><td>Windows</td><td style="background-color:#abf8a1">Yes</td></tr>
<tr><td><a href="https://arma3.com">ArmA 3</a></td><td>&gt;= 0.5</td><td>Install the <a href="https://dev.withsix.com/projects/mumblelink">dedicated mod</a></td><td>Any</td><td style="background-color:#abf8a1">Yes</td></tr>
<tr><td><a href="http://store.steampowered.com/app/362890">Black Mesa</a></td><td>Any</td><td>No</td><td>Any</td><td style="background-color:#abf8a1"><a href="#Source_Engine"><span title="See &quot;Source Engine&quot;">Yes*</span></a></td></tr>
<tr><td><a href="http://store.steampowered.com/app/730">Counter-Strike: Global Offensive</a></td><td>&gt;= <a href="http://blog.counter-strike.net/index.php/2014/08/10222">Aug 27 2014</a></td><td>No</td><td>Any</td><td style="background-color:#abf8a1"><a href="#Source_Engine"><span title="See &quot;Source Engine&quot;">Yes</span></a></td></tr>
<tr><td><a href="http://store.steampowered.com/app/240">Counter-Strike: Source</a></td><td>&gt;= <a href="http://store.steampowered.com/news/9893">Feb 5 2013</a></td><td>No</td><td>Any</td><td style="background-color:#abf8a1"><a href="#Source_Engine"><span title="See &quot;Source Engine&quot;">Yes</span></a></td></tr>
<tr><td><a href="http://www.dayofdefeat.com">Day of Defeat: Source</a></td><td>&gt;= <a href="http://store.steampowered.com/news/9221">1.0.0.46 (01:24:57 Oct 26 2012 (5101))</a></td><td>No</td><td>Any</td><td style="background-color:#abf8a1"><a href="#Source_Engine"><span title="See &quot;Source Engine&quot;">Yes</span></a></td></tr>
<tr><td><a href="http://ezquake.sf.net">ezQuake</a></td><td>&gt;= 2.0 alpha</td><td>No</td><td>Any</td><td>No</td></tr>
<tr><td><a href="https://www.guildwars2.com/">Guild Wars 2</a></td><td>&gt;= <a href="https://forum-en.guildwars2.com/forum/info/news/Game-Update-Notes-February-26-2013">Build 2/26/13</a> <sub><a href="http://wiki.guildwars2.com/wiki/Mumble">Official Wiki page</a></sub></td><td>No</td><td>Any</td><td style="background-color:#abf8a1">Yes</td></tr>
<tr><td><a href="http://store.steampowered.com/app/320/">Half-Life 2: Deathmatch</a></td><td>&gt;= <a href="http://store.steampowered.com/news/9221">1.0.0.37 (01:24:57 Oct 26 2012 (5101))</a></td><td>No</td><td>Any</td><td style="background-color:#abf8a1"><a href="#Source_Engine"><span title="See &quot;Source Engine&quot;">Yes</span></a></td></tr>
<tr><td><a href="http://store.steampowered.com/app/222880/">Insurgency</a></td><td>?</td><td>No</td><td>Any</td><td style="background-color:#abf8a1">Yes</td></tr>
<tr><td><a href="https://minecraft.net">Minecraft</a></td><td>&gt;=  Beta 1.3</td><td>Install the mod for <a href="https://github.com/zsawyer/MumbleLink">Forge</a> or <a href="https://github.com/magneticflux-/fabric-mumblelink-mod">Fabric.</a></td><td>Any</td><td style="background-color:#abf8a1">Yes (<a href="http://www.minecraftforum.net/topic/217587-164-mod-mumblelink-forge-smp-lan-mumble-realism-directional-voip/#developmentAddons">moddable</a>)</td></tr>
<tr><td><a href="http://www.minetest.net/">Minetest</a></td><td>Any</td><td>Install the <a href="https://forum.minetest.net/viewtopic.php?f=53&amp;t=21586">dedicated mod</a></td><td>Any</td><td style="background-color:#abf8a1">Yes</td></tr>
<tr><td><a href="http://store.steampowered.com/app/299740">Miscreated</a></td><td>&gt;= Patch <a href="http://miscreatedgame.com/news/patch-34">#34</a></td><td>No</td><td>Any</td><td style="background-color:#abf8a1">Yes</td></tr>
<tr><td><a href="https://www.gameconnect.net/projects/nucleardawn">Nuclear Dawn</a></td><td>&gt;= <a href="http://store.steampowered.com/news/9647">6.9</a></td><td>No</td><td>Any</td><td style="background-color:#abf8a1">Yes</td></tr>
<tr><td><a href="http://www.openarena.ws">Open Arena</a></td><td>&gt;= <a href="https://github.com/OpenArena/engine/commit/0ee3960225">commit 0ee3960225</a></td><td>Set <i>cl_useMumble</i> to <i>1</i></td><td>Any</td><td style="background-color:#abf8a1">Yes</td></tr>
<tr><td><a href="http://redeclipse.net/">Red Eclipse</a></td><td>Any</td><td>No</td><td>Any</td><td style="background-color:#abf8a1">Yes</td></tr>
<tr><td><a href="https://www.championsofregnum.com">Regnum-Online</a></td><td>?</td><td>No</td><td>Any</td><td>No</td></tr>
<tr><td><a href="http://rq3.com/">Reaction</a></td><td>Any</td><td>Set <i>cl_useMumble</i> to <i>1</i></td><td>Any</td><td style="background-color:#abf8a1">Yes</td></tr>
<tr><td><a href="http://ethernet.wasted.ch">Revenge of the Cats: Ethernet</a></td><td>&gt;= prototype 1.6</td><td>No</td><td>Any</td><td>No</td></tr>
<tr><td><a href="http://www.rigsofrods.com">Rigs of Rods</a></td><td>&gt;= 0.38.20</td><td>No</td><td>Any</td><td style="background-color:#abf8a1">Yes</td></tr>
<tr><td><a href="http://sauerbraten.org/">Sauerbraten</a></td><td>&gt;= <a href="http://sauerbraten.org/docs/history.html#_2008_06_17_ctf_edition">2008_06_17_ctf_edition</a></td><td>No</td><td>Any</td><td>No</td></tr>
<tr><td><a href="http://www.teamfortress.com/">Team Fortress 2</a></td><td>&gt;= <a href="http://store.steampowered.com/news/9221">1.2.3.5 (01:24:57 Oct 26 2012 (5101))</a></td><td>No</td><td>Any</td><td style="background-color:#abf8a1"><a href="#Source_Engine"><span title="See &quot;Source Engine&quot;">Yes</span></a></td></tr>
<tr><td><a href="http://tesseract.gg">Tesseract</a></td><td>&gt;= First Edition</td><td>No</td><td>Any</td><td style="background-color:#abf8a1">Yes</td></tr>
<tr><td><a href="https://ufoai.org">UFO: Alien Invasion</a></td><td>&gt;= <a href="https://ufoai.org/wiki/Changelog/2.5">2.5</a></td><td>Set <i>snd_mumble</i> to <i>1</i></td><td>Any</td><td style="background-color:#ddffdd"><a href="https://github.com/ufoai/ufoai/blob/136e4441b489643d378357fbff5666d60bc916ed/src/client/sound/s_mumble.cpp#L86">Context only</a></td></tr>
<tr><td><a href="https://www.unvanquished.net/">Unvanquished</a></td><td>Any</td><td>Set <i>cl_useMumble</i> to <i>1</i></td><td>Any</td><td style="background-color:#abf8a1">Yes</td></tr>
<tr><td><a href="http://www.warsow.net/">Warsow</a></td><td>&gt;= <a href="http://www.warsow.net/forum/thread/t/191271#post-191271">0.6</a><br>Improved &gt;= <a href="http://www.warsow.net/forum/thread/t/206086#post-206086">1.0</a></td><td>Set <i>cl_mumble</i> to <i>1</i></td><td>Any</td><td>No</td></tr>
</tbody>
</table>

## Support Through Plugin

Outdated plugins are marked in <span style="background-color:#FFFFC0">yellow</span>.

<table>
<thead>
<tr><th>Game</th><th>Version</th><th>Action required</th><th>Platform</th><th>Extended support<br>(Context, identity)</th></tr>
</thead>
<tbody>
<tr><td><a href="http://store.steampowered.com/app/17510">Age of Chivalry</a></td><td>Build 4104</td><td>No</td><td>Windows</td><td style="background-color:#ddffdd"><span title="Based on host IP address and port">Context only</span></td></tr>
<tr class="outdated"><td><a href="http://www.arma2.com">ArmA 2</a></td><td><a href="http://steamcommunity.com/games/arma2/announcements/detail/1027014441254296137">1.08</a> while the latest is <a href="http://steamcommunity.com/games/arma2/announcements/detail/1370438673676192803">1.10</a></td><td>No</td><td>Windows</td><td>No</td></tr>
<tr><td><a href="https://www.battlefield.com">Battlefield 1</a></td><td><ahref="https://www.battlefield.com/news/update-notes/spring-update">1.0.49.52296</a></td><td>No</td><td>Windows</td><td style="background-color:#abf8a1"><span title="Context based on server name. Identity: Team, Squad and Squad leader">Yes*</span></td></tr>
<tr><td><a href="http://www.battlefield.com/battlefield-1942">Battlefield 1942</a></td><td>1.61b</td><td>No</td><td>Windows</td><td>No</td></tr>
<tr><td><a href="http://www.battlefield.com/battlefield-2">Battlefield 2</a></td><td>1.50</td><td>No</td><td>Windows</td><td style="background-color:#abf8a1"><span title="Context based on host IP address and port. Identity: Host IP address and port, Commander, Squad Leader, Squad, Team, On VoIP, On VoIP com, Target squad ID">Yes*</span></td></tr>
<tr><td><a href="http://www.battlefield.com/battlefield-2142">Battlefield 2142</a></td><td><a href="http://largedownloads.ea.com/pub/patches/BF2142/">1.51</a></td><td>No</td><td>Windows</td><td style="background-color:#abf8a1"><span title="Context based on host IP address and port. Identity: Host IP address and port, Team ID, Squad ID, Commander, Squad Leader, Target squad ID, On VoIP, On VoIP com, Player">Yes*</span></td></tr>
<tr><td><a href="http://www.battlefield.com/battlefield3">Battlefield 3</a></td><td>Build 1147186 - End Game DLC</td><td>No</td><td>Windows</td><td style="background-color:#abf8a1"><span title="Context based on host IP address and port. Identity: Team, Squad, Squad Leader">Yes*</span></td></tr>
<tr><td><a href="http://www.battlefield.com/battlefield4">Battlefield 4</a></td><td>1.8.2.48475</td><td>No</td><td>Windows</td><td style="background-color:#abf8a1"><span title="Context based on Server ID. Identity: Host IP address and port, Team, Squad, Squad Leader, Squad State">Yes*</span></td></tr>
<tr><td><a href="http://www.battlefield.com/battlefield-bad-company-2-vietnam">Battlefield Bad Company 2</a></td><td>Build 795745</td><td>No</td><td>Windows</td><td>No</td></tr>
<tr><td><a href="http://www.battlefieldheroes.com">Battlefield Heroes</a></td><td>?</td><td>No</td><td>Windows</td><td>No</td></tr>
<tr class="outdated"><td><a href="http://borderlandsthegame.com">Borderlands</a></td><td><a href="http://borderlands.wikia.com/wiki/Patch_1.4.0">1.4.0</a> while the latest is <a href="http://borderlands.wikia.com/wiki/Patch_1.4.1">1.4.1</a></td><td>No</td><td>Windows</td><td style="background-color:#ddffdd"><span title="Based on host IP address and port">Context only*</span></td></tr>
<tr class="outdated"><td><a href="http://store.steampowered.com/app/49520">Borderlands 2</a></td><td><a href="http://borderlands.wikia.com/wiki/Patch_1.8.X_(Borderlands_2)">1.8.3</a> while the latest is <a href="http://borderlands.wikia.com/wiki/Patch_1.8.X_(Borderlands_2)">1.8.4</a></td><td>No</td><td>Windows</td><td style="background-color:#ddffdd"><span title="Character name">Identity only*</span></td></tr>
<tr><td><a href="http://breachgame.com">Breach</a></td><td>1.1.0</td><td>No</td><td>Windows</td><td>No</td></tr>
<tr><td><a href="http://store.steampowered.com/app/2630">Call of Duty 2</a></td><td>1.3</td><td>No</td><td>Windows</td><td>No</td></tr>
<tr><td><a href="http://store.steampowered.com/app/7940">Call of Duty 4: Modern Warfare</a></td><td>1.7.568</td><td>No</td><td>Windows</td><td style="background-color:#ddffdd"><span title="Based on host IP address and port">Context only*</span></td></tr>
<tr><td><a href="https://steamcommunity.com/app/10190">Call of Duty: Modern Warfare 2 Multiplayer</a></td><td>1.2.208</td><td>No</td><td>Windows</td><td>No</td></tr>
<tr><td><a href="http://store.steampowered.com/app/10180">Call of Duty: Modern Warfare 2 Special Ops</a></td><td>1.1</td><td>No</td><td>Windows</td><td>No</td></tr>
<tr><td><a href="http://store.steampowered.com/app/10090">Call of Duty: World at War</a></td><td>1.7.1263</td><td>No</td><td>Windows</td><td>No</td></tr>
<tr><td><a href="http://store.steampowered.com/app/10">Counter-Strike 1.6</a></td><td>1.6</td><td>No</td><td>Windows</td><td style="background-color:#ddffdd"><span title="Based on host IP address and port">Context only*</span></td></tr>
<tr><td><a href="http://www.dystopia-game.com">Dystopia</a></td><td>Build 4104</td><td>No</td><td>Windows</td><td style="background-color:#ddffdd"><span title="Based on host IP address and port">Context only*</span></td></tr>
<tr><td><a href="http://steamcommunity.com/app/10000">Enemy Territory: Quake Wars</a></td><td>1.50</td><td>No</td><td>Windows</td><td style="background-color:#ddffdd"><span title="Based on host IP address and port">Context only*</span></td></tr>
<tr><td><a href="http://www.finalfantasyxiv.com">Final Fantasy XIV</a></td><td>2016.11.11.0000.0000</td><td>No</td><td>Windows</td><td style="background-color:#abf8a1"><span title="Context based on Map. Identity: Map and Player">Yes*</span></td></tr>
<tr class="outdated"><td>Garry's Mod 11</td><td>Build 5692</td><td>No</td><td>Windows</td><td style="background-color:#ddffdd"><span title="Based on host IP address and port">Context only*</span></td></tr>
<tr><td><a href="http://www.rockstargames.com/IV">Grand Theft Auto IV</a></td><td><a href="https://support.rockstargames.com/hc/en-us/articles/200145406--May-28-2010-Grand-Theft-Auto-IV-Patch-7-Title-Update-v-1-0-7-0-English-1-0-6-1-Russian-1-0-5-2-Japanese">1.0.7.0</a></td><td>No</td><td>Windows</td><td>No</td></tr>
<tr><td><a href="https://www.rockstargames.com/sanandreas">Grand Theft Auto: San Andreas</a></td><td>1.0</td><td>No</td><td>Windows</td><td>No</td></tr>
<tr><td><a href="http://www.rockstargames.com/GTAOnline">Grand Theft Auto V</a></td><td>1.50 (Steam) <a href="https://support.rockstargames.com/hc/en-us/articles/115004482908">1.38</a> (Retail)</td><td>No</td><td>Windows</td><td style="background-color:#ddffdd"><span title="Player nickname, vehicle, location and street">Identity only*</span></td></tr>
<tr><td><a href="https://www.guildwars.com">Guild Wars</a></td><td>Build 36001</td><td>No</td><td>Windows</td><td style="background-color:#ddffdd"><span title="Based on location + area">Partial context only*</span></td></tr>
<tr><td><a href="http://store.steampowered.com/app/17700">Insurgency: Modern Infantry Combat</a></td><td>Build 4044</td><td>No</td><td>Windows</td><td style="background-color:#ddffdd"><span title="Based on host IP address and port">Context only*</span></td></tr>
<tr><td><a href="http://store.steampowered.com/app/8190">Just Cause 2</a></td><td>1.0.0.2</td><td>No</td><td>Windows</td><td>No</td></tr>
<tr><td><a href="http://leagueoflegends.com">League of Legends</a></td><td>1.0.0.145</td><td>No</td><td>Windows</td><td style="background-color:#ddffdd"><span title="Based on host IP address and port">Context only*</span></td></tr>
<tr><td><a href="http://www.l4d.com">Left 4 Dead</a></td><td><a href="http://www.l4d.com/blog/post.php?id=20984">1.0.3.1</a></td><td>No</td><td>Windows</td><td style="background-color:#abf8a1"><span title="Context based on Server ID. Identity: Host IP address and port, map name and player SteamID">Yes*</span></td></tr>
<tr><td><a href="http://www.l4d.com">Left 4 Dead 2</a></td><td><a href="http://www.l4d.com/blog/post.php?id=22240">2.1.4.6</a></td><td>No</td><td>Windows, Linux</td><td style="background-color:#abf8a1"><span title="Context based on Server ID. Identity: Host IP address and port, server name, map name, player nickname and player SteamID">Yes*</span></td></tr>
<tr class="outdated"><td><a href="https://www.lotro.com">Lord of the Rings Online</a></td><td><a href="https://lotro-wiki.com/index.php/Patches">Update 4</a> while the latest is <a href="https://lotro-wiki.com/index.php/Patches">Update 18</a></td><td>No</td><td>Windows</td><td style="background-color:#ddffdd"><span title="Based on region and instance">Partial context only*</span></td></tr>
<tr><td><a href="https://www.quakelive.com">Quake Live</a></td><td><a href="http://steamcommunity.com/games/282440/announcements/detail/876328108049672536">1069</a></td><td>No</td><td>Windows</td><td style="background-color:#abf8a1"><span title="Context based on host IP address and port. Identity: Map and team">Yes*</span></td></tr>
<tr><td><a href="https://www.rocketleaguegame.com/">Rocket League</a></td><td><a href="https://rocketleaguegame.com/news/patch-notes-v1-42">1.42</a></td><td>No</td><td>Windows, Linux</td><td>No</td></tr>
<tr><td><a href="http://playstaxel.com/">Staxel</a></td><td>Any</td><td>No</td><td>Windows</td><td style="background-color:#abf8a1">Yes</td></tr>
<tr><td><a href="http://subrosagame.com">Sub Rosa</a></td><td>0.07b</td><td>No</td><td>Windows</td><td>No</td></tr>
<tr><td><a href="http://store.steampowered.com/app/13240">Unreal Tournament (UT99)</a></td><td>436</td><td>No</td><td>Windows</td><td style="background-color:#ddffdd"><span title="Based on server name">Partial context only*</span></td></tr>
<tr><td><a href="http://store.steampowered.com/app/13230">Unreal Tournament 2004</a></td><td>Build 3369</td><td>No</td><td>Windows</td><td>No</td></tr>
<tr><td><a href="http://store.steampowered.com/app/13210">Unreal Tournament 3</a></td><td>2.1</td><td>No</td><td>Windows</td><td>No</td></tr>
<tr><td><a href="https://www.etlegacy.com/">Wolfenstein: Enemy Territory</a></td><td>2.60b</td><td>No</td><td>Windows</td><td style="background-color:#ddffdd"><span title="Based on host IP address, map and team">Context only*</span></td></tr>
<tr><td><a href="http://blizzard.com/games/wow">World of Warcraft</a></td><td>7.0.3.22810</td><td>No</td><td>Windows</td><td style="background-color:#abf8a1"><span title="Context based on Realm's name. Identity: Player's nickname">Yes*</span></td></tr>
</tbody>
</table>
