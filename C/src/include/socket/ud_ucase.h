#include <sys/un.h>
#include <sys/socket.h>
#include <ctype.h>
#include "tlpi_hdr.h"

#define SV_SOCKET_PATH "/tmp/ud_ucase"
#define BUF_SIZE 10

#define sockaddr_un struct sockaddr_un
