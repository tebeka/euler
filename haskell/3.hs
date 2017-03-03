-- The prime factors of 13195 are 5, 7, 13 and 29.
--
-- What is the largest prime factor of the number 600851475143 ?

roundSqrt :: Double -> Integer
roundSqrt = round . sqrt

firstCandidate :: Integer -> Integer
firstCandidate n = 
    let c = roundSqrt (fromIntegral n) in
        if even c then c - 1 else c

isPrime' _ 0 = False
isPrime' _ 1 = True
isPrime' n p = if n `mod` p == 0 then False else isPrime' n (p - 2)

isPrime n = 
    if n > 2 && even n then
        False
    else
        let p = firstCandidate n in
            isPrime' n p
first n = 
    let c = n / 2
        if even c then c - 1 else c

primeFactor n p = (n `mod` p == 0) && isPrime p
