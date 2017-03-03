; A permutation is an ordered arrangement of objects. For example, 3124 is one
; possible permutation of the digits 1, 2, 3 and 4. If all of the permutations
; are listed numerically or alphabetically, we call it lexicographic order. The
; lexicographic permutations of 0, 1 and 2 are:
; 
; 012   021   102   120   201   210
; 
; What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4,
; 5, 6, 7, 8 and 9?
;
; Answer: 2783915460

(defn last-permutation? [n]
  (= n (reverse (sort n))))

(defn next-head [n]
  (first (filter #(> % (first n)) (sort (rest n)))))

(defn permutation-step [n]
  (if (last-permutation? n)
    nil
    (let [sub (permutation-step (rest n))]
      (if sub
        (cons (first n) sub)
        (let [head (next-head n)]
          (cons head (sort (cons (first n) (remove #(= % head) (rest n))))))))))

(defn perm->str [n]
  (apply str (map #(str (bigint %)) n)))

(let [p (nth (iterate permutation-step (range 10)) 999999)]
  (prn (perm->str p)))
