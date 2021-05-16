---
title: Developer Translation Documentation
date: 2021-03-14
lastmod: 2021-03-14
aliases: 
  - "/documentation/translation/"
---
Mumble makes use of the Qt framework. This is also the case for translations. We follow the Qt tools and translation workflow.

Translators can translate our project in our [Mumble Weblate project](https://hosted.weblate.org/projects/mumble/mumble-client/) without a need for Git, Code, or programs other than their web browser.

The `.ts` translation files are located in the Mumble repository in `src/mumble/mumble_*.ts`.

We currently version a `mumble_en.ts` as a reference file for translations, and individual files for each locale we support (which each holds the source strings too anyway).

Qt supports both an [ID-based approach](https://doc.qt.io/qt-5/linguist-id-based-i18n.html) as well as a source + translated text approach. Weblate calls the ID based approach monolingual (a ts file does not contain the source text) and bilingual (a ts files contains both source and translated text).

We effectively work with the bilingual file approach, but have the Weblate project set up as if it were ID based. _This is for historic reasons and should be changed in the future._

## Adding a language

To add a new language the ts file is created with `lupdate`, and a file reference is added to the `ts_files` variable in `src/mumble/CMakeLists.txt`.

## Creating translatable strings

In Qt translation strings have a context – a class context. The fallback is the QApplication context. But any `Q_OBJECT` class introduces their own context. In the `ts` file texts are grouped by context.

Literal text is translated (“marked”), typically through the [`QObject::tr`](https://doc.qt.io/qt-5/qobject.html#tr) method, which in turn calls [QCoreApplication::translate](https://doc.qt.io/qt-5/qcoreapplication.html#translate) with an adequate.

NOTE: To provide structure to the many translation strings, `tr` should be called on the most reasonably specific class. *(Calling it on the base `QObject` will make it fall back to the generic `QApplication` context.)*

Classes not extending from QObject [can also be extended with the `tr` context functionality](https://doc.qt.io/qt-5/i18n-source-translation.html#translating-non-qt-classes).

See also upstream documentation [Using tr() for All Literal Text](https://doc.qt.io/qt-5/i18n-source-translation.html#using-tr-for-all-literal-text) (and the context around it).

### Numbers and Dates

Special care has to be taken for formatting date and time and numbers. The format of these differs by locale, and may be disassociated from the language a user uses too.

See also QLocale and other Qt classes for date, time and numbers.

* [Using tr() to Localize Numbers](https://doc.qt.io/qt-5/i18n-source-translation.html#using-tr-to-localize-numbers) (shows only one case/may not be the best approach)

### Disambiguation

Identical text in the same context can be disambiguated with an additional identifying string.

```c++
    QLabel *senderLabel = new QLabel(tr("Name:", "sender"));
    QLabel *recipientLabel = new QLabel(tr("Name:", "recipient"));
```

### Plurals

Plurals have special handling in general and between languages. Depending on the locale/language, a plural (word) may differ between 0, 1, 2. And they may differ in different ways.

Numbered plurals are marked with a `(s)` word postfix like `%n word(s)`.

Notably this will not show up as `word(s)` in the program, but as `word` or `words` depending on the parameter value.

The translation files (ts files) may have any number of `<numerusform>` elements as translations, rather than just one translation text.

```xml
    <message numerus="yes">
        <source>Ban List - %n Ban(s)</source>
        <translation>
            <numerusform>Bannliste - %n Bann</numerusform>
            <numerusform>Bannliste - %n Bann(s)</numerusform>
        </translation>
    </message>
```

Other languages may have three forms.

[Handling Plurals](https://doc.qt.io/qt-5/i18n-source-translation.html#handling-plurals)
[Translation Rules for Plurals](https://doc.qt.io/qt-5/i18n-plural-rules.html)

### Numbered Arguments (Ordered)

The order of arguments within a string may change with a text translation. Numbering them and passing values with `arg()` will allow translators to change the order as they see fit.

```c++
    label.setText(tr("%1 of %2 files copied.\nCopying: %3")
                  .arg(done)
                  .arg(total)
                  .arg(currentFile));
```

A translator comment may help them identify parameters if not obvious from the source text.

### Mnemonics

Programs typically show these while holding down ALT, and single characters can be used to activate controls allowing for keyboard-based usage of GUI applications. While ALT is pressed the mnemonic key is usually underlined.

The typical ampersand `&` markers are used to define mnemonics on a control text.

For example:

* `E&xit` would lead to `x` being the mnemonic
* `&Exit` would lead to `e` being the mnemonic

In both cases, the displayed text on the control will be `Exit`.

### Accelerator Values

Use [`QKeySequence`](https://doc.qt.io/qt-5/qkeysequence.html) to use keyboard shortcuts so they are translatable too.

```c++
exitAct = new QAction(tr("E&xit"), this);
exitAct->setShortcuts(QKeySequence::Quit);
```

* [Using QKeySequence() for Accelerator Values](https://doc.qt.io/qt-5/i18n-source-translation.html#using-qkeysequence-for-accelerator-values)

### Translating text outside of classes

Classes provide code structure. Yet sometimes text may end up outside of classes and Qt classes, and reasonably so.

Then `tr` can be called as a public static method on an appropriate class the text belongs to. Or `QCoreApplication::translate` can be called directly.

For text without context the `QT_TR_NOOP` and `QT_TRANSLATE_NOOP` macros can be used.

### Translator Comments

Translating text often requires context. It can be difficult to grasp a texts context from just the source string. Without the program or source code open and next to the source string, it can be impossible to grasp the context.

Translator comments can help provide context to a source string.

```c++
//: This name refers to a host name.
hostNameLabel->setText(tr("Name:"));

/*: This text refers to a C++ code example. */
QString example = tr("Example");
```

* [Translator Comments](https://doc.qt.io/qt-5/i18n-source-translation.html#translator-comments)

### Metadata

Additional metadata can be attached to translatable text for advanced behavior, also further down the line on translation platforms.

* [Qt: Adding Meta-Data to Strings](https://doc.qt.io/qt-5/i18n-source-translation.html#adding-meta-data-to-strings)
* [Weblate: Customizing behavior using flags](https://docs.weblate.org/en/latest/admin/checks.html#custom-checks)

## CMake integration

We make use of [`qt5_create_translation`](https://doc.qt.io/qt-5/qtlinguist-cmake-qt5-create-translation.html). This will call `lupdate` and `lrelease`.

The ts translation files are stripped of unrelated, unused texts (they contain text for multiple platforms while the program is compiled for just one), and then packaged into the application (through the Qt resource system; qrc definition file).

Regarding Qt resources see [The Qt Resource System](https://doc.qt.io/qt-5/resources.html).

## Upstream Qt Resources

* [Qt Linguist Manual: Developers](https://doc.qt.io/qt-5/linguist-programmers.html)
* [Qt Linguist Manual](https://doc.qt.io/qt-5/qtlinguist-index.html)
* [Writing Source Code for Translation](https://doc.qt.io/qt-5/i18n-source-translation.html)
* [Qt Linguist Manual: Text ID Based Translations](https://doc.qt.io/qt-5/linguist-id-based-i18n.html)
* [Qt Linguist Manual: Release Manager](https://doc.qt.io/qt-5/linguist-manager.html)
* [Qt Linguist Manual: Translators](https://doc.qt.io/qt-5/linguist-translators.html)

* [Weblate supported formats](https://docs.weblate.org/en/latest/formats.html#formats)
* [Weblate Qt ts file support](https://docs.weblate.org/en/latest/formats.html#qt-linguist-ts)
* [Weblate bilingual and monolingual formats](https://docs.weblate.org/en/latest/formats.html#bimono)
* [Weblate Translation types capabilities](https://docs.weblate.org/en/latest/formats.html#translation-types-capabilities)
