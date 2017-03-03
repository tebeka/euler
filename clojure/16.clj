; 2^(15) = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

; What is the sum of the digits of the number 2^(1000)?
;
; Answer: 1366

(defn char->int [c]
  (- (int c) (int \0)))

(defn pow [a b]
  (loop [acc a b b]
    (if (= b 1)
      acc
      (recur (* a acc) (dec b)))))

(let [s (str (pow 2 1000))]
  (prn (apply + (map char->int s))))
