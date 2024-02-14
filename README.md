### wllama

This is an alternate front-end for the [ollama](https://ollama.ai) command line tools.
This is pretty much [yak shaving](https://en.wiktionary.org/wiki/yak_shaving)
but hey it's another chance to use the wonderful [Wails](https://wails.io)
application development library.

#### Installation and development

You must have [Go](https://go.dev) installed. Once that's set up, you should follow
the [Wails installation instructions](https://wails.io/docs/gettingstarted/installation)
to set that up. Then run:

    ./go

#### Problems with ollama

The main idea with `ollama` is that it runs locally, without having to hit the
internet. This is why this is an actual application, rather than a web front end.

The main irritations with `ollama` (at the moment) is that when you install it,
it *insists* on setting itself up as a login item
[see bug #162](https://github.com/jmorganca/ollama/issues/162),
and it also *insists* on having you install the command-line tools as admin
[see bug #283](https://github.com/jmorganca/ollama/issues/283), and will not run
its front end until you do.

You can get around the second one by installing the command line tools via (e.g.)
the homebrew command `brew install ollama`. Then, you can run the server component
in one terminal with the command

    /opt/homebrew/opt/ollama/bin/ollama serve

And then run

    ollama run codellama:7b-instruct

*Look out, yak*.
