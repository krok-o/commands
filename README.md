# plugins
Sample plugins for Krok

## Conventions

The plugins must follow the following structure:

As a first, read from stdin because that's how the payload is transferred. And have a single first argument
in which the settings for this command are passed in.

The plugin must be executable. It can be anything that can run and follows this convention.
