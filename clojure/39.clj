; If p is the perimeter of a right angle triangle with integral length sides,
; {a,b,c}, there are exactly three solutions for p = 120.
; 
; {20,48,52}, {24,45,51}, {30,40,50}
; 
; For which value of p ≤ 1000, is the number of solutions maximised?
;
; Answer: 840

(defn triangle? [a b c]
  (= (* c c) (+ (* a a) (* b b))))

(defn triangles [n]
  (for [a (range 1 (- n 1)) b (range a (- n a 1)) 
        :when (triangle? a b (- n a b))]
    [a b (- n a b)]))

(defn max-triangles [max-p]
  (loop [p 1 best 1 best-n 0]
    (if (> p max-p)
      best
      (let [n (count (triangles p))]
        (if (> n best-n)
          (recur (inc p) p n)
          (recur (inc p) best best-n))))))

(prn (max-triangles 1000))
