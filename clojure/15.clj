; Starting in the top left corner of a 2×2 grid, there are 6 routes (without
; backtracking) to the bottom right corner.
; [15.gif]

; How many routes are there through a 20×20 grid?
;
; Answer: 137846528820


(defn num-paths [x y max-x max-y]
  (if (or (= x max-x) (= y max-y))
    1
    (+ (num-paths (inc x) y max-x max-y) (num-paths x (inc y) max-x max-y))))

(def num-paths (memoize num-paths))


(prn (num-paths 0 0 20 20))
