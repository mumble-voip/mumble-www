---
title: Building web services for Mumble
author: Svedrin
date: 2010-09-27T16:00:19+00:00
categories:
  - Community
  - Tech talk
tags:
  - channel viewer
  - json
  - web
  - xml

---
<figure style="max-width: 208px" class="wp-caption alignleft"><img style="margin-left: 10px; margin-bottom: 10px;" src="http://diesundas.funzt-halt.net/mumble/cvp_article/channelviewer.png" alt="Channel Viewer Screenshot" width="208" height="355" /><figcaption class="wp-caption-text">Mumble-Django channel viewer</figcaption></figure> 

Mumble provides a pretty neat online experience to its users as it is &#8211; still, people often look for some kind of an interface to query the server for status information over the web. The most common use for such an interface is an embedded channel viewer on a team&#8217;s web site, but it could be used for monitoring or other kinds of status notifications as well &#8211; the limit being your imagination.

<!--more-->

If you take a look at the <a title="mumble wiki" href="http://mumble.sourceforge.net/3rd_Party_Applications" target="_blank">3rd party applications</a> Wiki page, you will see that quite a few apps around Mumble have emerged recently, and probably a lot more are currently in development. This is great, as it makes the Mumble ecosystem a lot more vivid. The downside is that if you want to use these applications, you will soon run into problems if each application uses its own protocol to expose information, because using this method will require the server side component and the channel viewer application to match.

![Classic situation][1]

Needing to install a specific server-side component for each individual app would cause a lot of unneccessary work for both installation and maintenance, so being able to use a single server-side component for all apps would be preferable &#8211; especially if your favorite web interface already implemented the server side component.

![CVP with one server][2]

From a client&#8217;s perspective, being able to use your favorite app with a certain server would require the server admin to configure an interface for the client to use. Being able to simply use an existing interface would mean a lot less work, especially since web interface applications usually provide the server side for all servers that they manage.

![CVP with multiple=][3]

The ultimate goal is that each channel viewer is compatible with each server side implementation, no matter which one the respective administrators choose to install, while still being easily maintainable for each server admin.

![Mesh! Clouds! Enterprise! (tm)][4]

However, this can only be achieved when the components use the same data format, because otherwise they simply would not understand each other. To achieve this, the <a title="mumble wiki" href="http://mumble.sourceforge.net/Channel_Viewer_Protocol" target="_blank">Channel Viewer Protocol</a> has been agreed upon. If you want to create a new web application, you are highly encouraged to use the JSON or XML data format defined in the Protocol.

The Channel Viewer Protocol is aimed to be extendable, lightweight and easy to implement on both the server and client side. By using it, you will enable your users to freely choose which applications they prefer to use, and thereby support the idea of software freedom, because giving the user a choice is what this is all about. 🙂

If you have any proposals for the standard, please feel free to visit the [#mumble channel on IRC][5] or post in the [forums][6].

&#8212;

**Note by staff:** Svedrin, the author of this article, is a Mumble community member and developer of the 3rd party Web-Interface [Mumble-Django][7]. We thank him very much for contributing this article. We would love to be able to publish articles written by our community on a regular basis, so if you have an idea you think is worth getting into this blog, please [contact][8] us.

 [1]: http://diesundas.funzt-halt.net/mumble/cvp_article/classic_situation.png
 [2]: http://diesundas.funzt-halt.net/mumble/cvp_article/cvp_one_server.png
 [3]: http://diesundas.funzt-halt.net/mumble/cvp_article/cvp_qry_multiple_servers.png
 [4]: http://diesundas.funzt-halt.net/mumble/cvp_article/cvp_servers_mesh.png
 [5]: http://mumble.sourceforge.net/IRC
 [6]: http://sourceforge.net/apps/phpbb/mumble/
 [7]: http://mumble-django.org/
 [8]: http://mumble.sourceforge.net/FAQ/English#How_can_I_help_or_contact_you.3F