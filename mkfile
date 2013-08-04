VPROC=`{pwd}
GOPATH=`{pwd}

<$VPROC/mk/mkhdr

MKSHELL=rc

install:V:

all:V: install

<| go env

DIRS=\
	src\

GOC=$GOTOOLDIR/${GOCHAR}g
GLD=$GOTOOLDIR/${GOCHAR}l
GOAS=$GOTOOLDIR/${GOCHAR}a
GOPACK=$GOTOOLDIR/pack

O=$GOCHAR

<$VPROC/mk/mkdirs

<$VPROC/mk/mkcommon

#NUKEFILES=.goenv
