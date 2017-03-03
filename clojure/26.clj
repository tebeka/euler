; A unit fraction contains 1 in the numerator. The decimal representation of the
; unit fractions with denominators 2 to 10 are given:
; 
;     1/2	= 	0.5
;     1/3	= 	0.(3)
;     1/4	= 	0.25
;     1/5	= 	0.2
;     1/6	= 	0.1(6)
;     1/7	= 	0.(142857)
;     1/8	= 	0.125
;     1/9	= 	0.(1)
;     1/10 	= 	0.1
; 
; Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle. It can be
; seen that 1/7 has a 6-digit recurring cycle.
; 
; Find the value of d < 1000 for which 1/d contains the longest recurring
; cycle in its decimal fraction part.
; 
;


; The idea is to create step of divison that will return the digit and the
; leftover and run it with "iterate". Currently it seems like the best approach
; is to have an embedded function where "n" is bound (or maybe use "bindings")
(defn step [[digit a b]]
  ; (divide 1 7) -> [1 3]
  (let [a (* 10 a)
        digit (int (/ a b))
        leftover (- a (* b digit))]
     [digit leftover b]))

(defn digits [n]
  (defn step [digit leftover]
    
  (map first (iterate divide [1 n])))
