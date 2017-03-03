; By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see
; that the 6^(th) prime is 13.
; 
; What is the 10001^(st) prime number?
;
; Answer: 104743

(defn primes []
  (filter #(.isProbablePrime (bigint %) 10) (iterate inc 1)))

(prn (nth (primes) 10000))
