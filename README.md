GoSPEL
===============

This is an effort to learn Go Language and to understand Computer Architecture.

This effort started as early as 2006, has been unsuccessful implementation in C, C++ and Python.
With introduction of Go, this has got new wings.

GoSPEL is a hobby project which implements well defined interface for writing
expanding emulator/simulator.

Some of GoSPEL's  Features are

	  - Functional model
	  - Cycle accurate model (memory and caches)
	  - Multi-Arch support
	  - Multi-Platform support
	  - Plan9's ACID type debugger

Directory Tree:
---------------

<pre><code>
	  /-+---- plat/	- Platform specific directories
	    |      |
	    |      |______ mips/
	    |      |       |
	    |      |       +--- malta.go
	    |      |       +--- sead3.go
	    |      |       +--- other.go
	    |      |       +----- cavium/ (Cavium Specific boards, not implemented)
	    |      |       |
	    |      |       +----- <other>
	    |      |
	    |      |______ arm/
	    |              |
	    |              +--- realvieweb.go
	    |              +--- realviewpb.go
	    |              +--- vexpress.go
	    |              +---
	    |              +----- qualcomm/ (not Implemented)
	    |              +----- ti/  (not Implemented)
	    |
	    +---- cpu/	- Arch specific, also contains cache/TLB/ any extra
	    |      |      instructoins implemented
	    |      |
	    |      +-- cpu.go
	    |      +-- misc.go
	    |      +-- core.go
	    |      +-- regs.go
	    |      |
	    |      |_____ mips/
	    |      |       |
	    |      |       +--- mips.go
	    |      |       +--- cop0.go
	    |      |       +--- misc.go
	    |      |       |
	    |      |       +----- cavium/ (not Implemented)
	    |      |
	    |      |_____ arm/
	    |              |
	    |              +--- main.go
	    |              +--- cp15.go
	    |              +--- v4.go
	    |              +--- v5.go
	    |              +--- v6.go
	    |              +--- v7.go
	    |              |
	    |              +----- ti/ (not Implemented)
	    |              +----- qualcomm/ (not Implemented)
	    |
	    +---  dev/	- contains platform specific peripherals
	    |      |
		|      +--- dev.go (Generic Device infrastructure)
	    |      |
	    |      |_____ net/
	    |      |       |
	    |      |       +----- cavium/  (not Implemented)
	    |      |
	    |      |_____ serial/
	    |      |       |
	    |      |       +----- ti/ (not Implemented)
	    |      |
	    |      |_____ pci/
	    |      |       |
	    |      |       +----- cavium/ (not Implemented)
	    |      |
	    |      |_____ char/
	    |              |
	    |              +----- ti/ (not Implemented)
	    |
	    +---  cmd/	 -  root of all handling like command line, config file etc.
	    |  		    and responsible for log, debug etc.
	    |
	    +---  utils/ -  Not sure yet what goes in here

</code></pre>
</p>
