; A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
; a^(2) + b^(2) = c^(2)
; 
; For example, 3^(2) + 4^(2) = 9 + 16 = 25 = 5^(2).
; 
; There exists exactly one Pythagorean triplet for which a + b + c = 1000.
; Find the product abc.
;
; Answer: 31875000
 
(defn triplets []
  (for [a (range 1 999)
        b (range 1 (- 999 a))]
    [a b (- 1000 (+ a b))]))

(defn pythagorean? [a b c]
  (= (+ (* a a) (* b b)) (* c c)))

(let [abc (first (filter #(apply pythagorean? %) (triplets)))]
  (prn (apply * abc)))
