#!/usr/bin/env python
from __future__ import print_function
import sys
import random
import string
import argparse
import sox

def init_args():
    ap = argparse.ArgumentParser()
    ap.add_argument("--input")
    ap.add_argument("--output")
    args = ap.parse_args()
    if args.input == None:
        print("require input")
        exit()
    if args.output == None:
        print("require output")
        exit()
    return(args)

def randstr(n):
    r = [random.choice(string.ascii_letters + string.digits) for i in range(n)]
    return ''.join(r)

def downmix(src, dst):
    rand = randstr(20)
    tfm = sox.Transformer()
    tfm.remix({1:[1]})
    l1 = "/tmp/%s_l1.ogg" % rand
    tfm.build(src, l1)
    tfm = sox.Transformer()
    tfm.remix({1:[2]})
    c1 = "/tmp/%s_c1.ogg" % rand
    tfm.gain(-10.0)
    tfm.build(src, c1)
    tfm = sox.Transformer()
    tfm.remix({1:[3]})
    r1 = "/tmp/%s_r1.ogg" % rand
    tfm.build(src, r1)

    cbn = sox.Combiner()
    l2 = "/tmp/%s_l2.ogg" % rand
    cbn.build([l1,c1], l2, 'mix')
    cbn = sox.Combiner()
    r2 = "/tmp/%s_r2.ogg" % rand
    cbn.build([r1,c1], r2, 'mix')

    cbn = sox.Combiner()
    cbn.build([l2,r2], dst, 'merge')

def main():
    args = init_args()
    downmix(args.input, args.output)

main()
