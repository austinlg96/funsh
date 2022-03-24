import random


print("Level 5")
for i in range(0,10):
    x = int(f'{random.randint(0,9)}{random.randint(0,5)}')
    y = int(f'{random.randint(0,5)}{random.randint(0,5)}')
    
    output  = "{" + f'"{x}*{y}","{x*y}", 12345' + "},"
    print(output)


print("Level 6")
for i in range(0,10):
    x = random.randint(0,9)
    y = random.randint(0,9)
    z = random.randint(0,9)
    o = random.randint(0,1)
    o1 = "+" if o else "*"
    o2 = "*" if o else "+"
    ans = x + (y * z) if o else (x * y) + z
    output  = "{" + f'"{x}{o1}{y}{o2}{z}","{ans}", 6789' + "},"
    print(output)