---
title: Switch to Weblate as Translation Platform
author: Jan Klass
date: 2021-05-10
categories:
  - News
---

The Mumble client is available in many languages. For the translation process we recently moved from the
[Transifex](https://www.transifex.com/) online platform to [Weblate](https://weblate.org/). In this post we describe our
translation history, reasoning, and current approach.

Translations now happen in the
[Mumble client project on Hosted Weblate](https://hosted.weblate.org/projects/mumble/mumble-client/).

<!--more-->

## History - From local to online

Our Mumble client is making use of the Qt framework which provides various features. Amongst them also a system for
localization.

In the source folder `src/mumble` we produce a `.ts` file for each language, which contains the source and corresponding
translated texts. These files could be edited manually with a text or XML editor, or a translation application.

To lower the barrier for potential translators we moved to [Transifex](https://www.transifex.com/), a web based
translation platform in 2014. Using a web platform contributors and anyone else can access translation information with
a web browser, which made it a lot easier to work with. It also adds a few collaborative features that can be used.

To integrate the third party service we implemented a bot
([MumbleTransifexBot](https://github.com/mumble-voip/MumbleTransifexBot)) that would pull translation updates and create
pull requests on GitHub for us. Once established this significantly reduced the work necessary for maintaining this
system and integrating translation changes.

## Reevaluation

In the 6 years since we switched to Transifex, the translation platform changed significantly. Good alternatives
appeared, and we took some discussion from an issue ticket as initiation for a reevaluation of our situation and
opportunities.

Making systematic changes always results in foreseen and unforeseen effort, so not only did we have to evaluate whether
there were better alternatives around, but also whether it was worth it to invest into switching to it. A migration is
not only a change for our integration and hosting process, but also for existing translators that have to potentially
learn new processes and interfaces. We had to evaluate not only technical aspects but also terms of service, and human
aspects.

When looking for alternatives we quickly came to three big contenders:

1. Transifex - alread established and integrated
2. Crowdin - a promising alternative
3. Weblate - official hosted instance of the Weblate open source software

We documented our evaluation and also asked for feedback
[in a discussion topic _(#4641)_](https://github.com/mumble-voip/mumble/discussions/4641).

All three provide free hosting for open source (some limits apply). As a project with a small team and very limited
resources we are very greatful for this. Without these we would be in a much worse position.

Various useful features are shared between the platforms, while the user interface and user experience differs greatly
between them, as well as the technology behind it and resulting way to integrate them.

In the end we decided to initiate a move to [Hosted Weblate](https://hosted.weblate.org/). Apart from the feature
aspects we are also glad to move to a hosted open source software for translations.

## Migration to Weblate

The migration itself has mostly technical aspects to it, and a lot of related work. Deep-diving into the localization
setup and aspects of our project also showed some inconsistencies and questions for the future, and potential for
cleanup which we applied. The migration is/was tracked in
[ticket #4727](https://github.com/mumble-voip/mumble/issues/4727).

Weblate provides integration to GitHub and pull request creation. Some advanced configuration scenarios allow for
additional customization. Thanks to the integration we were able to phase out our integration and PR creation bot.

Interestingly enough even before finishing the migration and even with notices on Hosted Weblate to not translate yet we
received contributions on there.

We have been using Weblate for several weeks now, and have received numerous contributions. Thank you to all existing
and future translators!

The process is not as smooth as we would like it to be yet, but we are glad we made the transition already. And we hope
existing and future translators like the Hosted Weblate platform.

## Weblate Translation Platform

[Weblate](https://weblate.org/) is open source software. The hosted instance services not only us as FOSS but
[other projects](https://hosted.weblate.org/projects/) as well.

We are very thankful for the free hosting they provide to us. If you are looking for a translation platform for other
FOSS or [commercial projects (pricing)](https://weblate.org/en/hosting/), please consider using them.
[Direct donations to them](https://weblate.org/en/donate/) are helping them too, and were explicitly mentioned for
consideration after approving our project.

## Our Translation Project on Hosted Weblate

We support numerous languages, and also have some languages that are not complete enough for productive use yet. If you
are fluent in one of them, please consider contributing to
[our translations](https://hosted.weblate.org/projects/mumble/mumble-client/).

![Weblate translation status by language](https://hosted.weblate.org/widgets/mumble/-/mumble-client/horizontal-auto.svg)

If you want to translate a language that we did not add yet, and are willing to commit to it, please create a request in
[our issue tracker](https://github.com/mumble-voip/mumble/issues). We can add additional languages.

Have any feedback, positive or negative, praise or concerns? We would love to hear from you!

## Road to Mumble version 1.4

We recently landed blockers for 1.4, and are now in the process of preparing a last development snapshot, a feature
freeze, and then plan to publish release candidates.

We would love to release 1.4.0 with many, complete, and good translations. If you are a willing translator, please take
this opportunity to focus on finalizing translations towards the upcoming release. It would be great if we had finalized
translations within about a month or two. Thank you very much.
