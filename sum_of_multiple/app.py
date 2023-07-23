def sum_mul(n, m):
    if m < n :
        return "INVALID"
    
    index = 1
    last = n
    sum = 0
    
    while (last <= m ):
        index += 1
        sum += last
        last = index * n
    return sum

if __name__ == "__main__":
    res = sum_mul(2, 9)
    print(res)