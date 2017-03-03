; Let d(n) be defined as the sum of proper divisors of n (numbers less than n
; which divide evenly into n).
; If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and
; each of a and b are called amicable numbers.
; 
; For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55
; and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71
; and 142; so d(284) = 220.
; 
; Evaluate the sum of all the amicable numbers under 10000.
;
; Answer: 31626


(defn factors [n]
  (let [max (inc (int (Math/sqrt n)))]
   (loop [fs [] i 2]
     (if (= i max)
       (set (cons 1 fs)) ; Unique items
       (if (zero? (mod n i))
         (recur (concat fs [i (/ n i)]) (inc i))
         (recur fs (inc i)))))))

(defn d [n]
  (apply + (factors n)))

(defn amicable? [n]
  (let [i (d n)]
    (and (not (= n i)) (= (d i) n))))

(prn (apply + (filter amicable? (range 1 10001))))
