import os
from os.path import join
from tools.projpaths import *

bdir = TOPDIR
idir = join(bdir, 'pkg')

if COMMAND_LINE_TARGETS:
    targets = COMMAND_LINE_TARGETS
else:
    targets = 'install'

env = Environment(TOOLS=['default', 'go'])
env.Alias('install', idir)
env.Default('install')

objects = SConscript('src/SConscript',
	exports = ['env', 'bdir', 'idir'],
	variant_dir='#build',
	duplicate = 0)

