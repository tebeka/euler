; A palindromic number reads the same both ways. The largest palindrome made
; from the product of two 2-digit numbers is 9009 = 91 × 99.
;
; Find the largest palindrome made from the product of two 3-digit numbers.
;
; Answer: 906609

(defn palindrom? [s]
  (let [size (count s)
        middle (int (/ size 2))
        pre (.substring s 0 middle)
        i (if (even? size) middle (inc middle))]
    (= (list* pre) (reverse (.substring s i size)))))


(defn palindrom-numbers []
  (for [a (range 100 999)
        b (range 100 999)
        :when (palindrom? (format "%d" (* a b)))]
    (* a b)))

(prn (apply max (palindrom-numbers)))
