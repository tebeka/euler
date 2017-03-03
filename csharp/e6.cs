/*
The sum of the squares of the first ten natural numbers is,
1^2 + 2^2 + ... + 10^2 = 385

The square of the sum of the first ten natural numbers is,
(1 + 2 + ... + 10)^2 = 552 = 3025

Hence the difference between the sum of the squares of the first ten natural
numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred
natural numbers and the square of the sum.

Answer: 25164150
*/
using System;

class e6 {
    static ulong SumOfSuares(ulong n) {
        ulong sum = 0;
        for (ulong i = 1; i <= n; i++) {
            sum += (i * i);
        }

        return sum;
    }

    static ulong SquareOfSums(ulong n) {
        ulong sum = 0;
        for (ulong i = 1; i <= n; i++) {
            sum += i;
        }

        return sum * sum;
    }

    static void Main(string[] args) {
        ulong n = 100;
        Console.WriteLine(SquareOfSums(n) - SumOfSuares(n));
    }
}
        
