#!/usr/bin/env python
import base64,array,binascii
b64Str = "WUVMTE9XIFNVQk1BUklORQ=="
bArr = bytearray(base64.b64decode(b64Str))
for i in range(len(bArr)):
    bArr[i] = bArr[i] | 0x20
print bArr
print "Base64: "+base64.b64encode(bArr)
