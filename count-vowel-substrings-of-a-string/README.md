# 2062. Count Vowel Substrings of a String

A **substring** is a contiguous (non-empty) sequence of characters within a string.

A **vowel substring** is a substring that **only** consists of vowels (`a`, `e`, `i`, `o`, and `u`) and has **all five** vowels present in it.

Given a string `word`, return *the number of **vowel substrings** in `word`*.

Return *the minimum **score** of `nums` after applying the mentioned operation at most once for each index in it*.

**Example 1:**  
    **Input:** `word = "aeiouu"`  
    **Output:** `2`  
    **Explanation:** 
```
The vowel substrings of word are as follows (underlined):
- "aeiouu"
- "aeiouu"
```  

**Example 2:**  
    **Input:** `nums = [0,10], k = 2`  
    **Output:** `6`  
    **Explanation:** 
```
Not all 5 vowels are present, so there are no vowel substrings.
```  

**Example 3:**  
    **Input:** `nums = [1,3,6], k = 3`  
    **Output:** `0`  
    **Explanation:** 
```
The vowel substrings of word are as follows (underlined):
- "cuaieuouac"
- "cuaieuouac"
- "cuaieuouac"
- "cuaieuouac"
- "cuaieuouac"
- "cuaieuouac"
- "cuaieuouac"
```   

**Constraints:**  
    `1 <= word.length <= 100`  
    `word consists of lowercase English letters only.`  

## Submission Screenshot

![Image](./count-vowel-substrings-of-a-string.png)