from threading import Thread, Lock
import random
import math
import time

class Variavel():
    def __init__(self):
        self.valor = 0
        self.lock = Lock()

    def incrementa(self, value):
        self.lock.acquire()
        self.valor += value
        self.lock.release()

    def getValor(self):
        return self.valor

class MonteCarloThread(Thread):
    def __init__(self, id, done, inside, dots):
        super().__init__()
        self.threadid = id
        self.done = done
        self.inside = inside
        self.dots = dots

    def run(self):
        while self.done.getValor() < self.dots:
            # print("Thread", self.threadid, 'running', self.done.getValor())
            x = random.uniform(0, 1)
            y = random.uniform(0, 1)
            if math.sqrt(x*x + y*y) <= 1:
                self.inside.incrementa(1) 
            self.done.incrementa(1) 

def main():
    dots = int(input("Número de pontos: "))
    n_threads = int(input("Número de threads: "))
    if dots < 0 or n_threads < 0:
        return print("Ambos os números devem ser positivos.")
    print(f"Faremos {n_threads} threads rodando {dots} pontos.\n")
    done = Variavel()
    inside = Variavel()
    start = time.time()
    threads = [MonteCarloThread(i, done, inside, dots) for i in range(n_threads)]
    for thread in threads:
        thread.start()
    for thread in threads:
        thread.join()
    value = 4 * inside.getValor() / done.getValor() 
    end = time.time()   
    print("Encontramos 𝝅 =", value)
    print(end - start)


if __name__ == '__main__':
    main()