import random
import math
import time

def piFinder(dots):
    insideCircle = 0
    for i in range(dots):
        # print(i)
        x = random.uniform(0, 1)
        y = random.uniform(0, 1)
        if math.sqrt(x*x + y*y) <= 1:
            insideCircle += 1
    return 4 * insideCircle / dots

def main():
    dots = int(input("Número de pontos: "))
    if dots < 0:
        return print("O número de pontos deve ser positivo.")
    start = time.time()
    value = piFinder(dots)
    end = time.time()
    print("Encontramos 𝝅 =", value)
    print(end - start)


if __name__ == '__main__':
    main()