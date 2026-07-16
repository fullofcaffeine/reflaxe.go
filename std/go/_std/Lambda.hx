/*
 * Copyright (C)2005-2019 Haxe Foundation
 *
 * Permission is hereby granted, free of charge, to any person obtaining a
 * copy of this software and associated documentation files (the "Software"),
 * to deal in the Software without restriction, including without limitation
 * the rights to use, copy, modify, merge, publish, distribute, sublicense,
 * and/or sell copies of the Software, and to permit persons to whom the
 * Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
 * FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
 * DEALINGS IN THE SOFTWARE.
 */

import haxe.ds.List;

/**
	What
	- Implements the portable `Lambda` collection helpers as ordinary staged Haxe
	  source.

	Why
	- The mainstream Haxe stdlib implementation cannot be used unchanged on
	  `haxe.go` yet because its emitted generic helpers erase `Iterable<T>` to a
	  structural Go carrier. Native slices, staged lists, concrete iterator classes,
	  and typed callbacks are not assignable to that carrier under Go's invariant
	  type system.

	How
	- Keep every algorithm as a self-contained source function over `Iterable<T>`.
	  Exact call adapters cover all 19 public helpers after
	  `haxe_go-vfp.8.7.17` and `haxe_go-vfp.8.7.18`; they never implement traversal,
	  filtering, mapping, folding, or early-exit behavior. Preserve concrete nested
	  iterable types with constrained parameters so the Go boundary can build only
	  the required representation wrapper. `@:dce` eliminates unused emitted
	  helpers.
**/
@:dce
class Lambda {
	/**
		What
		- Gives compiler-generated `flatMap` callback adapters a nominal value that
		  implements the erased Go iterator method set.

		Why
		- Go does not let the map-shaped structural `Iterable<T>` carrier satisfy a
		  method interface directly, even though both represent the same Haxe
		  contract. The mainstream implementation therefore cannot receive an
		  adapted array, list, or concrete iterable callback result unchanged.

		How
		- Wrap the already-adapted structural iterable in the private source-declared
		  carrier. Retain this factory only when `flatten` or `flatMap` is reachable;
		  the exact compiler representation adapter calls it after Haxe dead-code
		  elimination has run.
	**/
	@:ifFeature("Lambda.flatten", "Lambda.flatMap")
	@:noCompletion
	public static function goIterableCarrierAdapter<A>(source:Iterable<A>):LambdaGoIterableCarrier<A> {
		return new LambdaGoIterableCarrier(source);
	}

	public static function array<A>(it:Iterable<A>):Array<A> {
		var out:Array<A> = [];
		for (value in it) {
			out.push(value);
		}
		return out;
	}

	public static function list<A>(it:Iterable<A>):List<A> {
		var out = new List<A>();
		for (value in it) {
			out.add(value);
		}
		return out;
	}

	public static function map<A, B>(it:Iterable<A>, transform:(item:A) -> B):Array<B> {
		var out:Array<B> = [];
		for (value in it) {
			out.push(transform(value));
		}
		return out;
	}

	public static function mapi<A, B>(it:Iterable<A>, transform:(index:Int, item:A) -> B):Array<B> {
		var out:Array<B> = [];
		var index = 0;
		for (value in it) {
			out.push(transform(index, value));
			index++;
		}
		return out;
	}

	public static function flatten<A, Nested:Iterable<A>>(it:Iterable<Nested>):Array<A> {
		var out:Array<A> = [];
		for (nested in it) {
			for (value in nested) {
				out.push(value);
			}
		}
		return out;
	}

	public static function flatMap<A, B, Nested:Iterable<B>>(it:Iterable<A>, transform:(item:A) -> Nested):Array<B> {
		var out:Array<B> = [];
		for (value in it) {
			for (mapped in transform(value)) {
				out.push(mapped);
			}
		}
		return out;
	}

	public static function has<A>(it:Iterable<A>, expected:A):Bool {
		var found = false;
		for (value in it) {
			if (value == expected) {
				found = true;
				break;
			}
		}
		return found;
	}

	public static function exists<A>(it:Iterable<A>, predicate:(item:A) -> Bool):Bool {
		var found = false;
		for (value in it) {
			if (predicate(value)) {
				found = true;
				break;
			}
		}
		return found;
	}

	public static function foreach<A>(it:Iterable<A>, predicate:(item:A) -> Bool):Bool {
		var allMatched = true;
		for (value in it) {
			if (!predicate(value)) {
				allMatched = false;
				break;
			}
		}
		return allMatched;
	}

	public static function iter<A>(it:Iterable<A>, visit:(item:A) -> Void):Void {
		for (value in it) {
			visit(value);
		}
	}

	public static function filter<A>(it:Iterable<A>, predicate:(item:A) -> Bool):Array<A> {
		var out:Array<A> = [];
		for (value in it) {
			if (predicate(value)) {
				out.push(value);
			}
		}
		return out;
	}

	public static function fold<A, B>(it:Iterable<A>, combine:(item:A, result:B) -> B, first:B):B {
		for (value in it) {
			first = combine(value, first);
		}
		return first;
	}

	public static function foldi<A, B>(it:Iterable<A>, combine:(item:A, result:B, index:Int) -> B, first:B):B {
		var index = 0;
		for (value in it) {
			first = combine(value, first, index);
			index++;
		}
		return first;
	}

	public static function count<A>(it:Iterable<A>, ?predicate:(item:A) -> Bool):Int {
		var result = 0;
		if (predicate == null) {
			for (_ in it) {
				result++;
			}
		} else {
			for (value in it) {
				if (predicate(value)) {
					result++;
				}
			}
		}
		return result;
	}

	public static function empty<A>(it:Iterable<A>):Bool {
		var result = true;
		for (_ in it) {
			result = false;
			break;
		}
		return result;
	}

	public static function indexOf<A>(it:Iterable<A>, expected:A):Int {
		var index = 0;
		var found = false;
		for (value in it) {
			if (value == expected) {
				found = true;
				break;
			}
			index++;
		}
		return found ? index : -1;
	}

	public static function find<A>(it:Iterable<A>, predicate:(item:A) -> Bool):Null<A> {
		var found:Null<A> = null;
		for (value in it) {
			if (predicate(value)) {
				found = value;
				break;
			}
		}
		return found;
	}

	public static function findIndex<A>(it:Iterable<A>, predicate:(item:A) -> Bool):Int {
		var index = 0;
		var found = false;
		for (value in it) {
			if (predicate(value)) {
				found = true;
				break;
			}
			index++;
		}
		return found ? index : -1;
	}

	public static function concat<A>(left:Iterable<A>, right:Iterable<A>):Array<A> {
		var out:Array<A> = [];
		for (value in left) {
			out.push(value);
		}
		for (value in right) {
			out.push(value);
		}
		return out;
	}
}

/**
	What
	- Provides a nominal Go-target view of the structural `Iterable<T>` value used
	  by the staged `Lambda.flatten` and `Lambda.flatMap` implementations.

	Why
	- The mainstream Haxe stdlib implementation cannot be used unchanged on
	  `haxe.go` because a constrained generic iterable becomes a Go method
	  interface, while portable Haxe iterables currently use a structural map.
	  Go does not let that map satisfy the method interface directly even though
	  both values represent the same Haxe iterator contract.

	How
	- Store the already-adapted structural iterable and delegate `iterator()` to
	  it. This private companion performs no traversal and owns no collection
	  algorithm; `Lambda` retains all public behavior in ordinary staged source.
**/
private class LambdaGoIterableCarrier<A> {
	final source:Iterable<A>;

	public function new(source:Iterable<A>) {
		this.source = source;
	}

	public function iterator():Iterator<A> {
		return source.iterator();
	}
}
