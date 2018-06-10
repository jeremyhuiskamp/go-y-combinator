#lang racket

(define length
  (lambda (list)
    (cond
      ((null? list) 0)
      (else (+ 1 (length (cdr list)))))))

(length (quote (a b c)))
