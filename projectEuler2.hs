fibs = 1 : 1 : zipWith (+) fibs (tail fibs)
ans = sum (takeWhile (<4000000) (filter even fibs))
ans
