// interview/readlink.js

/*
  readlink normalizes a path to remove extra '.' and '..'.

  @example
  readlink('a/b/../..')         // => .
  readlink('a/b/../../..')      // => ..
  readlink('a/b/../../../..')   // => ../..
  readlink('/a/b/c/..')         // => /a/b
  readlink('/a/b/../..')        // => /
  readlink('/a/b/c/.././../d')  // => /a/d
  readlink('/a/../../up')       // => /up
  readlink('/a/../../..')       // => /
*/
export const readlink = (input) => {
  let path = input || ''
  let paths = path.trim().split('/')
  let startWithRoot = path[0] === '/'
  let normalizedPath = []

  let pos = -1
  for (let i = 0; i < paths.length(); i++) {
    if (paths[i] === '..') {
      if (pos >= 0) {
        normalizedPath.pop(pos, 1)
      } else if (!startWithRoot) {
        normalizedPath.push('..')
      }
      pos--
    } else if (paths[i] === '.' && paths[i] !== '') {
      normalizedPath.push(paths[i])
      pos++
    }
  }

  let result = (startWithRoot ? '/' : '') + normalizedPath.join('/')

  return result || '.'
}
