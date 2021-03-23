package nexe

const Lib = BigJS + `
function $v(a) {
	if (a === undefined) {
		return "undefined";
	}

	if (a === true || a === false) {
		return a;
	}

	if (a instanceof Big) {
		return a.toString();
	}

	if (typeof a === 'string') {
		return a;
	}

	if (a.t.startsWith('[]')) {
		return a.v.map(b => $v(b));
	}

	return a.v
}

function $json(v) {
	if (typeof v === 'string') {
		return '"' + v + '"';
	}

	if (v === true) {
		return 'true';
	}

	if (v === false) {
		return 'false';
	}

	if (v instanceof Big) {
		return v.toString();
	}

	if (v[1][0] === '[') {
		return '[' + v[0].map(x => $json(x)).join(', ') + ']';
	}

	// Maps or objects.
	return '{' + Object.keys(v[0]).sort().map(key => 
		'"' + key + '": ' + $json(v[0][key])).join(', ') + '}';
}

function $string(v) {
	if (typeof v === 'string') {
		return v;
	}

	if (v === true) {
		return 'true';
	}

	if (v === false) {
		return 'false';
	}

	if (v instanceof Big) {
		return v.toString();
	}

	return $json(v);
}

const $callStack = [];

function $call(pos, fn) {
	$callStack.push(pos);
	return fn;
}

function $stack(e) {
	// This is for safety in case a JavaScript exception is thrown, which
	// shouldn't happen but if it does just print it out for debugging.
	if (!(e instanceof $Error)) {
		console.log(e);
		return [];
	}

	const stack = e.stack.
      split("\n").
      filter(x => x.includes('_fn')).
      map(x => x.match(/_fn\d+/)[0]);

    return 'Error: ' + $v(e.obj.Error) + '\n' + stack.
      map((x, index) => {
        const fn = tests__error_unhandled.$fns[x];
        return '  ' + (stack.length - index) + ' ' + fn[0] + '() at ' +
			$callStack[$callStack.length - index - 1];
      }).
      join('\n');
}

class $Error extends Error {
  constructor(obj) {
    super("");
    this.obj = obj;
  }
}

function print() {
	console.log(...Array.from(arguments).map(a => $string(a)));
	return [];
}

`
