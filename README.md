# Software Testing & Quality Assurance

## Hyphenated words
Testcase
``` javascript
it('should accept hyphenated words', () => {
    titleCase('X-men and the mission-critical mission').should.equal('X-Men and the Mission-Critical Mission');
  });

```

Regex to modify title

``` javascript
// capitalize characters after hyphen
  title = title.replace(/(-.)/g, (x) => {
    return '-' + x[1].toUpperCase();
  });

```

## Throw error on invalid input 
Testcase

``` javascript
it('should throw a TypeError for input other than a string', () => {
     should.Throw(() => titleCase(1));
   });

```

updated code

``` javascript
function titleCased(title) {
  // check type and throw error if not a string
  if (typeof(title) !== 'string') {
    throw new TypeError;
  }
  ...
```

