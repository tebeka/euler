-- 2520 is the smallest number that can be divided by each of the numbers from 1
-- to 10 without any remainder.
-- 
-- What is the smallest positive number that is evenly divisible by all of the
-- numbers from 1 to 20?
-- 
-- Answer: 232792560

div20 n = all (\x -> n `mod` x == 0) [1..20]

main = print $ head [x | x <- [20,40..], div20 x]

-- nice short one is
-- foldl lcm 2 [3..20]
