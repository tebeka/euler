; We shall say that an n-digit number is pandigital if it makes use of all the
; digits 1 to n exactly once. For example, 2143 is a 4-digit pandigital and is
; also prime.
; 
; What is the largest n-digit pandigital prime that exists?
;
; Answer:

(defn str->digits [s]
  (sort (map #(- (int %) (int \0)) s)))

(defn pandigital? [n]
  (let [n (str n)]
    (= (range 1 (inc (count n))) (str->digits n))))

(defn primes-from [n]
  (filter (fn [n] (.isProbablePrime (bigint n) 10)) 
          (take-while #(> % 1) (iterate dec n))))

(prn (max (filter pandigital? (primes-from 987654321))))
