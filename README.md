# Calling asdf tools recursively

I'm using `asdf` for most of my development work.

I'm currently working on a job where test runners spawn each other.
Our `nodejs` test runner starts up a `java` server for example,
with an empty env, except for `PATH`.

I've got all tools set up in `.tool-versions`.

When calling `java` from `nodejs` the error is:
```
unknown command: java. Perhaps you have to reshim?
```

This can be fixed by adding `ASDF_DATA_DIR` to the env in the `nodejs` code that
spawns `java`.

Is there a way of avoiding this?

## Reproduction

The code in this repo reproduces this situation, where `golang` calls into `nodejs`
without passing `ASDF_DATA_DIR`.

You can install the required `asdf` plugins:
```console
$ make asdf-plugins
```
Then install the tools:
```console
$ make install
```

You can run the `nodejs` program, to make sure it works:
```console
$ make run-node
node ./test.js
Hello from Node.js!
```

You can then run the `golang` program that calls `nodejs`:
```console
$ make run-go
go run ./test.go
Hello from go!
unknown command: node. Perhaps you have to reshim?
2024/01/25 12:36:37 exit status 1
exit status 1
make: *** [run-go] Error 1
```

You can fix the program by uncommenting line 18 in `test.go`.

```console
$ make run-go
go run ./test.go
Hello from go!
Hello from Node.js!
```
