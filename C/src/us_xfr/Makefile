BASE=$(shell pwd)
MUSTANG_TOP_BUILDDIR =$(BASE)/../..

include $(MUSTANG_TOP_BUILDDIR)/Makefile.common

GEN_EXE = xfr_cl xfr_sv 

LINUX_EXE = list_host_addresses \
	scm_cred_recv scm_cred_send \
	scm_multi_recv scm_multi_send \
	scm_rights_recv scm_rights_send \
	us_abstract_bind


LDLIBS += $(addprefix -l,$(MUSTANG_BASE_LIBS))
#EXE = ${GEN_EXE} ${LINUX_EXE}
EXE = ${GEN_EXE}

all : $(EXE) 
#	$(CC) $(CPPFLAGS) $(LDFLAGS) $^ $(LIBS) -o $@
#	gcc  -o sv sv.c -L../lib -lmustang -I../include/xfr -I../include

clean :
	${RM} ${GEN_EXE} *.dSYM

show :
	@echo "LIBS ${LIBS}"
	@echo "CPPFLAGS ${CPPFLAGS}"
	@echo "LDFLAGS ${LDFLAGS}"

${EXE} : ${TLPI_LIB}		# True as a rough approximation
