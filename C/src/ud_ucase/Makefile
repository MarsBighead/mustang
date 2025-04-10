BASE=$(shell pwd)
MUSTANG_TOP_BUILDDIR =$(BASE)/../..

include $(MUSTANG_TOP_BUILDDIR)/Makefile.common

GEN_EXE = xfr_cl xfr_sv 

LDLIBS += $(addprefix -l,$(MUSTANG_BASE_LIBS))

OBJECTS=$(patsubst %.c,%,$(wildcard *.c))

all : $(OBJECTS) 

clean :
	${RM} ${GEN_EXE} *.dSYM

show :
	@echo "LIBS ${LIBS}"
	@echo "OBJECTS ${OBJECTS}"
	@echo "CPPFLAGS ${CPPFLAGS}"
	@echo "LDFLAGS ${LDFLAGS}"
