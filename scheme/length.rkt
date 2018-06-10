#lang racket

(
 (lambda (make-length)
   ((lambda (recurser) (recurser recurser))
    (lambda (recurser)
      (make-length (lambda (list) ((recurser recurser) list))))))

 (lambda (length)
   (lambda (list)
     (cond
       ((null? list) 0)
       (else (+ 1 (length (cdr list)))))))
)
