# 2520 is the smallest number that can be divided by each of the numbers from 1
# to 10 without any remainder.
#
# What is the smallest number that is evenly divisible by all of the numbers
# from 1 to 20?
#
# Answer: 232792560

function is_divisible(n::Int)
    return all(map(i::Int -> n%i == 0, 1:20))
end


for n = 20:typemax(Int)
    if n > 232792560
        println("ooops")
        break
    end

    if is_divisible(n)
        println(n)
        break
    end
end
