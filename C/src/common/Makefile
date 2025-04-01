BASE=$(shell pwd)
MUSTANG_TOP_BUILDDIR =$(BASE)/../..

MUSTANG_DIRECTORY_INC = $(MUSTANG_TOP_BUILDDIR)/src/include

CPPFLAGS += -I$(MUSTANG_DIRECTORY_INC)


all: libcmcommon.so

libcmcommon.so:
	$(CC) -fPIC -shared $(CXXFLAGS) $(CPPFLAGS) $(wildcard *.cpp) -o libcmcommon.so