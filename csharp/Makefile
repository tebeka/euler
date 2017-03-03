exe = $(patsubst %.cs,%.exe,$(wildcard *.cs))

%.exe: %.cs
	gmcs $<

all: $(exe)

clean:
	rm -f $(exe)

fresh: clean all

.PHONY: all clean fresh
