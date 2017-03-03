; For a positive integer n, let σ(n) be the sum of all divisors of n, so e.g.
; σ(6) = 1 + 2 + 3 + 6 = 12.
; 
; A perfect number, as you probably know, is a number with σ(n) = 2n.
; 
; Let us define the perfection quotient of a positive integer as 
; p(n)	= σ(n) / n
; 
; 
; Find the sum of all positive integers n ≤ 10^(18) for which p(n) has the form
; k + 1⁄2, where k is an integer.

(defn factors [n]
  (let [max (inc (int (Math/sqrt n)))]
   (loop [fs [] i 2]
     (if (= i max)
       (set (cons n (cons 1 fs))) ; Unique items
       (if (zero? (mod n i))
         (recur (concat fs [i (/ n i)]) (inc i))
         (recur fs (inc i)))))))

(defn p [n]
 (/ (apply + (factors n)) n))

(defn k? [n]
  (let [pn (p n)]
    (and (= (class pn) clojure.lang.Ratio)
         (= (.denominator pn) 2))))

(defn sum-k [max-n]
  (apply + (pmap #(if (k? %) % 0) 
                 (take-while #(< % max-n) (iterate inc 1)))))

(prn (sum-k (Math/pow 10 18)))
(shutdown-agents)
