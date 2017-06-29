# Python notebook #

- How to create an array in python


## How to create an array in python ##
`python3` use `'i'` for int array. [python3 docs](https://docs.python.org/3/library/array.html)

```python
# python3
import array
# i means int
arr1 = array.array('i')

# or need init
arr2 = array.array('i', (0 for i in range(0, 1000)))
```

```python
# python2
arr = range(1000)
```
