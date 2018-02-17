#!/usr/bin/env python
import string


def longest_word(sen):
    longest = ""
    words = sen.translate(None, string.punctuation)
    words = words.split(' ')
    for word in words:
        if len(longest) < len(word):
            longest = word
    return longest


print longest_word(raw_input())