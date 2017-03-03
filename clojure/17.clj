; If the numbers 1 to 5 are written out in words: one, two, three, four, five,
; then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
; 
; If all the numbers from 1 to 1000 (one thousand) inclusive were written out
; in words, how many letters would be used?
; 
; NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and
; forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20
; letters. The use of "and" when writing out numbers is in compliance with
; British usage.
;
; Answer: 21124

(def numbers {
    0 "" 1 "one" 2 "two" 3 "three" 4 "four" 5 "five" 6 "six" 7 "seven" 
    8 "eight" 9 "nine" 10 "ten" 11 "eleven" 12 "twelve" 13 "thirteen" 
    14 "fourteen" 15 "fifteen" 16 "sixteen" 17 "seventeen" 18 "eighteen"
    19 "nineteen" 20 "twenty" 30 "thirty" 40 "forty" 50 "fifty"
    60 "sixty" 70 "seventy" 80 "eighty" 90 "ninety" })

(def tens ["" "" "hundred" "thousand"])

(defn div [n i]
  (int (/ n i)))

(defn num->eng [n]
  (if (> n 999)
    ; We know it can be only up to 1000
    (concat ["one" "thousand"] (num->eng (rem n 1000)))
    (if (<= n 20)
      (let [v (numbers n)]
        (if (empty? v) [] [v]))
      (if (> n 99)
        (if (= n (* 100 (div n 100)))
          [(numbers (div n 100)) "hundred"]
          (concat [(numbers (div n 100)) "hundred" "and"] (num->eng (rem n 100))))
        (concat [(numbers (* (div n 10) 10))] (num->eng (rem n 10)))))))

(defn num-length [n]
  (apply + (map count (num->eng n))))

(prn (apply + (map num-length (range 1 1001))))
