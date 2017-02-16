
#include <dlfcn.h>
#include <limits.h>
#include <stdlib.h>
#include <stdint.h>
#include <stdio.h>

// Define the header here
uintptr_t pluginOpen(const char* path, char** err);
void* pluginLookup(uintptr_t h, const char* name, char** err);


int main() {


	// Plugin 1
	printf( "Running Plugin 1\n" );
	char* err[255];
	const char*  path = "plugins/plugin1.so";

	// Getting this formula took about 2 hours of work, so please be grateful.
	// I had to cross compile a tweaked version of the Go 1.8 release branch,
	// and do some weird magic with imports to get the standard library to write
	// the string value of the variable they are using to build this string with
	// to /tmp/pluginpath because of the import test that won't let a user add
	// any log or fmt to the package. But hey.. I finally compiled Go..
	const char*  name = "plugins/plugin1.so.init";

	uintptr_t h;
	void* p;
    h = pluginOpen(path, err);
    p = pluginLookup(h, name, err);
    (p)();

	printf( "Running Plugin 2\n" );
}

uintptr_t pluginOpen(const char* path, char** err) {
	void* h = dlopen(path, RTLD_NOW|RTLD_GLOBAL);
	if (h == NULL) {
		*err = (char*)dlerror();
	}
	return (uintptr_t)h;
}


void* pluginLookup(uintptr_t h, const char* name, char** err) {
	void* r = dlsym((void*)h, name);
	if (r == NULL) {
		*err = (char*)dlerror();
	}
	return r;
}