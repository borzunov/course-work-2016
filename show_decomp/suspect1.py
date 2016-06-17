import sys

n = int(sys.argv[1])  # n >= 2

w = ''
for i in range(n):
    w = '10' * i + '0' + w
print(w)
