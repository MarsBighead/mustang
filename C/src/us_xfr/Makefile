BASE=$(shell pwd)
MUSTANG_TOP_BUILDDIR =$(BASE)/../..

include $(MUSTANG_TOP_BUILDDIR)/Makefile.common

GEN_EXE = cl sv 

LINUX_EXE = list_host_addresses \
	scm_cred_recv scm_cred_send \
	scm_multi_recv scm_multi_send \
	scm_rights_recv scm_rights_send \
	us_abstract_bind

#EXE = ${GEN_EXE} ${LINUX_EXE}
EXE = ${GEN_EXE}

all : 
	@echo "$(CC) $(CXXFLAGS) $(LDFLAGS) $(GEN_EXE)  $(LIBS) -o $@"
	$(CC) $(CPPFLAGS) $(LDFLAGS) $(LIBS) cl.c -o cl

clean :
	${RM} ${EXE} *.o

showall :
	@ echo ${EXE}

${EXE} : ${TLPI_LIB}		# True as a rough approximation
