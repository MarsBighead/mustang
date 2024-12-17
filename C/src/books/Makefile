BASE=$(shell pwd)
BOOK_TOP_BUILDDIR = ../..
SRC_PATH = $(BOOK_TOP_BUILDDIR)/src
COMM_LIB_PATH = $(SRC_PATH)/lib
BASE_BUILD_PATH = $(BOOK_TOP_BUILDDIR)/build

##########################################################################
# Enable options
##########################################################################

ifeq ($(USE_BUILD_MAKE), ON)
    BUILD_TYPE=$(MAKE_BUILD_TYPE)
else
    ifndef BUILD_TYPE
        BUILD_TYPE=Debug
    endif
endif


# Meta Configure 

PROJECT_SOURCE_DIR = $(BOOK_TOP_BUILDDIR)

CC = g++

CXXFLAGS=
CPPFLAGS=
LDFLAGS=
LIBS=

ifeq ($(BUILD_TYPE), Debug)
    BUILD_MODE = Debug
    OPTIMIZE_LEVEL = -O0 -g
else ifeq ($(BUILD_TYPE), Release)
    BUILD_MODE = Release
    ENABLE_MEMCHECK = OFF
    ENABLE_UT=OFF
    OPTIMIZE_LEVEL = -O2 -g3
else
    $(error unsupported BUILD_TYPE is $(BUILD_TYPE))
endif


COMM_INC = $(PROJECT_SOURCE_DIR)/src/include \
        $(PROJECT_SOURCE_DIR)/src/include/book


CPPFLAGS += $(addprefix -I,$(COMM_INC))

BOOK_LIB_PATH = $(PROJECT_SOURCE_DIR)/src/lib

OBJS = $(patsubst %.cpp,%.o,$(wildcard *.cpp))
LDFLAGS += -L$(BOOK_LIB_PATH)/book

all: books

# CXXFLAGS
books: $(OBJS)
	$(CC) -v $(CXXFLAGS) $(LDFLAGS) $^ $(LIBS) -o $@

clean:
	rm -f books *.o