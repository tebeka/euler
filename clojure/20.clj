; n! means n × (n − 1) × ... × 3 × 2 × 1
; 
; Find the sum of the digits in the number 100!
;
; Answer: 648

(defn fact [n]
 (apply * (range 1 (inc n))))


(defn char->int [c]
  (- (int c) (int \0)))

(let [s (str (fact 100))]
  (prn (apply + (map char->int s))))
