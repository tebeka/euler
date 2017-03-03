; The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
; 
; Find the sum of all the primes below two million.
;
; Answer: 142913828922

(defn primes []
  (filter (fn [n] (.isProbablePrime (bigint n) 10)) (iterate inc 1)))

(prn (apply + (take-while #(<= % 2000000) (primes))))
