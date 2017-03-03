; We shall say that an n-digit number is pandigital if it makes use of all the
; digits 1 to n exactly once; for example, the 5-digit number, 15234, is 1
; through 5 pandigital.
; 
; The product 7254 is unusual, as the identity, 39 × 186 = 7254, containing
; multiplicand, multiplier, and product is 1 through 9 pandigital.
; 
; Find the sum of all products whose multiplicand/multiplier/product identity
; can be written as a 1 through 9 pandigital.
;
; HINT: Some products can be obtained in more than one way so be sure to only include it once in your sum.
;
; Answer:

; FIXME: This is wrong!
;
(use 'clojure.contrib.combinatorics)
(use '[clojure.set :only (difference)])

(def *digits* #{\1 \2 \3 \4 \5 \6 \7 \8 \9})
(def *sizes* '(1 2 3 4))

(defn pandigital? [lhs rhs product]
  (= (sort *digits*) (sort (str lhs rhs product))))

(defn all-perms [nums sizes]
  (mapcat identity 
          (for [size sizes] (mapcat permutations (combinations nums size)))))

(defn all-pairs [digits sizes]
  (for [lhs (all-perms digits sizes)
        rhs (all-perms (difference digits lhs) sizes)]
    [lhs rhs]))

(defn chars->num 
  "(\1 \2 \3) -> 123"
  [cs]
  (Integer/valueOf (apply str cs)))

(defn map-fn [[lhs rhs]]
  (let [lhs (chars->num lhs) rhs (chars->num rhs) product (* lhs rhs)]
    (if (pandigital? lhs rhs product)
      product)))

(let [products (pmap map-fn (all-pairs *digits* *sizes*))]
  (println (reduce + (filter (complement nil?) products))))
(shutdown-agents)
