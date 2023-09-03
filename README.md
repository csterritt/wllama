### wllama

This is an alternate front-end for the `[ollama](https://ollama.ai)` command line tools.
This is pretty much [yak shaving](https://en.wiktionary.org/wiki/yak_shaving)
but hey it's another chance to use the wonderful [Wails](https://wails.io)
application development library.

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

In another terminal to interact with the model. This is fine for small experiments,
but for larger things this contains *another* irritation,
[which has made it to a pull request!](https://github.com/jmorganca/ollama/pull/416)
which sets it up to do a prompt *per line* instead of per (say) paragraph. So if you
want to create a big prompt, it's got to be on one line. *Look out, yak*.
