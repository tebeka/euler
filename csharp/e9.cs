/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
a^2 + b^2 = c^2

For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.  Find
the product a*b*c.

Answer: 31875000
*/
using System;

class e9 {
    static void Main() {
        bool seeking = true;
        for (int a = 1; (a <= 1000 && seeking); a++) {
            for (int b = 1; (b <= 1000 && seeking); b++) {
                for (int c = 1; (c <= 1000 && seeking); c++) {
                    if (((a*a + b*b) == (c * c)) && (a + b + c == 1000)) {
                        Console.WriteLine(a*b*c);
                        seeking = false;
                        break;
                    }
                }
            }
        }
    }
}
        
