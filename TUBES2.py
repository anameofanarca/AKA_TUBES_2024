import math

def is_prime_recursive(n, i=2):
    if n <= 1:
        return False
    if i > int(math.sqrt(n)):
        return True
    if n % i == 0:
        return False
    return is_prime_recursive(n, i + 1)

def find_smallest_prime_recursive(numbers, index=0, smallest_prime=float('inf')):
    if index == len(numbers):
        return smallest_prime if smallest_prime != float('inf') else None
    num = numbers[index]
    if is_prime_recursive(num) and num < smallest_prime:
        smallest_prime = num
    return find_smallest_prime_recursive(numbers, index + 1, smallest_prime)

# Contoh penggunaan
# 10
#numbers = [43, 98, 91, 19, 21, 19, 19, 7, 45, 25]
# 50
#numbers = [66, 60, 64, 9, 75, 38, 88, 83, 77, 24, 84, 57, 80, 33, 100, 20, 53, 59, 54, 80, 10, 55, 42, 33, 55, 4, 72, 26, 9, 8, 10, 91, 86, 89, 13, 84, 37, 96, 38, 14, 66, 56, 36, 67, 73, 40, 63, 57, 45, 90]
# 100
#numbers = [82, 30, 8, 37, 33, 51, 79, 28, 19, 83, 96, 29, 50, 20, 99, 80, 97, 1, 65, 5, 99, 92, 43, 48, 65, 24, 13, 31, 87, 89, 6, 22, 18, 45, 57, 54, 17, 68, 1, 25, 19, 51, 82, 7, 68, 77, 34, 8, 37, 99, 80, 4, 22, 11, 74, 81, 56, 4, 95, 100, 29, 44, 26, 7, 31, 31, 48, 61, 26, 65, 99, 56, 3, 60, 36, 24, 100, 18, 43, 38, 37, 80, 78, 48, 22, 12, 98, 4, 45, 8, 72, 30, 17, 9, 90, 5, 96, 92, 26, 73]
# 500
numbers = [56, 28, 34, 54, 12, 38, 60, 7, 13, 84, 100, 9, 63, 94, 24, 31, 50, 10, 30, 74, 97, 77, 54, 2, 81, 38, 45, 64, 40, 27, 53, 4, 25, 94, 74, 85, 55, 12, 50, 96, 56, 9, 28, 4, 90, 49, 78, 22, 43, 4, 56, 31, 60, 78, 24, 87, 7, 22, 74, 17, 87, 55, 37, 27, 65, 16, 78, 86, 32, 53, 88, 91, 82, 28, 79, 65, 90, 13, 95, 5, 30, 39, 34, 78, 97, 5, 72, 69, 95, 65, 65, 12, 42, 82, 18, 74, 67, 63, 60, 77, 48, 28, 21, 9, 99, 28, 56, 24, 82, 66, 3, 43, 22, 37, 33, 68, 18, 37, 48, 7, 73, 41, 11, 31, 68, 24, 13, 70, 17, 56, 59, 12, 33, 2, 34, 91, 23, 42, 19, 33, 10, 27, 28, 79, 70, 75, 99, 59, 53, 28, 11, 76, 23, 32, 32, 59, 32, 32, 42, 100, 39, 89, 93, 6, 32, 68, 61, 53, 74, 70, 78, 19, 88, 50, 17, 9, 63, 91, 26, 63, 41, 42, 32, 14, 34, 95, 46, 71, 9, 79, 15, 93, 59, 98, 84, 73, 84, 33, 40, 65, 65, 73, 32, 74, 5, 71, 21, 16, 5, 4, 45, 2, 78, 89, 48, 46, 63, 71, 32, 39, 50, 38, 39, 86, 51, 71, 91, 18, 24, 95, 39, 85, 67, 54, 94, 56, 53, 21, 47, 61, 90, 47, 31, 48, 31, 38, 87, 26, 96, 42, 75, 20, 100, 38, 52, 38, 61, 33, 91, 85, 25, 12, 16, 82, 11, 52, 51, 37, 18, 88, 82, 49, 61, 78, 52, 34, 31, 74, 50, 20, 98, 25, 54, 22, 6, 77, 57, 4, 32, 11, 55, 73, 84, 24, 31, 36, 97, 24, 28, 13, 24, 8, 25, 73, 18, 6, 72, 94, 73, 52, 43, 17, 99, 71, 22, 47, 18, 62, 71, 8, 84, 32, 39, 57, 3, 88, 22, 96, 34, 92, 19, 20, 77, 91, 27, 56, 6, 68, 73, 99, 36, 48, 66, 56, 43, 80, 89, 57, 5, 45, 4, 95, 84, 20, 80, 2, 27, 85, 79, 9, 18, 56, 24, 83, 44, 35, 83, 27, 87, 68, 42, 15, 88, 60, 31, 31, 6, 91, 19, 51, 79, 52, 88, 1, 65, 9, 54, 30, 45, 30, 100, 69, 55, 52, 83, 92, 63, 100, 16, 19, 18, 39, 32, 39, 93, 22, 12, 14, 64, 35, 86, 79, 54, 55, 78, 68, 7, 37, 6, 18, 76, 46, 55, 73, 48, 73, 13, 42, 3, 8, 49, 22, 22, 48, 6, 39, 92, 39, 43, 24, 92, 63, 99, 54, 11, 92, 77, 10, 20, 72, 2, 17, 95, 59, 2, 83, 57, 93, 70, 54, 16, 57, 19, 84, 65, 7, 40, 47, 59, 5, 21, 95, 29, 87, 92, 92, 97, 19, 61, 3, 90, 91, 62, 91, 78, 28, 87, 15, 83, 25, 41, 21, 19, 13, 53, 57, 90, 46, 67, 45]
#1000
#numbers = [56, 28, 34, 54, 12, 38, 60, 7, 13, 84, 100, 9, 63, 94, 24, 31, 50, 10, 30, 74, 97, 77, 54, 2, 81, 38, 45, 64, 40, 27, 53, 4, 25, 94, 74, 85, 55, 2, 50, 96, 56, 9, 28, 4, 90, 49, 78, 22, 43, 4, 56, 31, 60, 78, 24, 87, 7, 22, 74, 17, 87, 55, 37, 27, 65, 16, 78, 86, 32, 53, 88, 91, 82, 28, 79, 65, 90, 13, 95, 5, 30, 39, 34, 78, 97, 5, 72, 69, 95, 65, 65, 12, 42, 82, 18, 74, 67, 63, 60, 77, 48, 28, 21, 9, 99, 28, 56, 24, 82, 66, 3, 43, 22, 37, 33, 68, 18, 37, 48, 7, 73, 41, 11, 31, 68, 24, 13, 70, 17, 56, 59, 12, 33, 2, 34, 91, 23, 42, 19, 33, 10, 27, 28, 79, 70, 75, 99, 59, 53, 28, 11, 76, 23, 32, 32, 59, 32, 32, 42, 100, 39, 89, 93, 6, 32, 68, 61, 53, 74, 70, 78, 19, 88, 50, 17, 9, 63, 91, 26, 63, 41, 42, 32, 14, 34, 95, 46, 71, 9, 79, 15, 93, 59, 98, 84, 73, 84, 33, 40, 65, 65, 73, 32, 74, 5, 71, 21, 16, 5, 4, 45, 2, 78, 89, 48, 46, 63, 71, 32, 39, 50, 38, 39, 86, 51, 71, 91, 18, 24, 95, 39, 85, 67, 54, 94, 56, 53, 21, 47, 61, 90, 47, 31, 48, 31, 38, 87, 26, 96, 42, 75, 20, 100, 38, 52, 38, 61, 33, 91, 85, 25, 12, 16, 82, 11, 52, 51, 37, 18, 88, 82, 49, 61, 78, 52, 34, 31, 74, 50, 20, 98, 25, 54, 22, 6, 77, 57, 4, 32, 11, 55, 73, 84, 24, 31, 36, 97, 24, 28, 13, 24, 8, 25, 73, 18, 6, 72, 94, 73, 52, 43, 17, 99, 71, 22, 47, 18, 62, 71, 8, 84, 32, 39, 57, 3, 88, 22, 96, 34, 92, 19, 20, 77, 91, 27, 56, 6, 68, 73, 99, 36, 48, 66, 56, 43, 80, 89, 57, 5, 45, 4, 95, 84, 20, 80, 2, 27, 85, 79, 9, 18, 56, 24, 83, 44, 35, 83, 27, 87, 68, 42, 15, 88, 60, 31, 31, 6, 91, 19, 51, 79, 52, 88, 1, 65, 9, 54, 30, 45, 30, 100, 69, 55, 52, 83, 92, 63, 100, 16, 19, 18, 39, 32, 39, 93, 22, 12, 14, 64, 35, 86, 79, 54, 55, 78, 68, 7, 37, 6, 18, 76, 46, 55, 73, 48, 73, 13, 42, 3, 8, 49, 22, 22, 48, 6, 39, 92, 39, 43, 24, 92, 63, 99, 54, 11, 92, 77, 10, 20, 72, 2, 17, 95, 59, 2, 83, 57, 93, 70, 54, 16, 57, 19, 84, 65, 7, 40, 47, 59, 5, 21, 95, 29, 87, 92, 92, 97, 19, 61, 3, 90, 91, 62, 91, 78, 28, 87, 15, 83, 25, 41, 21, 19, 13, 53, 57, 90, 46, 67, 45,82, 30, 8, 37, 33, 51, 79, 28, 19, 83, 96, 29, 50, 20, 99, 80, 97, 1, 65, 5, 99, 92, 43, 48, 65, 24, 3, 31, 87, 89, 6, 22, 18, 45, 57, 54, 17, 68, 1, 25, 19, 51, 82, 7, 68, 77, 34, 8, 37, 99, 80, 4, 22, 11, 74, 81, 56, 4, 95, 100, 29, 44, 26, 7, 31, 31, 48, 61, 26, 65, 99, 56, 3, 60, 36, 24, 100, 18, 43, 38, 37, 80, 78, 48, 22, 12, 98, 4, 45, 8, 72, 30, 17, 9, 90, 5, 96, 92, 26, 73,82, 30, 8, 37, 33, 51, 79, 28, 19, 83, 96, 29, 50, 20, 99, 80, 97, 1, 65, 5, 99, 92, 43, 48, 65, 24, 3, 31, 87, 89, 6, 22, 18, 45, 57, 54, 17, 68, 1, 25, 19, 51, 82, 7, 68, 77, 34, 8, 37, 99, 80, 4, 22, 11, 74, 81, 56, 4, 95, 100, 29, 44, 26, 7, 31, 31, 48, 61, 26, 65, 99, 56, 3, 60, 36, 24, 100, 18, 43, 38, 37, 80, 78, 48, 22, 12, 98, 4, 45, 8, 72, 30, 17, 9, 90, 5, 96, 92, 26, 73,82, 30, 8, 37, 33, 51, 79, 28, 19, 83, 96, 29, 50, 20, 99, 80, 97, 1, 65, 5, 99, 92, 43, 48, 65, 24, 3, 31, 87, 89, 6, 22, 18, 45, 57, 54, 17, 68, 1, 25, 19, 51, 82, 7, 68, 77, 34, 8, 37, 99, 80, 4, 22, 11, 74, 81, 56, 4, 95, 100, 29, 44, 26, 7, 31, 31, 48, 61, 26, 65, 99, 56, 3, 60, 36, 24, 100, 18, 43, 38, 37, 80, 78, 48, 22, 12, 98, 4, 45, 8, 72, 30, 17, 9, 90, 5, 96, 92, 26, 73,82, 30, 8, 37, 33, 51, 79, 28, 19, 83, 96, 29, 50, 20, 99, 80, 97, 1, 65, 5, 99, 92, 43, 48, 65, 24, 3, 31, 87, 89, 6, 22, 18, 45, 57, 54, 17, 68, 1, 25, 19, 51, 82, 7, 68, 77, 34, 8, 37, 99, 80, 4, 22, 11, 74, 81, 56, 4, 95, 100, 29, 44, 26, 7, 31, 31, 48, 61, 26, 65, 99, 56, 3, 60, 36, 24, 100, 18, 43, 38, 37, 80, 78, 48, 22, 12, 98, 4, 45, 8, 72, 30, 17, 9, 90, 5, 96, 92, 26, 73,82, 30, 8, 37, 33, 51, 79, 28, 19, 83, 96, 29, 50, 20, 99, 80, 97, 1, 65, 5, 99, 92, 43, 48, 65, 24, 3, 31, 87, 89, 6, 22, 18, 45, 57, 54, 17, 68, 1, 25, 19, 51, 82, 7, 68, 77, 34, 8, 37, 99, 80, 4, 22, 11, 74, 81, 56, 4, 95, 100, 29, 44, 26, 7, 31, 31, 48, 61, 26, 65, 99, 56, 3, 60, 36, 24, 100, 18, 43, 38, 37, 80, 78, 48, 22, 12, 98, 4, 45, 8, 72, 30, 17, 9, 90, 5, 96, 92, 26, 73]
#5000
result = find_smallest_prime_recursive(numbers)
print(f"Bilangan prima terkecil (rekursif) adalah: {result}")