BASE=$(shell pwd)
MUSTANG_TOP_BUILDDIR =$(BASE)/../..

include $(MUSTANG_TOP_BUILDDIR)/Makefile.common
OBJECTS=$(patsubst %.c,%.o,$(wildcard *.c))


all : error_functions.o lib

lib : $(OBJECTS)

error_functions.o : ename.inc.h

ename.inc.h :
	sh build_ename.sh > $(MUSTANG_TOP_BUILDDIR)/src/include/ename.inc.h
	echo 1>&2 "ename.inc.h built"
clean :
	${RM} *.o ename.c.inc 