# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=gl/glu

GOFILES:=glu_defs.go
CGOFILES:=glu.go

GC="${O}g" -I ../gl/_obj

include $(GOROOT)/src/Make.pkg

DOLLAR:="$"

glu_defs.go: glu_defs.c
	godefs `echo -n $(CGO_CFLAGS) | sed 's/\([^ ^$(DOLLAR)]*\)/-f \1/g'` -g glu glu_defs.c > glu_defs.go
	gofmt -w glu_defs.go
