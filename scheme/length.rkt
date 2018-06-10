#lang racket

(lambda (h)
  ((lambda (f) (f f))
   (lambda (f)
     (h (lambda (x) ((f f) x))))))

(lambda (list)
  (cond
    ((null? list) 0)
    (else (+ 1 (length (cdr list))))))
