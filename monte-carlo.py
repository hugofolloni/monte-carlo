from threading import Thread, Lock
import random
import math
import time

class VariavelParcela():
    def __init__(self):
        self.valor = 0
        self.lock = Lock()

    def incrementa(self, value):
        self.lock.acquire()
        self.valor += value
        self.lock.release()

    def getValor(self):
        return self.valor

class MonteCarloParcelaThread(Thread):
    def __init__(self, id, parcela, inside, dots):
        super().__init__()
        self.threadid = id
        self.parcela = parcela
        self.inside = inside
        self.dots = dots

    def run(self):
        insideCirle = 0
        for i in range(self.parcela):
            # print("Thread", self.threadid, 'running', self.threadid * self.parcela + i)
            x = random.uniform(0, 1)
            y = random.uniform(0, 1)
            if math.sqrt(x*x + y*y) <= 1:
                insideCirle += 1
        self.inside.incrementa(insideCirle) 

def Parcela(dots, n_threads):
    inside = VariavelParcela()
    start = time.time()
    parcela = dots // n_threads
    threads = [MonteCarloParcelaThread(i, parcela, inside, dots) for i in range(n_threads)]
    for thread in threads:
        thread.start()
    for thread in threads:
        thread.join()
    value = 4 * inside.getValor() / dots 
    end = time.time()   
    print("Uilizando o algoritmo de parcelas, encontramos pi =", value)
    print("Tempo de processamento do algoritmo de parcelas: ", end - start)


def Sequencial(dots):
        
    def pi(dots):
        insideCircle = 0
        for i in range(dots):
            x = random.uniform(0, 1)
            y = random.uniform(0, 1)
            if math.sqrt(x*x + y*y) <= 1:
                insideCircle += 1
        return 4 * insideCircle / dots

    start = time.time()
    value = pi(dots)
    end = time.time()
    print("Uilizando o método sequencial, encontramos pi =", value)
    print("Tempo de processamento do algoritmo sequencial: ", end - start)

class VariavelBolsa():
    def __init__(self):
        self.valor = 0
        self.lock = Lock()

    def incrementa(self, value):
        self.lock.acquire()
        self.valor += value
        self.lock.release()

    def getValor(self):
        return self.valor

class MonteCarloBolsaThread(Thread):
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

def Bolsa(dots, n_threads):
    done = VariavelBolsa()
    inside = VariavelBolsa()
    start = time.time()
    threads = [MonteCarloBolsaThread(i, done, inside, dots) for i in range(n_threads)]
    for thread in threads:
        thread.start()
    for thread in threads:
        thread.join()
    value = 4 * inside.getValor() / done.getValor() 
    end = time.time()   
    print("Uilizando bolsa de tarefas, encontramos pi =", value)
    print("Tempo de processamento do algoritmo da bolsa: ", end - start)

def main():
    dots = int(input("Número de pontos: "))
    n_threads = int(input("Número de threads: "))
    if dots < 0 or n_threads < 0:
        return print("Ambos os números devem ser positivos.")
    print(f"Faremos {n_threads} threads rodando {dots} pontos.\n")
    Parcela(dots, n_threads)
    print('-------------------------------------------------------------------------------')
    Bolsa(dots, n_threads)
    print('-------------------------------------------------------------------------------')
    Sequencial(dots)

if __name__ == '__main__':
    main()