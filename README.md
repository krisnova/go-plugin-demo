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

### Why are they linux only? Why are they so big?

```bash
ls -lh | grep so

-rw-r--r-- 1 root root 3.1M Feb 16 16:42 plugin1.so
-rw-r--r-- 1 root root 3.1M Feb 16 16:42 plugin2.so

```

Let's look at the source code [here](https://github.com/golang/go/tree/release-branch.go1.8/src/plugin)

This gives us a hint into how Go plugins work. I have some experimental work in the [experimental_c_plugins](experimental_c_plugins) directory that offers an example of implementing a Go plugin in native C.


