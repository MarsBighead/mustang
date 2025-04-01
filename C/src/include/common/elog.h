#ifndef ELOG_H
#define ELOG_H
#endif
/* Error level codes */
#define DEBUG5                                 \
    10 /* Debugging messages, in categories of \
        * decreasing detail. */

#define DEBUG1 14 /* used by GUC debug_* variables */
#define LOG                                         \
    15 /* Server operational messages; sent only to \
        * server log by default. */

#define WARNING                                      \
    19 /* Warnings.  NOTICE is for expected messages \
        * like implicit sequence creation by SERIAL. \
        * WARNING is for unexpected messages. */
#define ERROR                                                 \
    20           /* user error - abort transaction; return to \
                  * known state */
#define FATAL 22 /* fatal error - abort process */