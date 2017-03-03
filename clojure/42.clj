; The n^(th) term of the sequence of triangle numbers is given by, t_(n) =
; ½n(n+1); so the first ten triangle numbers are:
; 
; 1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...
; 
; By converting each letter in a word to a number corresponding to its
; alphabetical position and adding these values we form a word value. For
; example, the word value for SKY is 19 + 11 + 25 = 55 = t_(10). If the word
; value is a triangle number then we shall call the word a triangle word.
; 
; Using words.txt (right click and 'Save Link/Target As...'), a 16K text file
; containing nearly two-thousand common English words, how many are triangle
; words?
; 
; Answer: 162

(defn word->int [w]
  (apply + (map #(inc (- (int %) (int \A))) w)))

(defn triangles [n]
  (map #(* (/ % 2) (inc %)) (range 1 (inc n))))

(let [words (re-seq #"[A-Z]+" (slurp "words.txt"))
      values (map word->int words)
      ts (set (triangles (apply max values)))]
  (prn (count (filter #(contains? ts %) values))))
