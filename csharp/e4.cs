/*
A palindromic number reads the same both ways. The largest palindrome made from
the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.

Answer: 906609
*/
using System;

class e4 {
    static bool isPalindrom(string s) {
        int left = 0;
        int right = s.Length - 1;

        while (left < right) {
            if (s[left] != s[right]) {
                return false;
            }
            left++;
            right--;
        }

        return true;
    }

    static void Main(string[] args) {
        int p = 0;
        for (int i = 100; i < 1000; ++i) {
            for (int j = 100; j < 1000; ++j) {
                int prod = i * j;
                if ((prod > p) && (isPalindrom(prod.ToString()))) {
                    p = prod;
                }
            }
        }

        Console.WriteLine(p);

    }
}
        
