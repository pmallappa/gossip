#! /usr/bin/env python2.6
# -*- mode: python; coding: utf-8; -*-

import os, sys, shelve, shutil
from os.path import join

# Way to get project root from any script importing this one :-)
PROJROOT = os.path.abspath(os.path.join(os.path.dirname(__file__), '../..'))
TOPDIR = os.path.abspath(os.path.join(os.path.dirname(__file__), '..'))

BUILDDIR        = join(PROJROOT, 'build')

projpaths = {
    'BUILDDIR'          : BUILDDIR,
    'TOPDIR'		: TOPDIR,
}

