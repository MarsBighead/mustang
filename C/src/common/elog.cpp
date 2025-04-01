#include "common/elog.h"
#include <iostream>

#define _(x) x

volatile int log_min_messages = WARNING;



/*
 * is_log_level_output -- is elevel logically >= log_min_level?
 *
 * We use this for tests that should consider LOG to sort out-of-order,
 * between ERROR and FATAL.  Generally this is the right thing for testing
 * whether a message should go to the postmaster log, whereas a simple >=
 * test is correct for testing whether the message should go to the client.
 */

static bool is_log_level_output(int elevel, int log_min_level)
{
    if (elevel == LOG) {
        if (log_min_level == LOG || log_min_level <= ERROR) {
            return true;
        }
    } else if (log_min_level == LOG) {
        /* elevel not equal to LOG */
        if (elevel >= FATAL) {
            return true;
        }
    } else if (elevel >= log_min_level) {
        /* Neither is LOG */
        return true;
    }

    return false;
}


/*
 * Write errors to stderr (or by equal means when stderr is
 * not available).
 */
void write_runlog(int elevel, const char* fmt, ...)
{
    /* Get whether the record will be logged into the file. */
    if (!is_log_level_output(elevel, log_min_messages)) {
        return;
    }

    /* Obtaining international texts. */
    fmt = _(fmt);

    va_list ap;
    va_start(ap, fmt);
    //WriteRunLogv(elevel, fmt, ap);
    va_end(ap);
}