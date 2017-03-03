# By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see
# that the 6^(th) prime is 13.
# 
# What is the 10001^(st) prime number?
#
# Answer: 104743

n = 2
idx = 1
while true
    if isprime(n)
        idx += 1
        if idx == 10000
            print(n)
            break
        end
    end
    n += 1
end
