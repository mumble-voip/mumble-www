---
title: "wiki Shortcode Test Page"
draft: true
---

# The `{{</* wiki /*/>}}` Shortcode

## Home

| Shortcode           | Result        |
| ------------------- | ------------- |
| `{{</* wiki /*/>}}` | {{< wiki />}} |

## Positional parameters

| Shortcode                                            | Result                                         |
| ---------------------------------------------------- | ---------------------------------------------- |
| `{{</* wiki Skinning /*/>}}`                         | {{< wiki Skinning />}}                         |
| `{{</* wiki "ACL Tutorial/Deutsch" /*/>}}`           | {{< wiki "ACL Tutorial/Deutsch" />}}           |
| `{{</* wiki Skins MyTitle /*/>}}`                    | {{< wiki Skins MyTitle />}}                    |
| `{{</* wiki Skins "Spa ced" /*/>}}`                  | {{< wiki Skins "MyTitle Spaced" />}}           |
| `{{</* wiki "ACL Tutorial/Deutsch" "Spa ced" /*/>}}` | {{< wiki "ACL Tutorial/Deutsch" "Spa ced" />}} |

## Named Parameters

| Shortcode                                                             | Result                                                      |
| --------------------------------------------------------------------- | ----------------------------------------------------------- |
| `{{</* wiki path="Contributing" /*/>}}`                               | {{< wiki path="Contributing" />}}                           |
| `{{</* wiki path="ACL Tutorial/Deutsch" /*/>}}`                       | {{< wiki path="ACL Tutorial/Deutsch" />}}                   |
| `{{</* wiki path="ACL_Tutorial/Deutsch" /*/>}}`                       | {{< wiki path="ACL_Tutorial/Deutsch" />}}                   |
| `{{</* wiki path="Contributing" */>}}Contr.{{</* /wiki */>}}`         | {{< wiki path="Contributing" >}}Contr.{{</wiki>}}           |
| `{{</* wiki path="ACL Tutorial/Deutsch" */>}}Tut DE{{</* /wiki */>}}` | {{< wiki path="ACL Tutorial/Deutsch" >}}Tut DE{{< /wiki >}} |

# Test

Within {{< wiki "Test" />}} Text.

No{{< wiki "Test" />}}Space.
