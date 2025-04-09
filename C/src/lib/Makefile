BASE=$(shell pwd)
MUSTANG_TOP_BUILDDIR =$(BASE)/../..

include $(MUSTANG_TOP_BUILDDIR)/Makefile.common
OBJECTS=$(patsubst %.c,%.o,$(wildcard *.c))


all : error_functions.o libmustang.a libmustang.so

libmustang.so: 
	$(CC) -fPIC -shared $(CXXFLAGS) $(CPPFLAGS) $(OBJECTS) -o libmustang.so 

error_functions.o : ename.inc.h

ename.inc.h :
	sh build_ename.sh > $(MUSTANG_TOP_BUILDDIR)/src/include/ename.inc.h
	echo 1>&2 "ename.inc.h built"

libmustang.a: $(OBJECTS)
	ar crs $@ $^

clean :
	${RM} *.o ename.c.inc 