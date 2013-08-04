#!/usr/bin/env python2
import os,sys

from os.path import *

home = os.environ['HOME']
scons_dir = ".scons/site_scons/site_tools"

gotooldir = join(home, scons_dir)

if not exists(gotooldir):
	os.makedirs(gotooldir)

gotoollink = join(gotooldir,'go')

srcgotool = abspath(join(dirname(__file__), 'go'))

if not exists(join(gotooldir, 'go')):
	print gotoollink, srcgotool
	try:
		os.symlink(srcgotool, gotoollink)
	except:
		print "error"
        print "Symlink created re-run scons"
        sys.exit(0)
