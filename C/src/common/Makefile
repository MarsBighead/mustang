BASE=$(shell pwd)
MUSTANG_TOP_BUILDDIR =$(BASE)/../..
include $(MUSTANG_TOP_BUILDDIR)/Makefile.common

MUSTANG_DIRECTORY_INC = $(MUSTANG_TOP_BUILDDIR)/src/include

CPPFLAGS += -I$(MUSTANG_DIRECTORY_INC)
CPPFLAGS += -I/usr/lib/gcc/aarch64-linux-gnu/12/include


all: libcommon.so

libcommon.so:
	$(CC) -fPIC -shared $(CXXFLAGS) $(CPPFLAGS) $(wildcard *.cpp) -o libcommon.so

clean :
	${RM} *.o *.so