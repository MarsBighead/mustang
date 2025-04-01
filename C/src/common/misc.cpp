#include<unistd.h>
#ifdef __APPLE__
#include <crt_externs.h>
#define environ (*_NSGetEnviron())
#else
extern char **environ;
#endif
#define _GNU_SOURCE
#include <string.h>

void print_environ(void)
{
    int i;

    //write_runlog(LOG, "begin printing environment variables.\n");
    for (i = 0; environ[i] != NULL; i++) {
        if (strcasestr(environ[i], "SESSION_ID") != NULL || strcasestr(environ[i], "PASSWD") != NULL) {
            continue;
        }
        //write_runlog(LOG, "%s\n", environ[i]);
    }
    //write_runlog(LOG, "end printing environment variables\n");
}