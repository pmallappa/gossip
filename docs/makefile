Each directory should contain makefile like this

in dir:mips
===================
PROG=\
	m\
	n\

common_PROG_SRCS=p.c q.c
m_SRCS=m.c

LIB=\
    a\
    b\

DIRS=\
    x\
    y\

common_LIB_SRCS=p.c q.c

a_SRCS=a.c
b_SRCS=b.c

include $(ROOT_DIR)/common.mk
==================


This should build pkg/mips/a.a and pkg/mips/b.a 
pkg/mips/a.a should be built using common_SRCS + a_SRCS
pkg/mips/b.a should be built using common_SRCS + b_SRCS

then should recurse into each directory listed by DIRS


