/*
2520 is the smallest number that can be divided by each of the numbers from 1 to
10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the
numbers from 1 to 20?

Answer: 232792560
*/
using System;

class e5 {
    static bool isDiv20(ulong n) {
        for (ulong i = 2; i <= 20; i++) {
            if ((n % i) != 0) {
                return false;
            }
        }
        return true;
    }

    static void Main(string[] args) {
        ulong n = 20;
        while (true) {
            if (isDiv20(n)) {
                Console.WriteLine(n);
                break;
            }
            n++;
        }
    }
}
        
