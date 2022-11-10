def alphabetTriangle():
    for i in range(ord('A'), ord('G')+1):
        for j in range(ord('A'), i+1):
            print(chr(j), end=' ')
        print()
        
def main():
    alphabetTriangle()
    
if __name__ == "__main__":
    main()