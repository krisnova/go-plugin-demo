
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