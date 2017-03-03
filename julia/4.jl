# A palindromic number reads the same both ways. The largest palindrome made
# from the product of two 2-digit numbers is 9009 = 91 × 99.
#
# Find the largest palindrome made from the product of two 3-digit numbers.
#
# Answer: 906609

function is_palindrom(n)
    n = string(n)
    return n == reverse(n)
end

function palindroms()
    function inner()
        for a = 100:999
            for b = 100:999
                n = a * b
                if is_palindrom(n)
                    produce(n)
                end
            end
        end
    end

    return @task inner()
end

println(maximum(palindroms()))
