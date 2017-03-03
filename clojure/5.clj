; 2520 is the smallest number that can be divided by each of the numbers from 1
; to 10 without any remainder.
;
; What is the smallest number that is evenly divisible by all of the numbers
; from 1 to 20?
;
; Answer: 232792560

(defn divisible? [n]
  (every? #(zero? (mod n %)) (range 1 20)))

(prn (first (filter divisible? (iterate inc 19))))
