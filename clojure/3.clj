; The prime factors of 13195 are 5, 7, 13 and 29.
;
; What is the largest prime factor of the number 600851475143 ?
;
; Answer: 6857

(defn first-candidate [n]
 (let [c (int (Math/sqrt n))]
  (if (even? c)
   (dec c)
   c)))

(defn dec2 [n]
 (- n 2))

(defn biggest-factor [n]
 (first (filter (fn [x] (and (zero? (mod n x)) 
                             (.isProbablePrime (bigint x) 10)))
                (iterate dec2 (first-candidate n)))))

(println (biggest-factor 600851475143))
