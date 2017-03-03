; See http://projecteuler.net/index.php?section=problems&id=96 :)
;
; Answer: 

(refer 'clojure.set :only '(difference))

(def problem '(
  "003020600"
  "900305001"
  "001806400"
  "008102900"
  "700000008"
  "006708200"
  "002609500"
  "800203009"
  "005010300"))

(defn char->num [c]
  (- (int c) (int \0)))

(defstruct cell :row :col)

(defn parse-row [grid row s]
  (loop [grid grid i 0]
    (if (= i (count s))
      grid
      (let [v (char->num (nth s i))]
        (if (zero? v)
          (recur grid (inc i))
          (recur (assoc grid (struct cell row i) v) (inc i)))))))

(defn parse-rows [rows]
  (loop [grid {} row 0]
    (if (= row (count rows))
      grid
      (recur (parse-row grid row (nth rows row)) (inc row)))))

(defn cell-row [c]
  (for [col (range 9)]
    (struct cell (:row c) col)))

(defn cell-col [c]
  (for [row (range 9)]
    (struct cell row (:col c))))

(defn box-first [n]
  (int (Math/floor (/ n 3))))

(defn cell-box [c]
  (let [row (box-first (:row c))
        col (box-first (:col c))]
    (for [r (range 3) c (range 3)]
      (struct cell (+ row r) (+ col c)))))

(defn cell-peers [c]
  (let [peers (set (concat (cell-row c) (cell-col c) (cell-box c)))]
    (disj peers c)))

