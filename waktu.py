import time
import random
import math
import sys

# Tingkatkan batas kedalaman rekursi untuk mencegah RecursionError pada input besar
sys.setrecursionlimit(3000)

# Fungsi untuk memeriksa bilangan prima (iteratif)
def is_prime_iterative(n):
    if n <= 1:
        return False
    for i in range(2, int(math.sqrt(n)) + 1):
        if n % i == 0:
            return False
    return True

# Fungsi untuk mencari bilangan prima terkecil secara iteratif
def find_smallest_prime_iterative(numbers):
    smallest_prime = float('inf')  # Representasi nilai tak terhingga
    for num in numbers:
        if is_prime_iterative(num) and num < smallest_prime:
            smallest_prime = num
    return smallest_prime if smallest_prime != float('inf') else None

# Fungsi untuk memeriksa bilangan prima (rekursif)
def is_prime_recursive(n, i=2):
    if n <= 1:
        return False
    if i > int(math.sqrt(n)):
        return True
    if n % i == 0:
        return False
    return is_prime_recursive(n, i + 1)

# Fungsi untuk mencari bilangan prima terkecil secara rekursif
def find_smallest_prime_recursive(numbers, index=0, smallest_prime=float('inf')):
    if index == len(numbers):
        return smallest_prime if smallest_prime != float('inf') else None
    num = numbers[index]
    if is_prime_recursive(num) and num < smallest_prime:
        smallest_prime = num
    return find_smallest_prime_recursive(numbers, index + 1, smallest_prime)

# Fungsi untuk menghasilkan daftar bilangan acak
def generate_random_numbers(size, min_val=1, max_val=1000):
    return [random.randint(min_val, max_val) for _ in range(size)]

# Uji dengan ukuran input yang berbeda dan tampilkan waktu dalam milidetik
input_sizes = [5, 25, 50, 100, 200, 600, 800, 1100, 1300, 1500, 1700, 2000, 2200, 2500]  # Batas disesuaikan untuk mencegah error

print("| Ukuran Input | Metode      | Waktu Eksekusi (ms) | Bilangan Prima Terkecil |")
print("|--------------|-------------|---------------------|--------------------------|")

for size in input_sizes:
    numbers = generate_random_numbers(size)

    # Pengujian metode iteratif
    start_time = time.time()
    smallest_iterative = find_smallest_prime_iterative(numbers)
    end_time = time.time()
    iterative_time_ms = (end_time - start_time) * 1000  # Konversi ke milidetik
    smallest_iterative_str = smallest_iterative if smallest_iterative is not None else "Tidak Ada"
    print(f"| {size:<12} | Iteratif    | {iterative_time_ms:<19.2f} | {smallest_iterative_str:<24} |")

    # Pengujian metode rekursif
    start_time = time.time()
    smallest_recursive = find_smallest_prime_recursive(numbers)
    end_time = time.time()
    recursive_time_ms = (end_time - start_time) * 1000  # Konversi ke milidetik
    smallest_recursive_str = smallest_recursive if smallest_recursive is not None else "Tidak Ada"
    print(f"| {size:<12} | Rekursif    | {recursive_time_ms:<19.2f} | {smallest_recursive_str:<24} |")
