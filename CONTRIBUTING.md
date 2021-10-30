# Mumble Website Contributing Guidelines

We appreciate pull requests for any issues you find.

If you want to make more substantial content, layout or structure changes, an issue ticket should be used for documenting and discussing the changes, to make sure you are not putting a lot of work into changes we disagree with.

Simple typo fixes and obvious text changes do not need elaborate documentation. You can suggest these in a pull request.

When creating more elaborate suggestions and changes please consider that the other party needs some introduction to your idea and changes. Laying out changes, screenshots, and reasoning helps speed up getting everyone on the same discussion base, and making a conclusive decision.

## Hugo Code Structure

Because we do not use themes, all code files like layouts and templates remain in a top level layouts folder.

## Text Content

The front matter should be defined in YAML `---` for consistency to existing content. We want to move to TOML `+++` in the future though.

All general content should be understandable to inexperienced readers. Provide introductory context where it makes sense. Complexity and knowledge requirements may increase in sub-sections (like developer or server administrator documentation).

Content should be created in Markdown (`.md` files) wherever possible. When Markdown alone is no longer enough, we embed HTML within the Markdown text/source code.
