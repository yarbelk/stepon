# Elephant Carpaccio

There is a fun exercise on story splitting, called elephant carpaccio.

After finishing this, I re-implemented it in golang, made a library
from said implementation, and then made an android app (binding 
the golang lib to the java).

make sure your androids are all set up, and your studios and jdks as
well.

change the gobind in `android/libstepon/build.gradle` to point to
your GOPATH and GO

Get the gomobile tool

    go get golang.org/x/mobile/cmd/gomobile
    gomobile init

to set up gomobile

then in the `android` dir,

    ./gradlew assemble

or whatever command you want.  YMMV.

also, from within the root dir, you can

    go build

and you will get a binary, `stepon`, which is a CLI interface.

see:

    ./stepon -h

It also accepts input from STDIN, where you must have

    [units] [prince] [state]

on each line of stdin.
