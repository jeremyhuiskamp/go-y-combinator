#lang racket

(lambda (list)
  (cond
    ((null? list) 0)
    (else (+ 1 (length (cdr list))))))
