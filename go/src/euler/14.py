#!/usr/bin/env python

_collatz = {1:1}
def collatz(n):
    value = _collatz.get(n)
    if value:
        return value

    if (n % 2) == 0:
        _collatz[n] = value = 1 + collatz(n/2)
    else:
        _collatz[n] = value = 1 + collatz(3*n+1)

    return value

def main():
    n, value = 1, 1
    for i in xrange(2, 1000001):
        c = collatz(i)
        if c > value:
            n, value = i, c

    print(n)

if __name__ == "__main__":
    main()
