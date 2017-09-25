#!/usr/bin/python
# requires python-zxcvbn, https://pypi.python.org/pypi/zxcvbn-python
# https://github.com/dwolfhub/zxcvbn-python
from zxcvbn import zxcvbn

results = zxcvbn('salakala')

print results 
