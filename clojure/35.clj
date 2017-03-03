; The number, 197, is called a circular prime because all rotations of the
; digits: 197, 971, and 719, are themselves prime.
; 
; There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71,
; 73, 79, and 97.
; 
; How many circular primes are there below one million?
;
; Answer: 55

(defn rotations [n]
  (let [n (str n) size (count n)]
    (for [i (range size)]
      (Integer/parseInt (apply str (.substring n i) (.substring n 0 i))))))

(defn primes-seq []
  (filter (fn [n] (.isProbablePrime (bigint n) 10)) (iterate inc 1)))

(defn primes-upto [n]
  (take-while #(< % n) (primes-seq)))

(defn circular? [n primes]
  (every? #(contains? primes %) (rotations n)))

(defn num-circular [n]
  (let [primes (set (primes-upto n))]
    (count (filter #(circular? % primes) primes))))

(prn (num-circular 1000000))
