#include <sys/un.h>
#include <sys/socket.h>
#include "tlpi_hdr.h"

#define SV_SOCKET_PATH "/tmp/us_xfr"
#define BUF_SIZE 100

#define sockaddr_un struct sockaddr_un
