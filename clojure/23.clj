; A perfect number is a number for which the sum of its proper divisors is
; exactly equal to the number. For example, the sum of the proper divisors of
; 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.
; 
; A number whose proper divisors are less than the number is called deficient
; and a number whose proper divisors exceed the number is called abundant.
; 
; As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest
; number that can be written as the sum of two abundant numbers is 24. By
; mathematical analysis, it can be shown that all integers greater than 28123
; can be written as the sum of two abundant numbers. However, this upper limit
; cannot be reduced any further by analysis even though it is known that the
; greatest number that cannot be expressed as the sum of two abundant numbers
; is less than this limit.
; 
; Find the sum of all the positive integers which cannot be written as the sum
; of two abundant numbers.
;
; Answer: 4179871
 
(defn factors [n]
  (let [max (inc (int (Math/sqrt n)))]
   (loop [fs [] i 2]
     (if (= i max)
       (set (cons 1 fs)) ; Unique items
       (if (zero? (mod n i))
         (recur (concat fs [i (/ n i)]) (inc i))
         (recur fs (inc i)))))))

(def factors (memoize factors))

(defn abundant? [n]
  (> (apply + (factors n)) n))

(defn sum-two-abundant [n]
  (some (fn [i] (and (abundant? i) (abundant? (- n i))))
        (range 1 (inc (Math/ceil (/ n 2))))))

;(prn (apply + (filter (fn [i] (not (sum-two-abundant i))) (range 1 28123))))
(dorun (map #(println % (sum-two-abundant %)) (range 1 28123)))
