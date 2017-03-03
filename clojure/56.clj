; A googol (10^(100)) is a massive number: one followed by one-hundred zeros;
; 100^(100) is almost unimaginably large: one followed by two-hundred zeros.
; Despite their size, the sum of the digits in each number is only 1.
; 
; Considering natural numbers of the form, a^(b), where a, b < 100, what is the
; maximum digital sum?
;
; Answer: 972

(defn char->num [n]
  (- (int n) (int \0)))

(defn sum-digits [n]
  (apply + (map char->num (str (bigint n)))))

(defn max-digits [n]
  (let [ab (for [a (range 1 n) b (range 1 n)] [(bigint a) (bigint b)])]
    (apply max (map #(sum-digits (.pow (first %) (last %))) ab))))

(prn (max-digits 100))
