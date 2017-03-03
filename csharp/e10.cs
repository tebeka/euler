/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.

Answer:
*/
using System;
using System.Collections;

class e10 {

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
        long sum = 0;
        foreach (long p in primes()) {
            if (p > 2000000) {
                break;
            }

            Console.WriteLine(p);
            sum += p;
        }

        Console.WriteLine(sum);
    }
}
        
