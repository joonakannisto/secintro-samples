#!/usr/bin/python

import math

PRIME = 4294967087
GENERATOR = 2

print PRIME
print GENERATOR

# Mark = M
# Jonathan = J

M_Secret =123

print M_Secret

M_Comp_1 = (math.pow(GENERATOR,M_Secret) % PRIME )

print M_Comp_1
