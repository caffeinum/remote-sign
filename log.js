function time() {
  return new Date().toLocaleTimeString()
}

function log(...args) {
  args.unshift('[' + time() + ']')
  console.log.apply(console, args)
}

module.exports = log
