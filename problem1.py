def sumMult_3_5(n):
    ans = 0
    for i in range(1,n//3+1):
        ans += i * 3
        if i%5 == 0:
            ans -= i * 3
        if i < n//5:
            ans += i * 5
    return ans

def main(): 
    print(sumMult_3_5(1000)) 
  
  
# Using the special variable  
# __name__ 
if __name__=="__main__": 
    main()
