/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

Answer: 6857
*/
using System;
using System.Collections;

class e3 {
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
        long n = 600851475143;
        long p = 0;
        foreach (long i in primes()) {
            if (i > n) {
                break;
            }
            if ((n % i) == 0) {
                p = i;
                n /= p;
            }
        }

        Console.WriteLine(p);
    }
}
        
