# A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
# a^(2) + b^(2) = c^(2)
# 
# For example, 3^(2) + 4^(2) = 9 + 16 = 25 = 5^(2).
# 
# There exists exactly one Pythagorean triplet for which a + b + c = 1000.
# Find the product abc.
#
# Answer: 31875000

function triplets()
    function inner()
        for a = 1:999
            for b = 1:(999-a)
                produce(a, b, 1000-a-b)
            end
        end
    end
    return @task inner()
end

for (a, b, c) = triplets()
    if a^2 + b^2 == c^2
        println(a * b * c)
        break
    end
end
