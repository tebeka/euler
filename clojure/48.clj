; The series, 1^(1) + 2^(2) + 3^(3) + ... + 10^(10) = 10405071317.

; Find the last ten digits of the series, 1^(1) + 2^(2) + 3^(3) + ... +
; 1000^(1000).
;
; Answer: 9110846700 

(defn power-seq [n]
  (map #(.pow % %) (map bigint (range 1 (inc n)))))

(let [n (str (apply + (power-seq 1000)))]
  (prn (.substring n (- (count n) 10))))

