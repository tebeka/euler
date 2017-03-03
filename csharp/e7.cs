/*
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that
the 6th prime is 13.

What is the 10001st prime number?

Answer: 
*/
using System;
using System.Collections;

class e7 {
    static IEnumerable longs() {
        long i = 2;
        while (true) {
            yield return i++;
        }
    }

    static IEnumerable filter(long p, IEnumerable nums) {
        foreach (long i in nums) {
            if ((i % p) != 0) {
                yield return i;
            }
        }
    }

    static long head(IEnumerable items) {
        foreach (long i in items) {
            return i;
        }

        // We'll never get here, just make the compiler happy
        return -1;
    }

    static IEnumerable primes() {
        IEnumerable stream = longs();
        while (true) {
            long p = head(stream);
            stream = filter(p, stream);
            yield return p;
        }
    }

    static void Main() {
        int n = 0;
        foreach (long p in primes()) {
            n++;
            if ((n % 100) == 0) {
                Console.WriteLine(n);
            }
            if (n == 10001) {
                Console.WriteLine(p);
                break;
            }
        }

    }
}
        
