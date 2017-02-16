# Go plugins in 1.8

<p align="center">
  <img src="img/gopher.png" width="180"> </image>
</p>

### What is a Go plugin?

##### The Plugin (shared object)

A Go plugin is essentially a shared object. We recognize these from our close neighbors in the `C` and `C++` programming languages.

Go plugin's are *NOT* part of the original program. They are standalone binaries that adhere to an ABI (Application Binary Interface) that another Go program can chose to attempt to run.


##### The Program

A Go program can choose to implement a Go plugin (remember this is a shared object or `.so` file) at *runtime*. This is huge because we no loner have to recompile anything to drastically change the
behavior of a Go program.


# Demo

### Attach to the official `golang:1.8` docker container

```bash
make
```

Which is essentially a wrapper for

```bash
docker run \
    -i \
    -t \
    -v $GOPATH/src/github.com/kris-nova/go-plugin-demo:/go/src/github.com/kris-nova/go-plugin-demo \
    -w /go/src/github.com/kris-nova/go-plugin-demo \
    --rm 
```

### Compile all the things

From the docker container we can go ahead and natively compile the main program, as well as all the plugins.

```bash
make build
```

### Run the program

By default we will be running `plugin1`. Run the program with

```bash
make run
```

### Change the plugin at runtime

```bash
export PLUGIN_NUMBER=2
make run
```

# Inspecting our plugins

### Why are they linux only?

Let's look at the Go source code [here](https://github.com/golang/go/tree/release-branch.go1.8/src/plugin)

```C
#cgo linux LDFLAGS: -ldl
#include <dlfcn.h>
#include <limits.h>
#include <stdlib.h>
#include <stdint.h>
```

This includes the `dlfcn.h` file, and uses the traditional Linux linking functions. Exactly the same as this prototype:

```C

#include <stdlib.h>
#include <stdio.h>
#include <dlfcn.h>

int main(int argc, char **argv) {
    void *handle;
    void (*run)();
    char *error;

    handle = dlopen ("../plugins/plugin1.so", RTLD_LAZY);
    if (!handle) {
        fputs (dlerror(), stderr);
        printf("\n");
        exit(1);
    }

    run = dlsym(handle, "plugin/unnamed-4dc81edc69e27be0c67b8f6c72a541e65358fd88.init");
    if ((error = dlerror()) != NULL)  {
        fputs(error, stderr);
        printf("\n");
        exit(1);
    }

    (*run)();
    dlclose(handle);
}
```

This gives us a hint into how Go plugins work, they use POSIX dynamic loading [more information](https://en.wikipedia.org/wiki/Dynamic_loading).

Right now there is only support for handling the linux version in the C implementation. The good news is that there is already resources for building shared objects for Windows and other archtypes.


